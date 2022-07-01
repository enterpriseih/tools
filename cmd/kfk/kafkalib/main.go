package kafkalib

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var err error

type KafkaInterface interface {
	TopicInfo(string) Topic
	ListTopics() *Topics
	TopicGroups(t *Topic) *Topic
	TopicDirs(t *Topic) *Topic
	KafkaCommand
	KafkaAdmin
}
type KafkaCommand interface {
	CommandListTopics()
	//CommandDescribeTopic(string)
	CommandDescribeTopic(Topic)
	CommandTopicGroups(Topic)
	CommandTopicDirs(Topic)
}
type KafkaAdmin interface {
	CreateTopic(*Topic) error
	DeleteTopic(string)
	DeleteGroup(string)
	CommandListGroups()
}
type Kafka struct {
	Version    sarama.KafkaVersion
	Brokers    []string
	BrokersStr string
	client     sarama.Client
	admin      sarama.ClusterAdmin
}

//var (
//	kafka Kafka
//	err   error
//	topic string
//	list  bool
//)
//
//func init() {
//
//	flag.StringVar(&topic, "t", "", "describe kafka topic: -t <topic>")
//	flag.BoolVar(&list, "l", false, "list kafka topics: true")
//
//	flag.Parse()
//	kafka.version = os.Getenv("KAFKA_VERSION")
//	kafka.BrokersStr = os.Getenv("KAFKA")
//
//	// For Test
//	//kafka.version = "1.0.0"
//	//kafka.BrokersStr = "10.2.16.8:9092"
//
//	if kafka.version == "" || kafka.BrokersStr == "" {
//		fmt.Println("export KAFKA_VERSION=1.0.0")
//		fmt.Println("export KAFKA=127.0.0.1:9092")
//	}
//	kafka.Brokers = strings.Split(kafka.BrokersStr, ",")
//	kafka.Version, err = sarama.ParseKafkaVersion(kafka.version)
//	if err != nil {
//		fmt.Println("Error Kafka Version")
//		panic(err)
//	}
//}

func (k *Kafka) InitClient() {
	config := sarama.NewConfig()
	config.Version = k.Version
	k.client, err = sarama.NewClient(k.Brokers, config)
	if err != nil {
		fmt.Println("Error Init Kafka Client")
		panic(err)
	}
}
func (k *Kafka) InitAdminClient() {
	config := sarama.NewConfig()
	config.Version = k.Version
	k.admin, err = sarama.NewClusterAdmin(k.Brokers, config)
	if err != nil {
		fmt.Println("Error Init Kafka Client")
		panic(err)
	}
}

func (k *Kafka) ApiInitClient() error {
	config := sarama.NewConfig()
	config.Version = k.Version
	k.client, err = sarama.NewClient(k.Brokers, config)
	if err != nil {
		//fmt.Println("Error Init Kafka Client")
		//panic(err)
		return err

	}
	return nil
}
func (k *Kafka) ApiInitAdminClient() error {
	config := sarama.NewConfig()
	config.Version = k.Version
	k.admin, err = sarama.NewClusterAdmin(k.Brokers, config)
	if err != nil {
		return err
	}
	return nil
}

//func main() {
//
//	kafka.InitClient()
//
//	var ki KafkaInterface = &kafka
//	// Topics
//	if list {
//		ki.CommandListTopics()
//	}
//	//topic = "example"
//	if topic != "" {
//		tp := ki.TopicInfo(topic)
//		// 加载Topic 数据目录大小
//		ki.TopicDirs(&tp)
//
//		ki.CommandDescribeTopic(tp)
//		ki.CommandTopicDirs(tp)
//		// 加载Consumer Group数据
//		ki.TopicGroups(&tp)
//		ki.CommandTopicGroups(tp)
//	}
//
//}
