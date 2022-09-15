package k8s

import "testing"

func TestDecode(t *testing.T) {
	Decode(listData)
}

const listData = `
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "65e0015548312573d2853bab7045e6d4762f6ea71fe4a3012bbcac8df5eb77d7",
                    "cni.projectcalico.org/podIP": "192.168.219.199/32",
                    "cni.projectcalico.org/podIPs": "192.168.219.199/32"
                },
                "creationTimestamp": "2022-09-15T08:14:37Z",
                "generateName": "nginx-deployment-9456bbbf9-",
                "labels": {
                    "app": "nginx",
                    "pod-template-hash": "9456bbbf9"
                },
                "name": "nginx-deployment-9456bbbf9-476ff",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "nginx-deployment-9456bbbf9",
                        "uid": "326f8493-fe44-41ca-86f0-bfc740bf2dab"
                    }
                ],
                "resourceVersion": "116484",
                "uid": "9be2b7d5-83e2-4949-88e2-1e218bf4fd0b"
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx:1.14.2",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "nginx",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-vj9tg",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "yaoshicheng-kubernetes-2.cloud.onecloud.io",
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "default",
                "serviceAccountName": "default",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "kube-api-access-vj9tg",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:50Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:50Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "docker://6bb0ce4231a7dd4d87b0f533e7b371174956155b100c65b5a0891c21db059395",
                        "image": "nginx:1.14.2",
                        "imageID": "docker-pullable://nginx@sha256:f7988fb6c02e0ce69257d9bd9cf37ae20a60f1df7563c3a2a6abe24160306b8d",
                        "lastState": {},
                        "name": "nginx",
                        "ready": true,
                        "restartCount": 0,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2022-09-15T08:14:49Z"
                            }
                        }
                    }
                ],
                "hostIP": "10.127.254.251",
                "phase": "Running",
                "podIP": "192.168.219.199",
                "podIPs": [
                    {
                        "ip": "192.168.219.199"
                    }
                ],
                "qosClass": "BestEffort",
                "startTime": "2022-09-15T08:14:37Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "c5bd261f5db3fe8ff16ca1748e8d848401cfeb3ed459d3030fb019a8ccc0d3da",
                    "cni.projectcalico.org/podIP": "192.168.219.200/32",
                    "cni.projectcalico.org/podIPs": "192.168.219.200/32"
                },
                "creationTimestamp": "2022-09-15T08:14:37Z",
                "generateName": "nginx-deployment-9456bbbf9-",
                "labels": {
                    "app": "nginx",
                    "pod-template-hash": "9456bbbf9"
                },
                "name": "nginx-deployment-9456bbbf9-ln28l",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "nginx-deployment-9456bbbf9",
                        "uid": "326f8493-fe44-41ca-86f0-bfc740bf2dab"
                    }
                ],
                "resourceVersion": "116497",
                "uid": "dc365fc1-d596-459d-afb5-769f2f7b3f01"
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx:1.14.2",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "nginx",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-8f5pt",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "yaoshicheng-kubernetes-2.cloud.onecloud.io",
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "default",
                "serviceAccountName": "default",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "kube-api-access-8f5pt",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:52Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:52Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "docker://2137c90274ba6b98a775981fc1b18b476bcc316479b624bb32da7c8930318ed4",
                        "image": "nginx:1.14.2",
                        "imageID": "docker-pullable://nginx@sha256:f7988fb6c02e0ce69257d9bd9cf37ae20a60f1df7563c3a2a6abe24160306b8d",
                        "lastState": {},
                        "name": "nginx",
                        "ready": true,
                        "restartCount": 0,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2022-09-15T08:14:51Z"
                            }
                        }
                    }
                ],
                "hostIP": "10.127.254.251",
                "phase": "Running",
                "podIP": "192.168.219.200",
                "podIPs": [
                    {
                        "ip": "192.168.219.200"
                    }
                ],
                "qosClass": "BestEffort",
                "startTime": "2022-09-15T08:14:37Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Pod",
            "metadata": {
                "annotations": {
                    "cni.projectcalico.org/containerID": "5e44ebb4cc4bcedc33bd8636aad19962feb450cebca87fc30834e333c88757c6",
                    "cni.projectcalico.org/podIP": "192.168.219.201/32",
                    "cni.projectcalico.org/podIPs": "192.168.219.201/32"
                },
                "creationTimestamp": "2022-09-15T08:14:37Z",
                "generateName": "nginx-deployment-9456bbbf9-",
                "labels": {
                    "app": "nginx",
                    "pod-template-hash": "9456bbbf9"
                },
                "name": "nginx-deployment-9456bbbf9-tqlmd",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "ReplicaSet",
                        "name": "nginx-deployment-9456bbbf9",
                        "uid": "326f8493-fe44-41ca-86f0-bfc740bf2dab"
                    }
                ],
                "resourceVersion": "116499",
                "uid": "979af1d4-d453-4a82-99b4-52fd102afedf"
            },
            "spec": {
                "containers": [
                    {
                        "image": "nginx:1.14.2",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "nginx",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File",
                        "volumeMounts": [
                            {
                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                                "name": "kube-api-access-624nx",
                                "readOnly": true
                            }
                        ]
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "enableServiceLinks": true,
                "nodeName": "yaoshicheng-kubernetes-2.cloud.onecloud.io",
                "preemptionPolicy": "PreemptLowerPriority",
                "priority": 0,
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "default",
                "serviceAccountName": "default",
                "terminationGracePeriodSeconds": 30,
                "tolerations": [
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/not-ready",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    },
                    {
                        "effect": "NoExecute",
                        "key": "node.kubernetes.io/unreachable",
                        "operator": "Exists",
                        "tolerationSeconds": 300
                    }
                ],
                "volumes": [
                    {
                        "name": "kube-api-access-624nx",
                        "projected": {
                            "defaultMode": 420,
                            "sources": [
                                {
                                    "serviceAccountToken": {
                                        "expirationSeconds": 3607,
                                        "path": "token"
                                    }
                                },
                                {
                                    "configMap": {
                                        "items": [
                                            {
                                                "key": "ca.crt",
                                                "path": "ca.crt"
                                            }
                                        ],
                                        "name": "kube-root-ca.crt"
                                    }
                                },
                                {
                                    "downwardAPI": {
                                        "items": [
                                            {
                                                "fieldRef": {
                                                    "apiVersion": "v1",
                                                    "fieldPath": "metadata.namespace"
                                                },
                                                "path": "namespace"
                                            }
                                        ]
                                    }
                                }
                            ]
                        }
                    }
                ]
            },
            "status": {
                "conditions": [
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "Initialized"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:52Z",
                        "status": "True",
                        "type": "Ready"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:52Z",
                        "status": "True",
                        "type": "ContainersReady"
                    },
                    {
                        "lastProbeTime": null,
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "status": "True",
                        "type": "PodScheduled"
                    }
                ],
                "containerStatuses": [
                    {
                        "containerID": "docker://287e7239f17c9bc4550092e4ace6f152cb58333d81c3d233c64db951b8e82d37",
                        "image": "nginx:1.14.2",
                        "imageID": "docker-pullable://nginx@sha256:f7988fb6c02e0ce69257d9bd9cf37ae20a60f1df7563c3a2a6abe24160306b8d",
                        "lastState": {},
                        "name": "nginx",
                        "ready": true,
                        "restartCount": 0,
                        "started": true,
                        "state": {
                            "running": {
                                "startedAt": "2022-09-15T08:14:51Z"
                            }
                        }
                    }
                ],
                "hostIP": "10.127.254.251",
                "phase": "Running",
                "podIP": "192.168.219.201",
                "podIPs": [
                    {
                        "ip": "192.168.219.201"
                    }
                ],
                "qosClass": "BestEffort",
                "startTime": "2022-09-15T08:14:37Z"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Service",
            "metadata": {
                "creationTimestamp": "2022-09-14T09:35:40Z",
                "labels": {
                    "component": "apiserver",
                    "provider": "kubernetes"
                },
                "name": "kubernetes",
                "namespace": "default",
                "resourceVersion": "209",
                "uid": "95742bdc-22a7-4895-8918-f15f6dc2a749"
            },
            "spec": {
                "clusterIP": "172.16.0.1",
                "clusterIPs": [
                    "172.16.0.1"
                ],
                "internalTrafficPolicy": "Cluster",
                "ipFamilies": [
                    "IPv4"
                ],
                "ipFamilyPolicy": "SingleStack",
                "ports": [
                    {
                        "name": "https",
                        "port": 443,
                        "protocol": "TCP",
                        "targetPort": 6443
                    }
                ],
                "sessionAffinity": "None",
                "type": "ClusterIP"
            },
            "status": {
                "loadBalancer": {}
            }
        },
        {
            "apiVersion": "apps/v1",
            "kind": "Deployment",
            "metadata": {
                "annotations": {
                    "deployment.kubernetes.io/revision": "1",
                    "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx\"},\"name\":\"nginx-deployment\",\"namespace\":\"default\"},\"spec\":{\"replicas\":3,\"selector\":{\"matchLabels\":{\"app\":\"nginx\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"containers\":[{\"image\":\"nginx:1.14.2\",\"name\":\"nginx\",\"ports\":[{\"containerPort\":80}]}]}}}}\n"
                },
                "creationTimestamp": "2022-09-15T08:14:37Z",
                "generation": 1,
                "labels": {
                    "app": "nginx"
                },
                "name": "nginx-deployment",
                "namespace": "default",
                "resourceVersion": "116502",
                "uid": "c87aa538-4f23-4635-8e36-e28b9a97aa64"
            },
            "spec": {
                "progressDeadlineSeconds": 600,
                "replicas": 3,
                "revisionHistoryLimit": 10,
                "selector": {
                    "matchLabels": {
                        "app": "nginx"
                    }
                },
                "strategy": {
                    "rollingUpdate": {
                        "maxSurge": "25%",
                        "maxUnavailable": "25%"
                    },
                    "type": "RollingUpdate"
                },
                "template": {
                    "metadata": {
                        "creationTimestamp": null,
                        "labels": {
                            "app": "nginx"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "image": "nginx:1.14.2",
                                "imagePullPolicy": "IfNotPresent",
                                "name": "nginx",
                                "ports": [
                                    {
                                        "containerPort": 80,
                                        "protocol": "TCP"
                                    }
                                ],
                                "resources": {},
                                "terminationMessagePath": "/dev/termination-log",
                                "terminationMessagePolicy": "File"
                            }
                        ],
                        "dnsPolicy": "ClusterFirst",
                        "restartPolicy": "Always",
                        "schedulerName": "default-scheduler",
                        "securityContext": {},
                        "terminationGracePeriodSeconds": 30
                    }
                }
            },
            "status": {
                "availableReplicas": 3,
                "conditions": [
                    {
                        "lastTransitionTime": "2022-09-15T08:14:52Z",
                        "lastUpdateTime": "2022-09-15T08:14:52Z",
                        "message": "Deployment has minimum availability.",
                        "reason": "MinimumReplicasAvailable",
                        "status": "True",
                        "type": "Available"
                    },
                    {
                        "lastTransitionTime": "2022-09-15T08:14:37Z",
                        "lastUpdateTime": "2022-09-15T08:14:52Z",
                        "message": "ReplicaSet \"nginx-deployment-9456bbbf9\" has successfully progressed.",
                        "reason": "NewReplicaSetAvailable",
                        "status": "True",
                        "type": "Progressing"
                    }
                ],
                "observedGeneration": 1,
                "readyReplicas": 3,
                "replicas": 3,
                "updatedReplicas": 3
            }
        },
        {
            "apiVersion": "apps/v1",
            "kind": "ReplicaSet",
            "metadata": {
                "annotations": {
                    "deployment.kubernetes.io/desired-replicas": "3",
                    "deployment.kubernetes.io/max-replicas": "4",
                    "deployment.kubernetes.io/revision": "1"
                },
                "creationTimestamp": "2022-09-15T08:14:37Z",
                "generation": 1,
                "labels": {
                    "app": "nginx",
                    "pod-template-hash": "9456bbbf9"
                },
                "name": "nginx-deployment-9456bbbf9",
                "namespace": "default",
                "ownerReferences": [
                    {
                        "apiVersion": "apps/v1",
                        "blockOwnerDeletion": true,
                        "controller": true,
                        "kind": "Deployment",
                        "name": "nginx-deployment",
                        "uid": "c87aa538-4f23-4635-8e36-e28b9a97aa64"
                    }
                ],
                "resourceVersion": "116501",
                "uid": "326f8493-fe44-41ca-86f0-bfc740bf2dab"
            },
            "spec": {
                "replicas": 3,
                "selector": {
                    "matchLabels": {
                        "app": "nginx",
                        "pod-template-hash": "9456bbbf9"
                    }
                },
                "template": {
                    "metadata": {
                        "creationTimestamp": null,
                        "labels": {
                            "app": "nginx",
                            "pod-template-hash": "9456bbbf9"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "image": "nginx:1.14.2",
                                "imagePullPolicy": "IfNotPresent",
                                "name": "nginx",
                                "ports": [
                                    {
                                        "containerPort": 80,
                                        "protocol": "TCP"
                                    }
                                ],
                                "resources": {},
                                "terminationMessagePath": "/dev/termination-log",
                                "terminationMessagePolicy": "File"
                            }
                        ],
                        "dnsPolicy": "ClusterFirst",
                        "restartPolicy": "Always",
                        "schedulerName": "default-scheduler",
                        "securityContext": {},
                        "terminationGracePeriodSeconds": 30
                    }
                }
            },
            "status": {
                "availableReplicas": 3,
                "fullyLabeledReplicas": 3,
                "observedGeneration": 1,
                "readyReplicas": 3,
                "replicas": 3
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}`
