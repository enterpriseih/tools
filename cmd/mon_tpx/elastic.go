package main

import (
	"bytes"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SqlData struct {
	Columns []Columns       `json:"columns"`
	Rows    [][]interface{} `json:"rows,omitempty"`
	Cursor  string          `json:"cursor"`
	Error   Error           `json:"error"`
	Status  int             `json:"status"`
}

type RootCause struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}
type CausedBy struct {
	Type   string      `json:"type"`
	Reason interface{} `json:"reason"`
}
type Error struct {
	RootCause []RootCause `json:"root_cause"`
	Type      string      `json:"type"`
	Reason    string      `json:"reason"`
	CausedBy  CausedBy    `json:"caused_by"`
}

type Columns struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type EsSqlParam struct {
	Query          string `json:"query"`
	FetchSize      int    `json:"fetch_size"`
	TimeZone       string `json:"time_zone"`
	RequestTimeout string `json:"request_timeout"`
	PageTimeout    string `json:"page_timeout"`
}

var es = "http://172.24.116.23:9200"
var sql = new(EsSqlParam)

func CloseCursor(cursor string) {

	esUrl := fmt.Sprintf("%s%s", es, "/_sql/close")
	client := &http.Client{}

	req, err := http.NewRequest("POST", esUrl, strings.NewReader(fmt.Sprintf(`{  "cursor" : "%s"}`, cursor)))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

}
func QueryCursor(cursor string) *SqlData {
	log.Println("query cursor")

	esUrl := fmt.Sprintf("%s%s", es, "/_sql?format=json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", esUrl, strings.NewReader(fmt.Sprintf(`{  "cursor" : "%s","time_zone":"Asia/Shanghai","request_timeout":"%s","page_timeout":"%s"}`, cursor, "900s", "900s")))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(body))
	data := new(SqlData)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	log.Println("cursor end")

	return data
}
func QuerySql(q []byte) *SqlData {
	log.Println("query sql")
	esUrl := fmt.Sprintf("%s%s", es, "/_sql?format=json")

	client := &http.Client{}

	req, err := http.NewRequest("GET", esUrl, bytes.NewReader(q))

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(body))
	data := new(SqlData)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	if data.Error.Type != "" {
		log.Println(data.Error.Reason)
	}
	log.Println("sql end")
	return data
}
