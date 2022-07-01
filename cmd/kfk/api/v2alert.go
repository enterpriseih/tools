package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type Alert struct {
	Cluster string
	Topic   string
	Group   string
	Phone   string
	Lag     string
}

type SilenceAlert struct {
	ID        string     `json:"id"`
	Matchers  []Matchers `json:"matchers"`
	StartsAt  time.Time  `json:"startsAt"`
	EndsAt    time.Time  `json:"endsAt"`
	CreatedBy string     `json:"createdBy"`
	Comment   string     `json:"comment"`
}
type Matchers struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}

func AddAlert(c *gin.Context) {

	cluster := c.PostForm("Cluster")
	topic := c.PostForm("Topic")
	group := c.PostForm("Group")
	lag := c.PostForm("Lag")
	phone := c.PostForm("Phone")

	if cluster != "" && topic != "" && group != "" && lag != "" && phone != "" {

		topicSlice := strings.Split(topic, ",")
		for _, t := range topicSlice {
			key := fmt.Sprintf("PROMETHEUS/KAFKA/%s/%s/%s/%s", cluster, t, group, phone)
			err := ConsulPutKV(key, lag)
			if err != nil {
				log.Println(err)
				c.String(http.StatusInternalServerError, err.Error())
				break
			}
		}
		c.String(http.StatusOK, "")
	} else {
		c.String(http.StatusInternalServerError, "Error params")
	}

}

func ListAlerts(c *gin.Context) {
	phone := "18345463525"
	client := ConsulClient()
	kv := client.KV()
	keys, _, _ := kv.List("/PROMETHEUS/KAFKA/", nil)
	log.Println(keys)
	var alerts = make([]Alert, 0)
	for _, k := range keys {
		log.Println(k.Key)
		log.Println(k.Value)
		item := strings.Split(k.Key, "/")
		if len(item) == 6 {
			log.Println(item[5])
			if item[5] == phone {
				var alert Alert
				alert.Cluster = item[2]
				alert.Topic = item[3]
				alert.Group = item[4]
				alert.Phone = item[5]
				alert.Lag = string(k.Value)
				alerts = append(alerts, alert)
			}
		}
	}
	b, err := json.Marshal(alerts)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, string(b))
}

func DeleteAlert(c *gin.Context) {
	alert := Alert{
		Cluster: c.PostForm("Cluster"),
		Topic:   c.PostForm("Topic"),
		Group:   c.PostForm("Group"),
		Phone:   c.PostForm("Phone"),
		Lag:     c.PostForm("Lag"),
	}

	if alert.Cluster != "" && alert.Topic != "" && alert.Group != "" && alert.Lag != "" && alert.Phone != "" {
		key := fmt.Sprintf("PROMETHEUS/KAFKA/%s/%s/%s/%s", alert.Cluster, alert.Topic, alert.Group, alert.Phone)
		err := ConsulDelKV(key)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())

		}
		c.String(http.StatusOK, "")
	} else {
		c.String(http.StatusInternalServerError, "Error params")
	}
}

func DoSilenceAlert(c *gin.Context) {
	s := SilenceAlert{
		ID:        "",
		Matchers:  nil,
		StartsAt:  time.Time{},
		EndsAt:    time.Time{},
		CreatedBy: "",
		Comment:   "",
	}
	s.Matchers = append(s.Matchers, Matchers{
		Name:    "cluster",
		Value:   c.PostForm(""),
		IsRegex: false,
	})

	s.Matchers = append(s.Matchers, Matchers{
		Name:    "consumergroup",
		Value:   "",
		IsRegex: false,
	})

	s.Matchers = append(s.Matchers, Matchers{
		Name:    "topic",
		Value:   "",
		IsRegex: false,
	})
	s.Matchers = append(s.Matchers, Matchers{
		Name:    "phone",
		Value:   "",
		IsRegex: false,
	})
	log.Println(s)
	c.String(http.StatusOK, "")

}
