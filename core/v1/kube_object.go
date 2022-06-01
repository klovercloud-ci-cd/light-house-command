package v1

import (
	"github.com/klovercloud/lighthouse-command/enums"
)

type KubeObject interface {
	Save(extra map[string]string) error
	Delete(agent string) error
	Update(oldObj interface{},agent string) error
}

func GetObject(object enums.RESOURCE_TYPE) KubeObject {
	if object == enums.CLUSTER_ROLE {
		return &ClusterRole{
			Obj: K8sClusterRole{
				TypeMeta: TypeMeta{
					Kind:       "ClusterRole",
					APIVersion: "rbac.authorization.k8s.io/v1",
				},
			},
		}
	} else if object == enums.CLUSTER_ROLE_BINDGING {
		return &ClusterRoleBinding{
			Obj: k8sClusterRoleBinding{
				TypeMeta: TypeMeta{
					Kind:       "ClusterRoleBinding",
					APIVersion: "rbac.authorization.k8s.io/v1",
				},
			}}
	} else if object == enums.CONFIG_MAP {
		return &ConfigMap{
			Obj: K8sConfigMap{
				TypeMeta: TypeMeta{
					Kind:       "ConfigMap",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.DAEMONSET {
		return &DaemonSet{
			Obj: K8sDaemonSet{
				TypeMeta: TypeMeta{
					Kind:       "DaemonSet",
					APIVersion: "apps/v1",
				},
			},
		}
	} else if object == enums.DEPLOYMENT {
		return &Deployment{
			Obj: K8sDeployment{
				TypeMeta: TypeMeta{
					Kind:       "Deployment",
					APIVersion: "apps/v1",
				},
			},
		}
	} else if object == enums.INGRESS {
		return &Ingress{
			Obj: K8sIngress{
				TypeMeta: TypeMeta{
					Kind:       "Ingress",
					APIVersion: "extensions/v1beta1",
				},
			},
		}
	} else if object == enums.NAMESPACE {
		return &Namespace{
			Obj: K8sNamespace{
				TypeMeta: TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.NETWORK_POLICY {
		return &NetworkPolicy{
			Obj: K8sNetworkPolicy{
				TypeMeta: TypeMeta{
					Kind:       "NetworkPolicy",
					APIVersion: "networking.k8s.io/v1",
				},
			},
		}
	} else if object == enums.POD {
		return &Pod{
			Obj: K8sPod{
				TypeMeta: TypeMeta{
					Kind:       "Pod",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.PERSISTENT_VOLUME {
		return &PersistentVolume{
			Obj: K8sPersistentVolume{
				TypeMeta: TypeMeta{
					Kind:       "PersistentVolume",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.PERSISTENT_VOLUME_CLAIM {
		return &PersistentVolumeClaim{
			Obj: K8sPersistentVolumeClaim{
				TypeMeta: TypeMeta{
					Kind:       "PersistentVolumeClaim",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.REPLICASET {
		return &ReplicaSet{
			Obj: K8sReplicaSet{
				TypeMeta: TypeMeta{
					Kind:       "ReplicaSet",
					APIVersion: "apps/v1",
				},
			},
		}
	} else if object == enums.ROLE {
		return &Role{
			Obj: K8sRole{
				TypeMeta: TypeMeta{
					Kind:       "Role",
					APIVersion: "rbac.authorization.k8s.io/v1",
				},
			},
		}
	} else if object == enums.ROLE_BINDING {
		return &RoleBinding{
			Obj: K8sRoleBinding{
				TypeMeta: TypeMeta{
					Kind:       "RoleBinding",
					APIVersion: "rbac.authorization.k8s.io/v1",
				},
			},
		}
	} else if object == enums.SECRET {
		return &Secret{
			Obj: K8sSecret{
				TypeMeta: TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.SERVICE {
		return &Service{
			Obj: K8sService{
				TypeMeta: TypeMeta{
					Kind:       "Service",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.SERVICE_ACCOUNT {
		return &ServiceAccount{
			Obj: K8sServiceAccount{
				TypeMeta: TypeMeta{
					Kind:       "ServiceAccount",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.STATEFULSET {
		return &StatefulSet{
			Obj: K8sStatefulSet{
				TypeMeta: TypeMeta{
					Kind:       "StatefulSet",
					APIVersion: "apps/v1",
				},
			},
		}
	} else if object == enums.NODE {
		return &Node{
			Obj: K8sNode{
				TypeMeta: TypeMeta{
					Kind:       "Node",
					APIVersion: "v1",
				},
			},
		}
	} else if object == enums.CERTIFICATE {
		return &Certificate{
			Obj: K8sCertificate{
				TypeMeta: TypeMeta{
					Kind:       "Certificate",
					APIVersion: "cert-manager.io/v1",
				},
			},
		}
	}else if object==enums.EVENT{
		return &Event{
			Obj:          K8sEvent{
			},
		}
	}
	return nil

}
