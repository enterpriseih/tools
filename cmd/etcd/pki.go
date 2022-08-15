package main

import "fmt"

var caTmpl string = `
[req]
req_extensions = v3_ca
distinguished_name = dn

[dn]
CN=%s 

[v3_ca]
keyUsage = critical,digitalSignature,keyEncipherment,keyCertSign
basicConstraints = critical,CA:TRUE
subjectKeyIdentifier=hash
`

func etcdCA() {
	fmt.Printf(caTmpl, "etcd-ca")
}

func k8sCA() {
	fmt.Printf(caTmpl, "kubernetes")
}

func frontCA() {
	fmt.Printf(caTmpl, "front-proxy-ca")
}
