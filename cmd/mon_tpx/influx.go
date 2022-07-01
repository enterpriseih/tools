package main

import (
	"fmt"
	"time"
)

func (c *NewClient) WriteInflux(data [][]interface{}) {

	//client := influxdb2.NewClientWithOptions("http://172.24.115.219:8086", "",
	//	influxdb2.DefaultOptions().SetBatchSize(500))

	// Get the blocking write client
	// Supply a string in the form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
	// Org name is not used
	fmt.Println(data)

	writeAPI := c.Influx.WriteAPI("", "ngx/rp_ngx")
	// create point using full params constructor

	for _, i := range data {
		t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i[0]))

		p := influxdb2.NewPoint("ngx_tpx",
			map[string]string{
				"server_name": fmt.Sprintf("%s", i[1]),
			},
			map[string]interface{}{
				"avg":  i[2],
				"tp99": i[3],
				"tp95": i[4],
				"tp90": i[5],
				"tp50": i[6],
			},

			t)
		// Write data
		writeAPI.WritePoint(p)

	}
	writeAPI.Flush()
	c.Influx.Close()

}

func (c *NewClient) SaveToInflux(b []byte) {
	var cursor string
	d := QuerySql(b)
	c.WriteInflux(d.Rows)

	if d.Cursor == "" {
	} else {
		cursor = d.Cursor
		for {
			b := QueryCursor(cursor)
			c.WriteInflux(b.Rows)

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
