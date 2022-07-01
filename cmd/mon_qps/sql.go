package main

import "fmt"

func Sql(start int64) EsSqlParam {

	end := start - 5*60*1000

	seconds := EsSqlParam{
		Query:          fmt.Sprintf(`select HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time, host,status  ,count(1) as qps from "nginx-*" where "@timestamp" > %d and "@timestamp" <  %d group by time,status,host   `, end, start),
		FetchSize:      5000,
		TimeZone:       "Asia/Shanghai",
		RequestTimeout: "180s",
		PageTimeout:    "180s",
	}

	return seconds
}
