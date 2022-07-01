package main

import (
	"encoding/json"
	"fmt"
	"git..com/ops-tools/kafkalib"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func VerifyKafkaClient(c *gin.Context) (kafkalib.Kafka, error) {
	var k kafkalib.Kafka
	errs := ""
	kafkaVersion := c.PostForm("kafkaVersion")
	if kafkaVersion == "" {
		errs = fmt.Sprintf("%s, Parameter Lost: kafkaVersion", errs)
	}
	k.Version, err = sarama.ParseKafkaVersion(kafkaVersion)
	if err != nil {
		errs = fmt.Sprintf("%s, Parameter Error: kafkaVersion, %s", errs, err.Error())
	}

	k.Brokers = c.PostFormArray("kafkaBrokers")
	log.Println(k.Brokers)
	if len(k.Brokers) == 0 {
		errs = fmt.Sprintf("%s, Parameter Lost: kafkaBrokers", errs)
	}

	return k, err
}

func MakeKafkaClient(c *gin.Context) (kafkalib.Kafka, error) {

	k, err := VerifyKafkaClient(c)
	if err != nil {
		return k, err
	}
	err = k.ApiInitClient()
	if err != nil {
		return k, err
	}
	return k, nil
}
func MakeKafkaAdminClient(c *gin.Context) (kafkalib.Kafka, error) {
	k, err := VerifyKafkaClient(c)
	if err != nil {
		return k, err
	}
	err = k.ApiInitAdminClient()
	if err != nil {
		return k, err
	}
	return k, nil
}
func v1GetTopics(c *gin.Context) {
	k, err := MakeKafkaClient(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	b, err := json.Marshal(k.ListTopics())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, string(b))
	log.Println(string(b))
}

func v1CreateTopic(c *gin.Context) {
	k, err := MakeKafkaAdminClient(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	var partitionCount int
	if c.PostForm("partition") != "" {
		partitionCount, err = strconv.Atoi(c.PostForm("partition"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Parameter Error: partition %s", err.Error()),
			})
			return
		}
	} else {
		partitionCount = 1
	}
	var replicationFactor int
	//if c.PostForm("replication") != "" {
	//	replicationFactor, err = strconv.Atoi(c.PostForm("replication"))
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"err": fmt.Sprintf("Parameter Error: replication %s", err.Error()),
	//		})
	//		return
	//	}
	//} else {
	//	replicationFactor = 2
	//}
	replicationFactor = 2
	var retentionBytes int
	if c.PostForm("retention.bytes") != "" {
		retentionBytes, err = strconv.Atoi(c.PostForm("retention.bytes"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Parameter Error: retention.bytes %s", err.Error()),
			})
			return
		}
	} else {
		retentionBytes = 21474836480
	}
	var retentionMs int
	if c.PostForm("retention.ms") != "" {
		retentionMs, err = strconv.Atoi(c.PostForm("retention.ms"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Parameter Error: retention.ms %s", err.Error()),
			})
			return
		}
	} else {
		retentionMs = 86400000
	}
	t := kafkalib.Topic{
		Name:              c.PostForm("topic"),
		PartitionCount:    int32(partitionCount),
		ReplicationFactor: int16(replicationFactor),
		RetentionBytes:    retentionBytes,
		RetentionMs:       retentionMs,
	}
	err = k.CreateTopic(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "")
}
