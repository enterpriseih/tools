package kafkalib

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jedib0t/go-pretty/table"

	"os"
)

type PartitionsDirSize struct {
	Partition int32
	Size      int64
}

type BrokersPartitionSize struct {
	Broker     int32
	Partitions []PartitionsDirSize
}

func (k *Kafka) TopicDirs(topic *Topic) *Topic {
	bpds := make([]BrokersPartitionSize, 0)
	topic.PartitionDirSize = make(map[int32]int64, topic.PartitionCount)

	for _, broker := range k.client.Brokers() {
		if err := broker.Open(k.client.Config()); err != nil && err != sarama.ErrAlreadyConnected {
			fmt.Printf("Cannot connect to broker %d: %v", broker.ID(), err)
		}
		defer broker.Close()

		var ddt sarama.DescribeLogDirsRequestTopic
		ddt.Topic = topic.Name
		ddt.PartitionIDs = topic.Partitions

		describeLogDirs, _ := broker.DescribeLogDirs(&sarama.DescribeLogDirsRequest{
			Version:        0,
			DescribeTopics: []sarama.DescribeLogDirsRequestTopic{ddt},
		})
		var bps BrokersPartitionSize
		bps.Broker = broker.ID()
		bps.Partitions = make([]PartitionsDirSize, 0)
		for _, describeLogDirsResponseDirMetadata := range describeLogDirs.LogDirs {
			for _, t := range describeLogDirsResponseDirMetadata.Topics {
				for _, describeLogDirsResponsePartition := range t.Partitions {
					// Leader Data Size
					if topic.PartitionLeader[describeLogDirsResponsePartition.PartitionID] == broker.ID() {
						topic.PartitionDirSize[describeLogDirsResponsePartition.PartitionID] = describeLogDirsResponsePartition.Size
					}

					// ALL Partition Broker Size
					var pds PartitionsDirSize
					pds.Partition = describeLogDirsResponsePartition.PartitionID
					pds.Size = describeLogDirsResponsePartition.Size
					bps.Partitions = append(bps.Partitions, pds)
				}
			}
		}
		if len(bps.Partitions) > 0 {
			bpds = append(bpds, bps)
		}

	}
	topic.BrokerPartitionDirSize = bpds
	return topic
}
func (k *Kafka) CommandTopicDirs(tp Topic) {
	//tp := k.TopicInfo(topic)
	//k.TopicDirs(&tp)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(fmt.Sprintf("Topics: %s", tp.Name))
	t.AppendHeader(table.Row{"Partition", "Broker", "Size"})
	for _, bpds := range tp.BrokerPartitionDirSize {
		for _, bpd := range bpds.Partitions {
			t.AppendRow([]interface{}{bpd.Partition, bpds.Broker, bytefmt.ByteSize(uint64(bpd.Size) * bytefmt.BYTE)})
		}
	}
	var sort = table.SortBy{
		Name:   "Partition",
		Number: 1,
		Mode:   0,
	}
	t.SortBy([]table.SortBy{sort})
	t.Render()
}
