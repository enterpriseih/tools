package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
)

func main() {
	cluster := `
etcd:
  - name: etcd0
    ip: 10.127.254.252
  - name: etcd1
    ip: 10.127.254.251
  - name: etcd2
    ip: 10.127.254.250
apiServer:
  port: 6443
  networking:
    dnsDomain: cluster.local
    serviceSubnet: 10.96.0.0/16
    podSubnet: 10.16.0.0/16
  host:
    - 10.127.254.252
    - 10.127.254.251
    - 10.127.254.250

`
	k8sCluster := K8SCluster{}

	err := yaml.Unmarshal([]byte(cluster), &k8sCluster)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	k8sCluster.EtcdService()
	k8sCluster.ApiService()
	etcdCA()
	k8sCA()
	frontCA()
	fmt.Println(controllerManagerService())
	fmt.Println(schedulerService())
	fmt.Println(kubeletService())
}
