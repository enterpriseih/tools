package kafkalib

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

type Group struct {
	Name                   string          `json:"Name"`
	Lag                    map[int32]int64 `json:"Lag"`
	LagSum                 int64           `json:"LagSum"`
	PartitionCurrentOffset map[int32]int64 `json:"PartitionCurrentOffset"`
	Members                []string
}

func (k *Kafka) TopicGroups(t *Topic) *Topic {

	groupIds := make([]string, 0)
	topicGroups := make([]Group, 0)

	for _, broker := range k.client.Brokers() {
		if err := broker.Open(k.client.Config()); err != nil && err != sarama.ErrAlreadyConnected {
			fmt.Printf("Cannot connect to broker %d: %v", broker.ID(), err)
		}

		defer broker.Close()
		groups, err := broker.ListGroups(&sarama.ListGroupsRequest{})
		if err != nil {
			fmt.Printf("Cannot get consumer group: %v", err)
		}
		for groupId := range groups.Groups {
			groupIds = append(groupIds, groupId)
		}
		describeGroups, err := broker.DescribeGroups(&sarama.DescribeGroupsRequest{Groups: groupIds})
		if err != nil {
			fmt.Printf("Cannot get describe groups: %v", err)
		}
		var g Group
		for _, group := range describeGroups.Groups {

			offsetFetchRequest := sarama.OffsetFetchRequest{ConsumerGroup: group.GroupId, Version: 1}
			for _, partition := range t.Partitions {
				offsetFetchRequest.AddPartition(t.Name, partition)
			}
			if offsetFetchResponse, err := broker.FetchOffset(&offsetFetchRequest); err != nil {
				fmt.Printf("Cannot get offset of group %s: %v", group.GroupId, err)
			} else {

				for topic, partitions := range offsetFetchResponse.Blocks {
					// If the topic is not consumed by that consumer group, skip it
					topicConsumed := false
					for _, offsetFetchResponseBlock := range partitions {
						// Kafka will return -1 if there is no offset associated with a topic-partition under that consumer group
						if offsetFetchResponseBlock.Offset != -1 {
							topicConsumed = true
							break
						}
					}

					g.PartitionCurrentOffset = make(map[int32]int64, t.PartitionCount)
					g.Lag = make(map[int32]int64, t.PartitionCount)
					if topicConsumed {
						var currentOffsetSum int64
						var lagSum int64
						for partition, offsetFetchResponseBlock := range partitions {
							err := offsetFetchResponseBlock.Err
							if err != sarama.ErrNoError {
								fmt.Printf("Error for  partition %d :%v", partition, err.Error())
								continue
							}
							currentOffset := offsetFetchResponseBlock.Offset
							currentOffsetSum += currentOffset
							g.PartitionCurrentOffset[partition] = currentOffset

							if offset, ok := t.PartitionCurrentOffset[partition]; ok {
								// If the topic is consumed by that consumer group, but no offset associated with the partition
								// forcing lag to -1 to be able to alert on that
								var lag int64
								if offsetFetchResponseBlock.Offset == -1 {
									lag = -1
								} else {
									lag = offset - offsetFetchResponseBlock.Offset
									lagSum += lag
								}
								g.Lag[partition] = lag

							} else {
								fmt.Printf("No offset of topic %s partition %d, cannot get consumer group lag", topic, partition)
							}
							g.Name = group.GroupId
							g.LagSum = lagSum

						}
					}
				}
			}
			if len(g.PartitionCurrentOffset) > 0 {
				members := make([]string, 0)
				for _, i := range group.Members {
					gmm, _ := i.GetMemberMetadata()
					for _, gmmT := range gmm.Topics {
						if gmmT == t.Name {
							members = append(members, fmt.Sprintf("%s%s", i.ClientId, i.ClientHost))
						}
					}
				}
				//fmt.Println(members)
				g.Members = members
				topicGroups = append(topicGroups, g)
			}

		}
	}
	//groupStr, _ := json.Marshal(topicGroups)
	//fmt.Println(string(groupStr))
	t.Groups = topicGroups
	return t
}
func (k *Kafka) CommandTopicGroups(tp Topic) {
	//tp := k.TopicInfo(topic)
	//k.TopicGroups(&tp)
	if len(tp.Groups) > 0 {
		tb := table.NewWriter()
		tb.SetOutputMirror(os.Stdout)
		tb.SetTitle(fmt.Sprintf("Topics: %s", tp.Name))
		tb.AppendHeader(table.Row{"Group", "CURRENT-OFFSET Sum", "LOG-END-OFFSE Sum", "LAG Sum", "CONSUMER Num"})

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle(fmt.Sprintf("Topics: %s", tp.Name))
		t.AppendHeader(table.Row{"Group", "partition", "CURRENT-OFFSET", "LOG-END-OFFSET", "LAG", "CONSUMERS"})
		for _, group := range tp.Groups {
			tb.AppendRow([]interface{}{group.Name, tp.PartitionCurrentOffsetSum - group.LagSum,
				tp.PartitionCurrentOffsetSum, group.LagSum, len(group.Members)})

			for _, partition := range tp.Partitions {
				t.AppendRow([]interface{}{group.Name, partition,
					group.PartitionCurrentOffset[partition],
					tp.PartitionCurrentOffset[partition],
					group.Lag[partition], group.Members})
			}
		}
		tb.Render()
		t.Render()
	}
}
func (k *Kafka) CommandListGroups() {
	groups, err := k.admin.ListConsumerGroups()
	if err != nil {
		fmt.Println(err)
	}
	for k, _ := range groups {
		fmt.Println(k)
	}
}
