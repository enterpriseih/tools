package main

import (
	"fmt"
	"os"
	"text/template"
)

func (k *K8SCluster) ApiService() {
	k.etcdUrl()

	for _, i := range k.APIServer.Host {
		ts, _ := template.New("ApiService").Parse(fmt.Sprintf(apiServiceTmpl(), i))
		_ = ts.Execute(os.Stdout, k)
		fmt.Println()
	}
}

func apiServiceTmpl() string {
	return `cat > /etc/systemd/system/kube-apiserver.service <<EOF
[Unit]
Description=Kubernetes API Server
Documentation=https://github.com/kubernetes/kubernetes
After=network.target

[Service]
ExecStart=/usr/local/bin/kube-apiserver  --advertise-address=%s --secure-port=6443  --etcd-servers={{ .EtcdUrl }}  \
  --service-cluster-ip-range={{ .APIServer.Networking.ServiceSubnet }} --service-account-issuer=https://kubernetes.default.svc.{{ .APIServer.Networking.DNSDomain }} \
  --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key \
  --allow-privileged=true --authorization-mode=Node,RBAC --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true \
  --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key \
  --client-ca-file=/etc/kubernetes/pki/ca.crt  --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key \
  --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key \
  --requestheader-allowed-names=front-proxy-client  --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt \
  --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key \
  --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group  --requestheader-username-headers=X-Remote-User  \
  --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname  

Restart=on-failure
RestartSec=10s
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF
`
}
