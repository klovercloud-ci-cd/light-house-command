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

type OBJECT string

const (
	CLUSTER_ROLE            = OBJECT("CLUSTERROLE")
	CLUSTER_ROLE_BINDGING   = OBJECT("CLUSTERROLEBINDING")
	CONFIG_MAP              = OBJECT("CONFIGMAP")
	DAEMON_SET              = OBJECT("DAEMONSET")
	DEPLOYMENT              = OBJECT("DEPLOYMENT")
	INGRESS                 = OBJECT("INGRESS")
	NAMESPACE               = OBJECT("NAMESPACE")
	NETWORK_POLICY          = OBJECT("NETWORKPOLICY")
	POD                     = OBJECT("POD")
	PERSISTENT_VOLUME       = OBJECT("PERSISTENTVOLUME")
	PERSISTENT_VOLUME_CLAIM = OBJECT("PERSISTENTVOLUMECLAIM")
	REPLICA_SET             = OBJECT("REPLICASET")
	ROLE                    = OBJECT("ROLE")
	ROLE_BINDING            = OBJECT("ROLEBINDING")
	SECRET                  = OBJECT("SECRET")
	SERVICE                 = OBJECT("SERVICE")
	SERVICE_ACCOUNT         = OBJECT("SERVICEACCOUNT")
	STATEFUL_SET            = OBJECT("STATEFULSET")
	NODE                    = OBJECT("NODE")
	CERTIFICATE             = OBJECT("CERTIFICATE")
)
