package main

type K8SCluster struct {
	Etcd      []Etcd    `yaml:"etcd"`
	APIServer APIServer `yaml:"apiServer"`
	EtcdUrl   string    `yaml:"etcdUrl"`
}
type Etcd struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}
type Networking struct {
	DNSDomain     string `yaml:"dnsDomain"`
	ServiceSubnet string `yaml:"serviceSubnet"`
	PodSubnet     string `yaml:"podSubnet"`
}
type APIServer struct {
	Port       int        `yaml:"port"`
	Networking Networking `yaml:"networking"`
	Host       []string   `yaml:"host"`
}
