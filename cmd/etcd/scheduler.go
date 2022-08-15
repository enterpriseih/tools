package main

func schedulerService() string {
	return `
cat > /etc/systemd/system/kube-scheduler.service << EOF

[Unit]
Description=Kubernetes Scheduler
Documentation=https://github.com/kubernetes/kubernetes
After=network.target

[Service]
ExecStart=/usr/local/bin/kube-scheduler --authentication-kubeconfig=/etc/kubernetes/scheduler.conf \
  --authorization-kubeconfig=/etc/kubernetes/scheduler.conf \
  --bind-address=127.0.0.1 --kubeconfig=/etc/kubernetes/scheduler.conf --leader-elect=true

Restart=always
RestartSec=10s

[Install]
WantedBy=multi-user.target

EOF
`
}
