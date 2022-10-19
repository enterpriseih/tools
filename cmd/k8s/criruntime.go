package k8s

import "k8s.io/kubernetes/pkg/kubelet/cri/remote"

func main() {

	remote.NewRemoteRuntimeService(remoteRuntimeEndpoint, kubeCfg.RuntimeRequestTimeout.Duration, kubeDeps.TracerProvider)
}
