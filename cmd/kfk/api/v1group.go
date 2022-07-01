package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func v2GetGroups(c *gin.Context) {
	cluster := c.Query("ClusterName")
	topic := c.Query("TopicName")

	log.Printf("kafka集群: %s Topic: %s", cluster, topic)
	if cluster != "" && topic != "" {

		b, err := json.Marshal(QueryTopicGroupList(cluster, topic))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))
	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster Name! ")
	}
}

func v2QueryAllGroups(c *gin.Context) {
	cluster := c.Query("ClusterName")

	log.Printf("kafka集群: %s", cluster)
	if cluster != "" {

		b, err := json.Marshal(QueryGroupList(cluster))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))
	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster Name! ")
	}
}

func v2GetAllGroups(c *gin.Context) {
	cluster := c.PostForm("ClusterName")

	log.Printf("kafka集群: %s", cluster)
	if cluster != "" {

		b, err := json.Marshal(QueryGroupList(cluster))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))
	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster Name! ")
	}
}

func v2GroupStatus(c *gin.Context) {
	cluster := c.PostForm("ClusterName")
	group := c.PostForm("GroupName")
	if cluster != "" && group != "" {
		b, err := json.Marshal(QyeryGroupTopics(group, cluster))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))

	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster/Group Name! ")
	}
}
func v2QueryGroupStatus(c *gin.Context) {
	cluster := c.Query("ClusterName")
	group := c.Query("GroupName")
	if cluster != "" && group != "" {
		b, err := json.Marshal(QyeryGroupTopics(group, cluster))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(b))
		c.String(http.StatusOK, string(b))

	} else {
		c.String(http.StatusInternalServerError, "Error Kafka Cluster/Group Name! ")
	}
}
