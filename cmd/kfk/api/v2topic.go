package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func v2PostTopics(c *gin.Context) {
	cluster := c.PostForm("ClusterName")
	log.Println(cluster)
	if cluster != "" {
		var kafka Kafka
		kafka.ClusterName = cluster
		kafka.QueryTopics()
		b, err := json.Marshal(kafka)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))
	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster Name! ")
	}

}

func v2GetTopics(c *gin.Context) {
	cluster := c.Query("ClusterName")
	log.Println(cluster)
	if cluster != "" {
		var kafka Kafka
		kafka.ClusterName = cluster
		kafka.QueryTopicsList()
		b, err := json.Marshal(kafka.Topics)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))
	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster Name! ")
	}
}
