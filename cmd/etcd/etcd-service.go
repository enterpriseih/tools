package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
	"text/template"
)

type EtcdCluster struct {
	Etcd []Etcd `yaml:"etcd"`
}
type Etcd struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}

func main() {
	cluster := `
etcd:
  - name: etcd0
    ip: 10.127.254.252
  - name: etcd1
    ip: 10.127.254.251
  - name: etcd2
    ip: 10.127.254.250
`
	etcdCluster := EtcdCluster{}

	err := yaml.Unmarshal([]byte(cluster), &etcdCluster)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	str := `cat > /etc/systemd/system/etcd.service <<EOF
[Unit]
Description=etcd
Documentation=https://github.com/coreos/etcd

[Service]
Type=notify
Restart=always
RestartSec=5s
LimitNOFILE=40000
TimeoutStartSec=0
# existing
ExecStart=/usr/local/bin/etcd --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt \
  --client-cert-auth=true  --cert-file=/etc/kubernetes/pki/etcd/server.crt  --key-file=/etc/kubernetes/pki/etcd/server.key \
  --peer-client-cert-auth=true  --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt  --peer-key-file=/etc/kubernetes/pki/etcd/peer.key \
  --data-dir=/var/lib/etcd --snapshot-count=10000  --listen-metrics-urls=http://127.0.0.1:2381 \
  --listen-peer-urls=https://{{ .IP }}:2380   --listen-client-urls=https://127.0.0.1:2379,https://{{ .IP }}:2379 \
  --name={{ .Name }}  --initial-cluster-state=new \
  --advertise-client-urls=https://{{ .IP }}:2379  --initial-advertise-peer-urls=https://{{ .IP }}:2380 \
  --initial-cluster=%s

[Install]
WantedBy=multi-user.target
EOF
`
	var buf bytes.Buffer

	t, _ := template.New("cluster").Parse(`{{ range .Etcd -}} {{ .Name}}=https://{{ .IP }}:2380, {{- end }}`)
	_ = t.Execute(&buf, etcdCluster)
	initialCluster := strings.TrimRight(buf.String(), ",")

	for _, i := range etcdCluster.Etcd {
		ts, _ := template.New("etcd").Parse(fmt.Sprintf(str, initialCluster))
		_ = ts.Execute(os.Stdout, i)
		fmt.Println()
	}

}
