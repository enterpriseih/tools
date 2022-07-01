package main

import (
	"encoding/json"
	"flag"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"log"
	"time"
)

var begin int64

func init() {

	flag.Int64Var(&begin, "b", 0, "")
	if begin == 0 {
		begin = time.Now().Unix()*1000 - 30*60*1000
	}
	flag.Parse()
}

type NewClient struct {
	Influx influxdb2.Client
}

func main() {
	var start int64
	start = begin
	c := NewClient{Influx: influxdb2.NewClientWithOptions("http://172.24.115.219:8086", "",
		influxdb2.DefaultOptions().SetBatchSize(500))}
	for {

		if start > time.Now().Unix()*1000-30*1000 {
			start = start - 10*60*1000
			time.Sleep(time.Second * 10)
		} else {
			start = start + 4*60*1000

		}
		sql := Sql(start)
		b, err := json.Marshal(sql)
		if err != nil {
			fmt.Println(err)
		}

		c.SaveToInflux(b)
		log.Printf("%s: %d\n", time.Unix(start/1000, 0), start)
	}
}
