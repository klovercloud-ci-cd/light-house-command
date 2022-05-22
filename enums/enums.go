package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

// RESOURCE_TYPE pipeline resource types
type RESOURCE_TYPE string

const ( // CERTIFICATE k8s certificate as resource
	CERTIFICATE = RESOURCE_TYPE("certificate")
	// CLUSTER_ROLE k8s cluster_role as resource
	CLUSTER_ROLE = RESOURCE_TYPE("clusterRole")
	// CLUSTER_ROLE_BINDGING k8s cluster_role_binding as resource
	CLUSTER_ROLE_BINDGING = RESOURCE_TYPE("clusterRoleBinding")
	// CONFIG_MAP k8s config_map as resource
	CONFIG_MAP = RESOURCE_TYPE("configMap")
	// DAEMONSET k8s daemonset as resource
	DAEMONSET = RESOURCE_TYPE("daemonset")
	// DEPLOYMENT k8s deployment as resource
	DEPLOYMENT = RESOURCE_TYPE("deployment")
	// INGRESS k8s ingress as resource
	INGRESS = RESOURCE_TYPE("ingress")
	// NAMESPACE k8s namespace as resource
	NAMESPACE = RESOURCE_TYPE("namespace")
	// NETWORK_POLICY k8s network_policy as resource
	NETWORK_POLICY = RESOURCE_TYPE("networkPolicy")
	// NODE k8s node as resource
	NODE = RESOURCE_TYPE("node")
	// PERSISTENT_VOLUME k8s persistent_volume as resource
	PERSISTENT_VOLUME = RESOURCE_TYPE("persistentVolume")
	// PERSISTENT_VOLUME_CLAIM k8s persistent_volume_claim as resource
	PERSISTENT_VOLUME_CLAIM = RESOURCE_TYPE("persistentVolumeClaim")
	// POD k8s pod as resource
	POD = RESOURCE_TYPE("pod")
	// REPLICASET k8s replicaset as resource
	REPLICASET = RESOURCE_TYPE("replicaset")
	// ROLE k8s role as resource
	ROLE = RESOURCE_TYPE("role")
	// ROLE_BINDING k8s role_binding as resource
	ROLE_BINDING = RESOURCE_TYPE("roleBinding")
	// SECRET k8s secret as resource
	SECRET = RESOURCE_TYPE("secret")
	// SERVICE k8s service as resource
	SERVICE = RESOURCE_TYPE("service")
	// SERVICE_ACCOUNT k8s service_account as resource
	SERVICE_ACCOUNT = RESOURCE_TYPE("serviceAccount")
	// STATEFULSET k8s statefulset as resource
	STATEFULSET = RESOURCE_TYPE("statefulset")
)

// Command kafka command
type Command string

const (
	// Kube object ADD command
	ADD = Command("ADD")
	// Kube object UPDATE command
	UPDATE = Command("UPDATE")
	// Kube object DELETE command
	DELETE = Command("DELETE")
)
