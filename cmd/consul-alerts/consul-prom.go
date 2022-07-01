package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

type Alerts struct {
	Group string `json:"group"`
	Alert string `json:"alert"`
	Rule  Rule   `json:"rule"`
}
type Labels struct {
	Severity string `json:"severity"`
}
type Annotations struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}
type Rule struct {
	Expr        string      `json:"expr"`
	For         string      `json:"for"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
}

func Put(node string, key string, value []byte) {
	address := fmt.Sprintf("%s:%d", node, 8500)
	config := &api.Config{
		Address:    address,
		Scheme:     "http",
		Datacenter: "dc1",
	}
	// Get a new client
	client, err := api.NewClient(config)

	if err != nil {
		panic(err)
	}
	kv := client.KV()
	p := &api.KVPair{
		Key:   key,
		Value: value,
	}
	if _, err := kv.Put(p, nil); err != nil {
		fmt.Println(err)
	}

}

func main() {
	consulNode := "172.25.47.74"
	//KV(consulNode)
	baseKey := "PROMETHEUS/GROUPS"

	consulTemplate := `
/ # cat /consul-template/config/config.hcl
consul {

  address = "172.25.47.74:8500"


  retry {
    enabled = true

    attempts = 12

    backoff = "250ms"

    max_backoff = "1m"
  }
}

reload_signal = "SIGHUP"

kill_signal = "SIGINT"

max_stale = "10m"

log_level = "warn"

pid_file = "/tmp/consul-template.pid"

wait {
  min = "5s"
  max = "10s"
}

deduplicate {
  enabled = true

  prefix = "consul-template/dedup/"
}

template {
  source = "/consul-template/config/prometheus.tpl"

  destination = "/consul-template/data/alerts.yml"

  create_dest_dirs = true

  command = ""

  command_timeout = "60s"

  error_on_missing_key = false

  perms = 0644

  backup = true

  left_delimiter  = "{{"
  right_delimiter = "}}"

  wait {
    min = "2s"
    max = "10s"
  }
}
/ # cat /consul-template/config/prometheus.tpl
groups:

{{ range $group, $pairs := tree "PROMETHEUS/GROUPS" | byKey }}
- name: {{ $group }}
  rules:
  {{ range $alert, $v :=  $pairs }}
  - alert: {{ $alert }}
{{ $v.Value | indent 4 }}
  {{ end }}
{{ end }}
`
	fmt.Println(consulTemplate)
	a := `
{
  "group": "group3",
  "alert": "alert1",
  "rule": {
    "expr": "up{job=\"cadvisor\"} == 0",
    "for": "5m",
    "labels": {
      "severity": "page"
    },
    "annotations": {
      "summary": "Docker Service {{ $labels.instance }} down",
      "description": "{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 5 minutes."
    }
  }
}`
	var data Alerts
	json.Unmarshal([]byte(a), &data)
	fmt.Println(data)
	out, _ := yaml.Marshal(data.Rule)
	fmt.Println(string(out))
	alertKey := fmt.Sprintf("%s/%s/%s", baseKey, data.Group, data.Alert)
	fmt.Println(alertKey)
	Put(consulNode, alertKey, out)
}
