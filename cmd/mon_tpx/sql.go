package main

import "fmt"

func Sql(start int64) EsSqlParam {

	end := start - 5*60*1000

	seconds := EsSqlParam{
		Query: fmt.Sprintf(`
SELECT HISTOGRAM("@timestamp", INTERVAL 5 SECONDS) AS time,host, 
ROUND(AVG(request_time),3) AS  rtavg,
ROUND(PERCENTILE(request_time,99),3) AS tp99 ,
ROUND(PERCENTILE(request_time,95),3) AS tp95 ,
ROUND(PERCENTILE(request_time,90),3) AS tp90 ,
ROUND(PERCENTILE(request_time,50),3) AS tp50 
FROM "nginx-*" where "@timestamp" > %d and "@timestamp" < %d 
GROUP BY host,time
  `, end, start),
		FetchSize:      5000,
		TimeZone:       "Asia/Shanghai",
		RequestTimeout: "180s",
		PageTimeout:    "180s",
	}

	return seconds
}
