package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"strings"
)

type KafkaClusters []KafkaCluster

type KafkaCluster struct {
	Name    string
	Version string
	Brokers []string
}

func v1Clusters(c *gin.Context) {
	client, err := consul.NewClient(&consul.Config{
		Address:    "consul.tyc.local",
		Scheme:     "http",
		Datacenter: "",
		Transport:  nil,
		HttpClient: nil,
		HttpAuth:   nil,
		WaitTime:   0,
		Token:      "",
		TokenFile:  "",
		Namespace:  "",
		TLSConfig:  consul.TLSConfig{},
	})

	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	keys, _, err := kv.List("/Kafka/", nil)
	if err != nil {
		panic(err)
	}
	var kcs KafkaClusters
	for _, i := range keys {
		sp := strings.Split(i.Key, "/")
		if len(sp) == 2 {
			if sp[1] != "" {
				var kc KafkaCluster

				err := json.Unmarshal(i.Value, &kc)
				if err != nil {
					log.Println(err)
				} else {
					kc.Name = sp[1]
					kcs = append(kcs, kc)
				}
			}
		}
	}
	data, _ := json.Marshal(kcs)
	log.Println(string(data))
	c.String(http.StatusOK, string(data))
}
func v1GetClusters(c *gin.Context) {
	client, err := consul.NewClient(&consul.Config{
		Address:    "consul.tyc.local",
		Scheme:     "http",
		Datacenter: "",
		Transport:  nil,
		HttpClient: nil,
		HttpAuth:   nil,
		WaitTime:   0,
		Token:      "",
		TokenFile:  "",
		Namespace:  "",
		TLSConfig:  consul.TLSConfig{},
	})

	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	keys, _, err := kv.List("/Kafka/", nil)
	if err != nil {
		panic(err)
	}
	var kcs KafkaClusters
	for _, i := range keys {
		sp := strings.Split(i.Key, "/")
		if len(sp) == 2 {
			if sp[1] != "" {
				var kc KafkaCluster

				err := json.Unmarshal(i.Value, &kc)
				if err != nil {
					log.Println(err)
				} else {
					kc.Name = sp[1]
					kcs = append(kcs, kc)
				}
			}
		}
	}
	data, _ := json.Marshal(kcs)
	log.Println(string(data))
	c.String(http.StatusOK, string(data))
}
