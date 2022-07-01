package main

import (
	"code.cloudfoundry.org/bytefmt"
	"context"
	"fmt"
	promApi "github.com/prometheus/client_golang/api"
	promV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
	"log"
	"os"
	"time"
)

type Topic struct {
	Partitions             []string
	TopicName              string
	PartitionNum           int
	DataSize               string
	PartitionCurrentOffset promModel.SampleValue
	PartitionLave          promModel.SampleValue
	MessagePerSecond       string
	MessagePerMinute       string
	HaveConsumerGroup      string
}
type Kafka struct {
	ClusterName string
	Topics      []Topic
}
type TopicConsumerGroup []ConsumerGroup
type ConsumerGroup struct {
	GroupName        promModel.LabelValue
	ConsumerGroupLag promModel.SampleValue
}

type GroupTopics []GroupTopic

type GroupTopic struct {
	TopicName     string
	CurrentOffset promModel.SampleValue
	GroupLag      promModel.SampleValue
}

func PromClient() promV1.API {
	client, err := promApi.NewClient(promApi.Config{
		Address: "http://172.24.115.201:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	v1api := promV1.NewAPI(client)
	return v1api
}

func QueryFloat(query string) promModel.Value {

	v1api := PromClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, warnings, err := v1api.Query(ctx, query, time.Now())

	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	return result
	//fmt.Printf("%#v\n", result)

}

func (k *Kafka) QueryKafkaDataSize() map[string]string {
	query := fmt.Sprintf(`sum(kafka_log_log_size{cluster="%s"}) by (topic)`, k.ClusterName)
	topicSize := make(map[string]string)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			topicSize[string(v.Metric["topic"])] = bytefmt.ByteSize(uint64(v.Value) * bytefmt.BYTE)
		}
	}
	return topicSize

}
func (k *Kafka) QueryCurrentOffset() map[string]promModel.SampleValue {
	query := fmt.Sprintf(`sum(kafka_topic_partition_current_offset{cluster='%s'}) by (topic) `, k.ClusterName)

	partitionCurrentOffsets := make(map[string]promModel.SampleValue)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			partitionCurrentOffsets[string(v.Metric["topic"])] = v.Value
		}
	}
	return partitionCurrentOffsets
}

func (k *Kafka) QueryOldestOffset() map[string]promModel.SampleValue {
	query := fmt.Sprintf(`sum(kafka_topic_partition_oldest_offset{cluster='%s'}) by (topic) `, k.ClusterName)

	partitionOldestOffset := make(map[string]promModel.SampleValue)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			partitionOldestOffset[string(v.Metric["topic"])] = v.Value
		}
	}
	return partitionOldestOffset
}

func (k *Kafka) QueryMessageInPerSecond() map[string]promModel.SampleValue {
	query := fmt.Sprintf(`sum(rate(kafka_topic_partition_current_offset{cluster="%s"}[1m])) by (topic) `, k.ClusterName)

	messagePerSecond := make(map[string]promModel.SampleValue)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			messagePerSecond[string(v.Metric["topic"])] = v.Value
		}
	}
	return messagePerSecond
}

func (k *Kafka) QueryMessageInPerMinute() map[string]promModel.SampleValue {
	query := fmt.Sprintf(`sum(delta(kafka_topic_partition_current_offset{cluster='%s'}[5m])/5) by (topic)`, k.ClusterName)

	messagePerMinute := make(map[string]promModel.SampleValue)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			messagePerMinute[string(v.Metric["topic"])] = v.Value
		}
	}
	return messagePerMinute
}

func (k *Kafka) QueryTopics() {
	v1api := PromClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lbls, warnings, err := v1api.Series(ctx, []string{
		fmt.Sprintf(`kafka_topic_partition_current_offset{job='kafka_exporter',cluster='%s',topic!='__consumer_offsets'}`, k.ClusterName),
	}, time.Now().Add(-time.Hour), time.Now())
	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	topics := make(map[string][]string, len(lbls))
	for _, lbl := range lbls {
		topics[string(lbl["topic"])] = append(topics[string(lbl["topic"])], fmt.Sprintf("%s", lbl["partition"]))
	}
	//log.Println(topics)
	topicSize := k.QueryKafkaDataSize()
	partitionCurrentOffsets := k.QueryCurrentOffset()
	partitionOldestOffsets := k.QueryOldestOffset()
	messagePerSeconds := k.QueryMessageInPerSecond()
	messagePerMinutes := k.QueryMessageInPerMinute()
	haveGroupTopics := QueryHaveGroupTopics(k.ClusterName)
	for topic, partations := range topics {

		var t = new(Topic)
		t.TopicName = topic
		if haveGroupTopics[topic] {
			t.HaveConsumerGroup = "有"
		} else {
			t.HaveConsumerGroup = "无"
		}
		t.Partitions = partations
		t.PartitionNum = len(partations)

		t.PartitionCurrentOffset = partitionCurrentOffsets[topic]
		t.PartitionLave = partitionCurrentOffsets[topic] - partitionOldestOffsets[topic]
		t.MessagePerSecond = fmt.Sprintf("%.2f", messagePerSeconds[topic])
		t.MessagePerMinute = fmt.Sprintf("%.2f", messagePerMinutes[topic])
		t.DataSize = topicSize[topic]

		k.Topics = append(k.Topics, *t)
	}

}
func (k *Kafka) QueryTopicsList() {
	v1api := PromClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lbls, warnings, err := v1api.Series(ctx, []string{
		fmt.Sprintf(`kafka_topic_partition_current_offset{job='kafka_exporter',cluster='%s'}`, k.ClusterName),
	}, time.Now().Add(-time.Hour), time.Now())
	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	topics := make(map[string][]string, len(lbls))
	for _, lbl := range lbls {
		topics[string(lbl["topic"])] = append(topics[string(lbl["topic"])], fmt.Sprintf("%s", lbl["partition"]))
	}
	//log.Println(topics)

	for topic, _ := range topics {
		var t = new(Topic)
		t.TopicName = topic

		k.Topics = append(k.Topics, *t)
	}

}

func QueryHaveGroupTopics(cluster string) map[string]bool {
	query := fmt.Sprintf(`sum(kafka_consumergroup_lag{cluster="%s"}) by ( topic) `, cluster)
	topics := make(map[string]bool)
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			topics[string(v.Metric["topic"])] = true
		}
	}
	log.Println(topics)
	return topics
}

func QueryTopicGroupList(cluster string, topic string) TopicConsumerGroup {

	query := fmt.Sprintf(`sum(kafka_consumergroup_lag{cluster="%s",topic="%s"}) by (consumergroup, topic,cluster) `, cluster, topic)
	var tg TopicConsumerGroup
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["consumergroup"] != "" {
			var consumerGroup ConsumerGroup
			consumerGroup.GroupName = v.Metric["consumergroup"]
			consumerGroup.ConsumerGroupLag = v.Value
			tg = append(tg, consumerGroup)
		}
	}
	return tg

}
func QueryGroupList(cluster string) TopicConsumerGroup {

	query := fmt.Sprintf(`sum(kafka_consumergroup_lag{cluster="%s"}) by (consumergroup, cluster) `, cluster)
	var tg TopicConsumerGroup
	for _, v := range QueryFloat(query).(promModel.Vector) {
		if v.Metric["consumergroup"] != "" {
			var consumerGroup ConsumerGroup
			consumerGroup.GroupName = v.Metric["consumergroup"]
			consumerGroup.ConsumerGroupLag = v.Value
			tg = append(tg, consumerGroup)
		}
	}
	return tg

}

func QyeryGroupTopics(group string, cluster string) GroupTopics {

	var topics = make([]string, 0)

	topicOffset := make(map[string]promModel.SampleValue)

	queryOffset := fmt.Sprintf(`sum (kafka_consumergroup_current_offset{consumergroup="%s",cluster="%s"} )by (topic,consumergroup,cluster)`, group, cluster)
	for _, v := range QueryFloat(queryOffset).(promModel.Vector) {
		if v.Metric["topic"] != "" {

			topicOffset[string(v.Metric["topic"])] = v.Value
			topics = append(topics, string(v.Metric["topic"]))
		}
	}
	topicGroupLen := make(map[string]promModel.SampleValue)

	queryGroupLag := fmt.Sprintf(`sum (kafka_consumergroup_lag{consumergroup="%s",cluster="%s"} )by (topic,consumergroup,cluster)`, group, cluster)
	for _, v := range QueryFloat(queryGroupLag).(promModel.Vector) {
		if v.Metric["topic"] != "" {
			topicGroupLen[string(v.Metric["topic"])] = v.Value
		}
	}

	var gts GroupTopics
	for _, topic := range topics {
		var gt GroupTopic
		gt.TopicName = topic
		gt.CurrentOffset = topicOffset[topic]
		gt.GroupLag = topicGroupLen[topic]
		gts = append(gts, gt)

	}
	return gts
}
