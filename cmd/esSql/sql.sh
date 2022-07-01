#./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 MINUTES) AS time, server_name,count(1)  from "nginx-2020.07.16" where "@timestamp" > 1594895250171 group by  time,server_name '
#./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 HOURS) AS time, server_name,count(1)  from "nginx-2020.07.16" where "@timestamp" > 1594895250171 group by  time,server_name '
./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time, server_name,count(1)  from "nginx-2020.07.16" where "@timestamp" > 1594895250171 and "@timestamp" < 1594895350171  group by  time,server_name ' -o a.out

./esSql -sql ' select  HISTOGRAM("@timestamp", INTERVAL 1 HOURS) AS time, server_name,count(1)  from "nginx-2020.07.18" group by  time,server_name  '

./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 HOURS) AS time ,server_name,request_uri  from "nginx-2020.07.21" where "@timestamp" > 1595287800000 and "@timestamp" < 1595291400000 and request_time > 1 group by  time,server_name'

./esSql -sql 'select   *  from "nginx-2020.07.21" where "@timestamp" > 1595287800000 and "@timestamp" < 1595291400000 and request_time > 1 '

./esSql -sql 'select server_name,request_uri from "nginx-2020.07.26" where "@timestamp" > 1595703600000 and "@timestamp" < 1595707200000  and request_uri like '\'/services/v3/search/humanSearch%\''  limit 10 '

./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time,server_name,count(1) as QPS from "nginx-2020.07.26" where "@timestamp" > 1595703600000 and "@timestamp" < 1595707200000  and request_uri like '\'/services/v3/search/humanSearch%\''  group by time,server_name ' -o nginx-2020.07.26.csv

./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time,server_name,count(1) as QPS from "nginx-2020.07.27" where "@timestamp" > 1595790000000 and "@timestamp" < 1595793600000 and request_uri like '\'/services/v3/search/humanSearch%\''  group by time,server_name ' -o nginx-2020.07.27.csv

./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time,server_name,count(1) as QPS from "nginx-2020.07.28" where "@timestamp" > 1595917800000 and "@timestamp" < 1595919600000    group by time,server_name ' -o nginx-2020.07.28.csv

./esSql -sql 'select  HISTOGRAM("@timestamp", INTERVAL 1 SECONDS) AS time,host,remote_addr,"geoip.country_name" as country,"geoip.region_name" as region ,count(1) as QPS from ngx where "@timestamp" >  1597052700000 and  "@timestamp"  < 1597053000000 and host = '\'zhuanli.com\'' group by time, host,remote_addr,country,region  ' -o out.csv

