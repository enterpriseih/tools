package kafkalib

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

type Topics []Topic

type Topic struct {
	Name                      string `json:"Name"`
	RetentionBytes            int
	RetentionMs               int
	PartitionCount            int32                  `json:"PartitionCount"`
	ReplicationFactor         int16                  `json:"ReplicationFactor"`
	Partitions                []int32                `json:"Partitions"`
	PartitionLeader           map[int32]int32        `json:"PartitionLeader"`
	Replicas                  map[int32][]int32      `json:"Replicas"`
	InSyncReplicas            map[int32][]int32      `json:"InSyncReplicas"`
	InSyncReplicaNum          int                    `json:"InSyncReplicaNum"`
	PartitionCurrentOffset    map[int32]int64        `json:"PartitionCurrentOffset"`
	PartitionOldestOffset     map[int32]int64        `json:"PartitionOldestOffset"`
	PartitionCurrentOffsetSum int64                  `json:"PartitionCurrentOffsetSum"`
	Groups                    []Group                `json:"Groups"`
	PartitionDirSize          map[int32]int64        `json:"PartitionDirSize"`
	BrokerPartitionDirSize    []BrokersPartitionSize `json:"BrokerPartitionDirSize"` // topic[broker][partition]size
}

func (k *Kafka) TopicInfo(topic string) Topic {
	var t Topic

	t.Name = topic
	t.Partitions, _ = k.client.Partitions(topic)
	t.PartitionCount = int32(len(t.Partitions))
	t.PartitionLeader = make(map[int32]int32, t.PartitionCount)
	t.PartitionCurrentOffset = make(map[int32]int64, t.PartitionCount)
	t.PartitionOldestOffset = make(map[int32]int64, t.PartitionCount)

	t.Replicas = make(map[int32][]int32, t.PartitionCount)
	t.InSyncReplicas = make(map[int32][]int32, t.PartitionCount)

	for _, partition := range t.Partitions {
		broker, err := k.client.Leader(topic, partition)
		if err != nil {
			fmt.Println(err)
		}
		t.PartitionLeader[partition] = broker.ID()
		// Current Offset
		t.PartitionCurrentOffset[partition], err = k.client.GetOffset(topic, partition, sarama.OffsetNewest)
		t.PartitionCurrentOffsetSum += t.PartitionCurrentOffset[partition]

		//  Oldest Offset
		t.PartitionOldestOffset[partition], err = k.client.GetOffset(topic, partition, sarama.OffsetOldest)

		// Replicas
		t.Replicas[partition], err = k.client.Replicas(topic, partition)
		if t.ReplicationFactor < int16(len(t.Replicas[partition])) {
			t.ReplicationFactor = int16(len(t.Replicas[partition]))
		}
		// InSync
		t.InSyncReplicas[partition], err = k.client.InSyncReplicas(topic, partition)
		if t.InSyncReplicaNum == 0 {
			t.InSyncReplicaNum = len(t.InSyncReplicas[partition])
		}
		if t.InSyncReplicaNum > len(t.InSyncReplicas[partition]) {
			t.InSyncReplicaNum = len(t.InSyncReplicas[partition])
		}
	}
	return t
}

func (k *Kafka) ListTopics() *Topics {
	topics, _ := k.client.Topics()

	var topicsInfo Topics
	for _, topic := range topics {
		if topic != "__consumer_offsets" {
			topicsInfo = append(topicsInfo, k.TopicInfo(topic))
		}
	}
	return &topicsInfo
}
func (k *Kafka) CommandListTopics() {
	topics := k.ListTopics()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Topic", "Partition Num", "Replica Num", "InSync Num", "PartitionCurrentOffset Sum"})
	for _, i := range *topics {
		t.AppendRow([]interface{}{i.Name, i.PartitionCount, i.ReplicationFactor, i.InSyncReplicaNum, i.PartitionCurrentOffsetSum})
	}
	t.Render()
}

//func (k *Kafka) CommandDescribeTopic(topic string) {
func (k *Kafka) CommandDescribeTopic(tp Topic) {

	//	tp := k.TopicInfo(topic)
	//k.TopicDirs(&tp)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(fmt.Sprintf("Topics: %s", tp.Name))
	t.AppendHeader(table.Row{"Partition", "Leader", "Replicas", "InSync", "PartitionCurrentOffset", "PartitionOldestOffset", "Data Len", "Data Size"})
	for _, i := range tp.Partitions {
		t.AppendRow([]interface{}{i, tp.PartitionLeader[i], tp.Replicas[i], tp.InSyncReplicas[i], tp.PartitionCurrentOffset[i],
			tp.PartitionOldestOffset[i], tp.PartitionCurrentOffset[i] - tp.PartitionOldestOffset[i],
			bytefmt.ByteSize(uint64(tp.PartitionDirSize[i]) * bytefmt.BYTE)})
	}
	t.Render()
}
