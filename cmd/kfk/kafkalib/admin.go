package kafkalib

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func (k *Kafka) CreateTopic(topic *Topic) error {

	retentionBytes := fmt.Sprintf("%d", topic.RetentionBytes)
	retentionMs := fmt.Sprintf("%d", topic.RetentionMs)

	entries := make(map[string]*string, 2)

	entries["retention.bytes"] = &retentionBytes
	entries["retention.ms"] = &retentionMs

	err = k.admin.CreateTopic(topic.Name, &sarama.TopicDetail{
		NumPartitions:     topic.PartitionCount,
		ReplicationFactor: topic.ReplicationFactor,
		ConfigEntries:     entries,
	}, false)
	if err != nil {
		return err
	}
	return err

}

func (k *Kafka) DeleteTopic(topic string) {

	err = k.admin.DeleteTopic(topic)
	if err != nil {
		fmt.Println(err)
	}
}
func (k *Kafka) DeleteGroup(group string) {
	err = k.admin.DeleteConsumerGroup(group)
	if err != nil {
		fmt.Println(err)
	}
}
