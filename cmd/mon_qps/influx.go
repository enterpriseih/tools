package main

import (
	"fmt"

	"time"
)

func (c *NewClient) WriteInflux(data [][]interface{}) {

	// Get the blocking write client
	// Supply a string in the form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
	// Org name is not used
	writeAPI := c.Influx.WriteAPI("", "ngx/rp_ngx")
	// create point using full params constructor
	for _, i := range data {
		t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i[0]))

		p := influxdb2.NewPoint("ngx_qps",
			map[string]string{
				"server_name": fmt.Sprintf("%s", i[1]),
				"status":      fmt.Sprintf("%.0f", i[2]),
			},
			map[string]interface{}{"qps": i[3]},
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
