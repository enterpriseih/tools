package main

func kubeletService() string {
	return `
cat > /etc/systemd/system/kubelet.service << EOF

[Unit]
Description=Kubernetes Kubelet
Documentation=https://github.com/kubernetes/kubernetes
After=crio.service
Requires=crio.service

[Service]
ExecStart=/usr/local/bin/kubelet --cluster-dns=10.96.0.10  --cluster-domain=cluster.local \
  --kubeconfig=/etc/kubernetes/kubelet.conf --cgroup-driver=systemd --container-runtime=remote  \
  --container-runtime-endpoint=unix:///var/run/crio/crio.sock 


[Install]
WantedBy=multi-user.target
EOF
`
}
