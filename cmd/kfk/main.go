package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/Ubbo-Sathla/tools/cmd/kfk/kafkalib"
	"os"
	"strings"
)

var (
	kfk         kafkalib.Kafka
	err         error
	topic       string
	deleteTopic string
	deleteGroup string
	listTopics  bool
	listGroups  bool

	kafkaVersion string
)

func init() {

	flag.StringVar(&topic, "t", "", "describe kafka topic: -t <topic>")
	flag.BoolVar(&listTopics, "lt", false, "list kafka topics: true")
	flag.BoolVar(&listGroups, "lg", false, "list kafka topics: true")

	flag.StringVar(&deleteTopic, "dt", "", "delete kafka topic: -dt <topic>")
	flag.StringVar(&deleteGroup, "dg", "", "delete kafka group: -dg <group>")
	flag.Parse()
	kfk.BrokersStr = os.Getenv("KAFKA")
	kafkaVersion = os.Getenv("KAFKA_VERSION")
	// For Test
	//kafkaVersion = "1.0.0"
	//kfk.BrokersStr = "10.2.16.8:9092"

	if kafkaVersion == "" || kfk.BrokersStr == "" {
		fmt.Println("export KAFKA_VERSION=1.0.0")
		fmt.Println("export KAFKA=127.0.0.1:9092")
	}
	kfk.Brokers = strings.Split(kfk.BrokersStr, ",")
	kfk.Version, err = sarama.ParseKafkaVersion(kafkaVersion)
	if err != nil {
		fmt.Println("Error Kafka Version")
		panic(err)
	}
}

func main() {

	//kfk.InitAdminClient()
	//var ki kafkalib.KafkaInterface = &kfk
	//ki.DeleteTopic("ex1s")
	if deleteTopic != "" {
		kfk.InitAdminClient()
		var ki kafkalib.KafkaInterface = &kfk
		ki.DeleteTopic(deleteTopic)
		return
	}

	if deleteGroup != "" {
		kfk.InitAdminClient()
		var ki kafkalib.KafkaInterface = &kfk

		ki.DeleteGroup(deleteGroup)
		return
	}
	if listGroups {
		kfk.InitAdminClient()
		var ki kafkalib.KafkaInterface = &kfk
		ki.CommandListGroups()
		return
	}
	kfk.InitClient()

	var ki kafkalib.KafkaInterface = &kfk
	// Topics
	if listTopics {
		ki.CommandListTopics()
	}

	//topic = "example"
	if topic != "" {
		tp := ki.TopicInfo(topic)
		// 加载Topic 数据目录大小
		ki.TopicDirs(&tp)

		ki.CommandDescribeTopic(tp)
		ki.CommandTopicDirs(tp)
		// 加载Consumer Group数据
		ki.TopicGroups(&tp)
		ki.CommandTopicGroups(tp)
	}

}
