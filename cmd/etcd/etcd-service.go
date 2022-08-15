package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func (k *K8SCluster) EtcdService() {

	for _, i := range k.Etcd {
		ts, _ := template.New("etcd").Parse(fmt.Sprintf(etcdServiceTmpl(), k.etcdPeerUrl))
		_ = ts.Execute(os.Stdout, i)
		fmt.Println()
	}
}
func (k *K8SCluster) etcdUrl() string {
	var buf bytes.Buffer
	t, _ := template.New("EtcdUrl").Parse(`{{ range .Etcd -}} https://{{ .IP }}:2379, {{- end }}`)
	_ = t.Execute(&buf, k)
	k.EtcdUrl = strings.TrimRight(buf.String(), ",")
	return k.EtcdUrl
}

func (k *K8SCluster) etcdPeerUrl() string {
	var buf bytes.Buffer
	t, _ := template.New("etcdPeerUrl").Parse(`{{ range .Etcd -}} {{ .Name}}=https://{{ .IP }}:2380, {{- end }}`)
	_ = t.Execute(&buf, k)
	return strings.TrimRight(buf.String(), ",")
}

func etcdServiceTmpl() string {
	return `cat > /etc/systemd/system/etcd.service <<EOF
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
}
