package main

import (
	"bytes"

	"encoding/json"
	"flag"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
var out string
var sql = new(EsSqlParam)

func init() {
	//flag.StringVar(&sql.TimeZone, "tz", "Asia/Shanghai", "-tz Asia/Shanghai: Time-zone in ISO 8601 used for executing the query on the server.\n https://docs.oracle.com/javase/8/docs/api/java/time/ZoneId.html")
	sql.TimeZone = "Asia/Shanghai"
	flag.IntVar(&sql.FetchSize, "ps", 1000, "-ps 1000: The maximum number of rows (or entries) to return in one response")
	flag.StringVar(&sql.Query, "sql", "show tables", "-sql 'select \"@timestamp\",server_name,status from \"nginx-2020.07.15\" limit 11'\n SQL query to execute.")
	flag.StringVar(&sql.RequestTimeout, "rt", "90s", "-rt 90s: The timeout before the request fails.")
	flag.StringVar(&sql.PageTimeout, "pt", "45s", "-pt 45: The timeout before a pagination request fails.")
	flag.StringVar(&out, "o", "", "-o out.csv :  output file")
	flag.Parse()
}

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
	fmt.Println(string(body))

}
func QueryCursor(cursor string) *SqlData {
	esUrl := fmt.Sprintf("%s%s", es, "/_sql?format=json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", esUrl, strings.NewReader(fmt.Sprintf(`{  "cursor" : "%s","time_zone":"Asia/Shanghai","request_timeout":"%s","page_timeout":"%s"}`, cursor, sql.RequestTimeout, sql.PageTimeout)))
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
	//fmt.Println(data.Cursor)

	//CloseCursor(data.Cursor)
	return data
}
func QuerySql(q []byte) *SqlData {
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
		fmt.Println(string(body))
		panic(data.Error.Reason)
	}
	return data
}

func (s *SqlData) UnmarshalJSON(data []byte) error {
	dat := &struct {
		Columns []Columns       `json:"columns"`
		Rows    [][]interface{} `json:"rows"`
		Cursor  string          `json:"cursor"`
		Error   Error           `json:"error"`
		Status  int             `json:"status"`
	}{}

	if err := json.Unmarshal(data, &dat); err != nil {
		return err
	}

	s.Columns = dat.Columns
	s.Cursor = dat.Cursor
	s.Error = dat.Error
	s.Status = dat.Status

	rows := [][]interface{}{}

	for _, row := range dat.Rows {
		rr := []interface{}{}
		for _, cell := range row {
			if cell == nil {
				rr = append(rr, "")
			} else {
				rr = append(rr, cell)
			}
		}
		rows = append(rows, rr)
	}

	s.Rows = rows
	return nil
}

func Terminal(b []byte) {
	num := 0

	var cursor string

	d := QuerySql(b)

	columns := table.Row{}

	for _, i := range d.Columns {
		columns = append(columns, i.Name)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(columns)
	for _, j := range d.Rows {
		t.AppendRow(j)
		num += 1

		if num > 10000 {
			CloseCursor(cursor)
			break
		}
	}
	if d.Cursor == "" {
	} else {
		cursor = d.Cursor
		for {
			fmt.Printf(".")

			b := QueryCursor(cursor)

			for _, k := range b.Rows {
				t.AppendRow(k)
				num += 1
				if num > 9999 {
					CloseCursor(cursor)
					break
				}
			}

			if num > 9999 {
				CloseCursor(cursor)
				break
			}
			if b.Cursor == "" {
				CloseCursor(cursor)
				break
			}
			if b.Cursor != cursor {
				cursor = b.Cursor
			}
		}
	}
	t.Render()
	log.Printf("Terminal only show 10000, number of lines: %d\n", num)
}

func SaveToFile(b []byte) {
	var cursor string

	d := QuerySql(b)

	columns := table.Row{}

	for _, i := range d.Columns {
		columns = append(columns, i.Name)
	}

	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t := table.NewWriter()
	t.SetOutputMirror(f)
	t.AppendHeader(columns)
	for _, j := range d.Rows {
		t.AppendRow(j)
	}
	t.RenderCSV()
	t.ResetHeaders()
	t.ResetRows()

	if d.Cursor == "" {
	} else {
		cursor = d.Cursor
		for {
			//fmt.Println(cursor)
			fmt.Printf(".")
			b := QueryCursor(cursor)

			for _, k := range b.Rows {
				t.AppendRow(k)
				//fmt.Printf("%#v\n", k)
			}
			t.RenderCSV()
			//t.ResetHeaders()
			t.ResetRows()

			if b.Cursor == "" {
				CloseCursor(cursor)
				break
			}
			if b.Cursor != cursor {
				cursor = b.Cursor
			}
		}
	}
}
func main() {

	b, err := json.Marshal(sql)
	if err != nil {
		fmt.Println(err)
	}

	if out != "" {
		SaveToFile(b)
	} else {
		Terminal(b)
	}

}
