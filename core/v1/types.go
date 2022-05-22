package v1

import (
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"github.com/klovercloud/lighthouse-command/enums"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"time"
)

type UID string
type Time struct {
	time.Time `protobuf:"-"`
}

type KubeEventMessage struct {
	Body   interface{}   `json:"body"`
	Header MessageHeader `json:"header"`
}

type MessageHeader struct {
	Offset  int               `json:"offset"`
	Command enums.Command     `json:"command"`
	Extras  map[string]string `json:"extras"`
}

type OwnerReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion" bson:"apiVersion"`
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind" bson:"kind"`
	// Name of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name" protobuf:"bytes,3,opt,name=name" bson:"name"`
	// UID of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID UID `json:"uid" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID" bson:"uid"`
	// If true, this reference points to the managing controller.
	// +optional
	Controller *bool `json:"controller,omitempty" protobuf:"varint,6,opt,name=controller" bson:"controller"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then
	// the owner cannot be deleted from the key-value store until this
	// reference is removed.
	// Defaults to false.
	// To set this field, a user needs "Delete" permission of the owner,
	// otherwise 422 (Unprocessable Entity) will be returned.
	// +optional
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" protobuf:"varint,7,opt,name=blockOwnerDeletion" bson:"blockOwnerDeletion"`
}
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind" bson:"kind"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	// +optional
	APIVersion string `json:"apiVersion" protobuf:"bytes,2,opt,name=apiVersion" bson:"apiVersion"`
}
type Initializers struct {
	// Pending is a list of initializers that must execute in order before this object is visible.
	// When the last pending initializer is removed, and no failing result is set, the initializers
	// struct will be set to nil and the object is considered as initialized and visible to all
	// clients.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Pending []Initializer `json:"pending" protobuf:"bytes,1,rep,name=pending" patchStrategy:"merge" patchMergeKey:"name" bson:"name"`
	// If result is set with the Failure field, the object will be persisted to storage and then deleted,
	// ensuring that other clients can observe the deletion.
	Result *Status `json:"result,omitempty" protobuf:"bytes,2,opt,name=result" bson:"result"`
}
type ListMeta struct {
	// selfLink is a URL representing this object.
	// Populated by the system.
	// Read-only.
	// +optional
	SelfLink string `json:"selfLink,omitempty" protobuf:"bytes,1,opt,name=selfLink" bson:"selfLink"`

	// String that identifies the server's internal version of this object that
	// can be used by clients to determine when objects have changed.
	// Value must be treated as opaque by clients and passed unmodified back to the server.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,2,opt,name=resourceVersion" bson:"resourceVersion"`

	// continue may be set if the user set a limit on the number of items returned, and indicates that
	// the server has more data available. The value is opaque and may be used to issue another request
	// to the endpoint that served this list to retrieve the next set of available objects. Continuing a
	// consistent list may not be possible if the server configuration has changed or more than a few
	// minutes have passed. The resourceVersion field returned when using this continue value will be
	// identical to the value in the first response, unless you have received this token from an error
	// message.
	Continue string `json:"continue,omitempty" protobuf:"bytes,3,opt,name=continue" bson:"continue"`
}
type StatusReason string

const (
	// StatusReasonUnknown means the server has declined to indicate a specific reason.
	// The details field may contain other information about this error.
	// Status code 500.
	StatusReasonUnknown StatusReason = ""

	// StatusReasonUnauthorized means the server can be reached and understood the request, but requires
	// the user to present appropriate authorization credentials (identified by the WWW-Authenticate header)
	// in order for the action to be completed. If the user has specified credentials on the request, the
	// server considers them insufficient.
	// Status code 401
	StatusReasonUnauthorized StatusReason = "Unauthorized"

	// StatusReasonForbidden means the server can be reached and understood the request, but refuses
	// to take any further action.  It is the result of the server being configured to deny access for some reason
	// to the requested resource by the client.
	// Details (optional):
	//   "kind" string - the kind attribute of the forbidden resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the forbidden resource
	// Status code 403
	StatusReasonForbidden StatusReason = "Forbidden"

	// StatusReasonNotFound means one or more resources required for this operation
	// could not be found.
	// Details (optional):
	//   "kind" string - the kind attribute of the missing resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the missing resource
	// Status code 404
	StatusReasonNotFound StatusReason = "NotFound"

	// StatusReasonAlreadyExists means the resource you are creating already exists.
	// Details (optional):
	//   "kind" string - the kind attribute of the conflicting resource
	//   "id"   string - the identifier of the conflicting resource
	// Status code 409
	StatusReasonAlreadyExists StatusReason = "AlreadyExists"

	// StatusReasonConflict means the requested operation cannot be completed
	// due to a conflict in the operation. The client may need to alter the
	// request. Each resource may define custom details that indicate the
	// nature of the conflict.
	// Status code 409
	StatusReasonConflict StatusReason = "Conflict"

	// StatusReasonGone means the item is no longer available at the server and no
	// forwarding address is known.
	// Status code 410
	StatusReasonGone StatusReason = "Gone"

	// StatusReasonInvalid means the requested create or Update operation cannot be
	// completed due to invalid data provided as part of the request. The client may
	// need to alter the request. When set, the client may use the StatusDetails
	// message field as a summary of the issues encountered.
	// Details (optional):
	//   "kind" string - the kind attribute of the invalid resource
	//   "id"   string - the identifier of the invalid resource
	//   "causes"      - one or more StatusCause entries indicating the data in the
	//                   provided resource that was invalid.  The code, message, and
	//                   field attributes will be set.
	// Status code 422
	StatusReasonInvalid StatusReason = "Invalid"

	// StatusReasonServerTimeout means the server can be reached and understood the request,
	// but cannot complete the action in a reasonable time. The client should retry the request.
	// This is may be due to temporary server load or a transient communication issue with
	// another server. Status code 500 is used because the HTTP spec provides no suitable
	// server-requested client retry and the 5xx class represents actionable errors.
	// Details (optional):
	//   "kind" string - the kind attribute of the resource being acted on.
	//   "id"   string - the operation that is being attempted.
	//   "retryAfterSeconds" int32 - the number of seconds before the operation should be retried
	// Status code 500
	StatusReasonServerTimeout StatusReason = "ServerTimeout"

	// StatusReasonTimeout means that the request could not be completed within the given time.
	// Clients can get this response only when they specified a timeout param in the request,
	// or if the server cannot complete the operation within a reasonable amount of time.
	// The request might succeed with an increased value of timeout param. The client *should*
	// wait at least the number of seconds specified by the retryAfterSeconds field.
	// Details (optional):
	//   "retryAfterSeconds" int32 - the number of seconds before the operation should be retried
	// Status code 504
	StatusReasonTimeout StatusReason = "Timeout"

	// StatusReasonTooManyRequests means the server experienced too many requests within a
	// given window and that the client must wait to perform the action again. A client may
	// always retry the request that led to this error, although the client should wait at least
	// the number of seconds specified by the retryAfterSeconds field.
	// Details (optional):
	//   "retryAfterSeconds" int32 - the number of seconds before the operation should be retried
	// Status code 429
	StatusReasonTooManyRequests StatusReason = "TooManyRequests"

	// StatusReasonBadRequest means that the request itself was invalid, because the request
	// doesn't make any sense, for example deleting a read-only object.  This is different than
	// StatusReasonInvalid above which indicates that the API call could possibly succeed, but the
	// data was invalid.  API calls that return BadRequest can never succeed.
	StatusReasonBadRequest StatusReason = "BadRequest"

	// StatusReasonMethodNotAllowed means that the action the client attempted to perform on the
	// resource was not supported by the code - for instance, attempting to Delete a resource that
	// can only be created. API calls that return MethodNotAllowed can never succeed.
	StatusReasonMethodNotAllowed StatusReason = "MethodNotAllowed"

	// StatusReasonNotAcceptable means that the accept types indicated by the client were not acceptable
	// to the server - for instance, attempting to receive protobuf for a resource that supports only json and yaml.
	// API calls that return NotAcceptable can never succeed.
	// Status code 406
	StatusReasonNotAcceptable StatusReason = "NotAcceptable"

	// StatusReasonRequestEntityTooLarge means that the request entity is too large.
	// Status code 413
	StatusReasonRequestEntityTooLarge StatusReason = "RequestEntityTooLarge"

	// StatusReasonUnsupportedMediaType means that the content type sent by the client is not acceptable
	// to the server - for instance, attempting to send protobuf for a resource that supports only json and yaml.
	// API calls that return UnsupportedMediaType can never succeed.
	// Status code 415
	StatusReasonUnsupportedMediaType StatusReason = "UnsupportedMediaType"

	// StatusReasonInternalError indicates that an internal error occurred, it is unexpected
	// and the outcome of the call is unknown.
	// Details (optional):
	//   "causes" - The original error
	// Status code 500
	StatusReasonInternalError StatusReason = "InternalError"

	// StatusReasonExpired indicates that the request is invalid because the content you are requesting
	// has expired and is no longer available. It is typically associated with watches that can't be
	// serviced.
	// Status code 410 (gone)
	StatusReasonExpired StatusReason = "Expired"

	// StatusReasonServiceUnavailable means that the request itself was valid,
	// but the requested service is unavailable at this time.
	// Retrying the request after some time might succeed.
	// Status code 503
	StatusReasonServiceUnavailable StatusReason = "ServiceUnavailable"
)

type Status struct {
	TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`

	// Status of the operation.
	// One of: "Success" or "Failure".
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status string `json:"status,omitempty" protobuf:"bytes,2,opt,name=status" bson:"status"`
	// A human-readable description of the status of this operation.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message" bson:"message"`
	// A machine-readable description of why this operation is in the
	// "Failure" status. If this value is empty there
	// is no information available. A Reason clarifies an HTTP status
	// code but does not override it.
	// +optional
	Reason StatusReason `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason,casttype=StatusReason" bson:"reason"`
	// Extended data associated with the reason.  Each reason may define its
	// own extended details. This field is optional and the data returned
	// is not guaranteed to conform to any schema except that defined by
	// the reason type.
	// +optional
	Details *StatusDetails `json:"details,omitempty" protobuf:"bytes,5,opt,name=details" bson:"details"`
	// Suggested HTTP return code for this status, 0 if not set.
	// +optional
	Code int32 `json:"code,omitempty" protobuf:"varint,6,opt,name=code" bson:"code"`
}
type CauseType string

const (
	// CauseTypeFieldValueNotFound is used to report failure to find a requested value
	// (e.g. looking up an ID).
	CauseTypeFieldValueNotFound CauseType = "FieldValueNotFound"
	// CauseTypeFieldValueRequired is used to report required values that are not
	// provided (e.g. empty strings, null values, or empty arrays).
	CauseTypeFieldValueRequired CauseType = "FieldValueRequired"
	// CauseTypeFieldValueDuplicate is used to report collisions of values that must be
	// unique (e.g. unique IDs).
	CauseTypeFieldValueDuplicate CauseType = "FieldValueDuplicate"
	// CauseTypeFieldValueInvalid is used to report malformed values (e.g. failed regex
	// match).
	CauseTypeFieldValueInvalid CauseType = "FieldValueInvalid"
	// CauseTypeFieldValueNotSupported is used to report valid (as per formatting rules)
	// values that can not be handled (e.g. an enumerated string).
	CauseTypeFieldValueNotSupported CauseType = "FieldValueNotSupported"
	// CauseTypeUnexpectedServerResponse is used to report when the server responded to the client
	// without the expected return type. The presence of this cause indicates the error may be
	// due to an intervening proxy or the server software malfunctioning.
	CauseTypeUnexpectedServerResponse CauseType = "UnexpectedServerResponse"
)

type StatusCause struct {
	// A machine-readable description of the cause of the error. If this value is
	// empty there is no information available.
	// +optional
	Type CauseType `json:"reason,omitempty" protobuf:"bytes,1,opt,name=reason,casttype=CauseType" bson:"CauseType"`
	// A human-readable description of the cause of the error.  This field may be
	// presented as-is to a reader.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message" bson:"message"`
	// The field of the resource that has caused this error, as named by its JSON
	// serialization. May include dot and postfix notation for nested attributes.
	// Arrays are zero-indexed.  Fields may appear more than once in an array of
	// causes due to fields having multiple errors.
	// Optional.
	//
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	// +optional
	Field string `json:"field,omitempty" protobuf:"bytes,3,opt,name=field" bson:"field"`
}
type StatusDetails struct {
	// The name attribute of the resource associated with the status StatusReason
	// (when there is a single name which can be described).
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// The group attribute of the resource associated with the status StatusReason.
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,2,opt,name=group" bson:"group"`
	// The kind attribute of the resource associated with the status StatusReason.
	// On some operations may differ from the requested resource Kind.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,3,opt,name=kind" bson:"kind"`
	// UID of the resource.
	// (when there is a single resource which can be described).
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	UID UID `json:"uid,omitempty" protobuf:"bytes,6,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID" bson:"uid"`
	// The Causes array includes more details associated with the StatusReason
	// failure. Not all StatusReasons may provide detailed causes.
	// +optional
	Causes []StatusCause `json:"causes,omitempty" protobuf:"bytes,4,rep,name=causes" bson:"causes"`
	// If specified, the time in seconds before the operation should be retried. Some errors may indicate
	// the client must take an alternate action - for those errors this field may indicate how long to wait
	// before taking the alternate action.
	// +optional
	RetryAfterSeconds int32 `json:"retryAfterSeconds,omitempty" protobuf:"varint,5,opt,name=retryAfterSeconds" bson:"retryAfterSeconds"`
}

// Initializer is information about an initializer that has not yet completed.
type Initializer struct {
	// name of the process that is responsible for initializing this object.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
}
type ObjectMeta struct {
	// Name must be unique within a namespace. Is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`

	// GenerateName is an optional prefix, used by the server, to generate a unique
	// name ONLY IF the Name field has not been provided.
	// If this field is used, the name returned to the client will be different
	// than the name passed. This value will also be combined with a unique suffix.
	// The provided value has the same validation rules as the Name field,
	// and may be truncated by the length of the suffix required to make the value
	// unique on the server.
	//
	// If this field is specified and the generated name exists, the server will
	// NOT return a 409 - instead, it will either return 201 Created or 500 with Reason
	// ServerTimeout indicating a unique name could not be found in the time allotted, and the client
	// should retry (optionally after the time indicated in the Retry-After header).
	//
	// Applied only if Name is not specified.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency
	// +optional
	GenerateName string `json:"generateName,omitempty" protobuf:"bytes,2,opt,name=generateName" bson:"generateName"`

	// Namespace defines the space within each name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/namespaces
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace" bson:"namespace"`

	// SelfLink is a URL representing this object.
	// Populated by the system.
	// Read-only.
	// +optional
	SelfLink string `json:"selfLink,omitempty" protobuf:"bytes,4,opt,name=selfLink" bson:"selfLink"`

	// UID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	//
	// Populated by the system.
	// Read-only.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	UID UID `json:"uid,omitempty" protobuf:"bytes,5,opt,name=uid,casttype=k8s.io/kubernetes/pkg/types.UID" bson:"uid"`

	// An opaque value that represents the internal version of this object that can
	// be used by clients to determine when objects have changed. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and passed unmodified back to the server.
	// They may only be valid for a particular resource or set of resources.
	//
	// Populated by the system.
	// Read-only.
	// Value must be treated as opaque by clients and .
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion" bson:"resourceVersion"`

	// A sequence number representing a specific generation of the desired state.
	// Populated by the system. Read-only.
	// +optional
	Generation int64 `json:"generation,omitempty" protobuf:"varint,7,opt,name=generation" bson:"generation" bson:"generation"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system.
	// Read-only.
	// Null for lists.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	CreationTimestamp Time `json:"creationTimestamp,omitempty" protobuf:"bytes,8,opt,name=creationTimestamp" bson:"creationTimestamp"`

	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This
	// field is set by the server when a graceful deletion is requested by the user, and is not
	// directly settable by a client. The resource is expected to be deleted (no longer visible
	// from resource lists, and not reachable by name) after the time in this field, once the
	// finalizers list is empty. As long as the finalizers list contains items, deletion is blocked.
	// Once the deletionTimestamp is set, this value may not be unset or be set further into the
	// future, although it may be shortened or the resource may be deleted prior to this time.
	// For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react
	// by sending a graceful termination signal to the containers in the pod. After that 30 seconds,
	// the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup,
	// remove the pod from the API. In the presence of network partitions, this object may still
	// exist after this timestamp, until an administrator or automated process can determine the
	// resource is fully terminated.
	// If not set, graceful deletion of the object has not been requested.
	//
	// Populated by the system when a graceful deletion is requested.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	DeletionTimestamp *Time `json:"deletionTimestamp,omitempty" protobuf:"bytes,9,opt,name=deletionTimestamp" bson:"deletionTimestamp"`

	// Number of seconds allowed for this object to gracefully terminate before
	// it will be removed from the system. Only set when deletionTimestamp is also set.
	// May only be shortened.
	// Read-only.
	// +optional
	DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" protobuf:"varint,10,opt,name=deletionGracePeriodSeconds" bson:"deletionGracePeriodSeconds"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels" bson:"labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations" bson:"annotations"`

	// List of objects depended by this object. If ALL objects in the list have
	// been deleted, this object will be garbage collected. If this object is managed by a controller,
	// then an entry in this list will point to this controller, with the controller field set to true.
	// There cannot be more than one managing controller.
	// +optional
	// +patchMergeKey=uid
	// +patchStrategy=merge
	OwnerReferences []OwnerReference `json:"ownerReferences,omitempty" patchStrategy:"merge" patchMergeKey:"uid" protobuf:"bytes,13,rep,name=ownerReferences" bson:"ownerReferences"`

	// An initializer is a controller which enforces some system invariant at object creation time.
	// This field is a list of initializers that have not yet acted on this object. If nil or empty,
	// this object has been completely initialized. Otherwise, the object is considered uninitialized
	// and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to
	// observe uninitialized objects.
	//
	// When an object is created, the system will populate this list with the current set of initializers.
	// Only privileged users may set or modify this list. Once it is empty, it may not be modified further
	// by any user.
	Initializers *Initializers `json:"initializers,omitempty" protobuf:"bytes,16,opt,name=initializers" bson:"initializers"`

	// Must be empty before the object is deleted from the registry. Each entry
	// is an identifier for the responsible component that will remove the entry
	// from the list. If the deletionTimestamp of the object is non-nil, entries
	// in this list can only be removed.
	// +optional
	// +patchStrategy=merge
	Finalizers []string `json:"finalizers,omitempty" patchStrategy:"merge" protobuf:"bytes,14,rep,name=finalizers" bson:"finalizers"`

	// The name of the cluster which the object belongs to.
	// This is used to distinguish resources with same name and namespace in different clusters.
	// This field is not set anywhere right now and apiserver is going to ignore it if set in create or Update request.
	// +optional
	ClusterName string `json:"clusterName,omitempty" protobuf:"bytes,15,opt,name=clusterName" bson:"clusterName"`
}

type PolicyRule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
	Verbs []string `json:"verbs" protobuf:"bytes,1,rep,name=verbs" bson:"verbs"`

	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of
	// the enumerated resources in any API group will be allowed.
	// +optional
	APIGroups []string `json:"apiGroups,omitempty" protobuf:"bytes,2,rep,name=apiGroups" bson:"apiGroups"`
	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	// +optional
	Resources []string `json:"resources,omitempty" protobuf:"bytes,3,rep,name=resources" bson:"resources"`
	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
	// +optional
	ResourceNames []string `json:"resourceNames,omitempty" protobuf:"bytes,4,rep,name=resourceNames" bson:"resourceNames"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
	// Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding.
	// Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
	// +optional
	NonResourceURLs []string `json:"nonResourceURLs,omitempty" protobuf:"bytes,5,rep,name=nonResourceURLs" bson:"nonResourceURLs"`
}

type AggregationRule struct {
	// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules.
	// If any of the selectors match, then the ClusterRole's permissions will be added
	// +optional
	ClusterRoleSelectors []LabelSelector `json:"clusterRoleSelectors,omitempty" protobuf:"bytes,1,rep,name=clusterRoleSelectors" bson:"clusterRoleSelectors"`
}

type LabelSelector struct {
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty" protobuf:"bytes,1,rep,name=matchLabels" bson:"matchLabels"`
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	// +optional
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,2,rep,name=matchExpressions" bson:"matchExpressions"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)

type LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key" bson:"key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator LabelSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator" bson:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	// +optional
	Values []string `json:"values,omitempty" protobuf:"bytes,3,rep,name=values" bson:"values"`
}

type Subject struct {
	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount".
	// If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind" bson:"kind"`
	// APIGroup holds the API group of the referenced subject.
	// Defaults to "" for ServiceAccount subjects.
	// Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	// +optional
	APIGroup string `json:"apiGroup,omitempty" protobuf:"bytes,2,opt.name=apiGroup" bson:"apiGroup"`
	// Name of the object being referenced.
	Name string `json:"name" protobuf:"bytes,3,opt,name=name" bson:"name"`
	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty
	// the Authorizer should report an error.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace" bson:"namespace"`
}

type RoleRef struct {
	// APIGroup is the group for the resource being referenced
	APIGroup string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup" bson:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind" bson:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name" protobuf:"bytes,3,opt,name=name" bson:"name"`
}

type DaemonSetUpdateStrategy struct {
	// Type of daemon set Update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	// +optional
	Type DaemonSetUpdateStrategyType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type" bson:"type"`

	// Rolling Update config params. Present only if type = "RollingUpdate".
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be. Same as Deployment `strategy.rollingUpdate`.
	// See https://github.com/kubernetes/kubernetes/issues/35345
	// +optional
	RollingUpdate *RollingUpdateDaemonSet `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate" bson:"rollingUpdate"`
}

type DeploymentStrategy struct {
	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// +optional
	Type DeploymentStrategyType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=DeploymentStrategyType" bson:"type"`

	// Rolling Update config params. Present only if DeploymentStrategyType =
	// RollingUpdate.
	//---
	// TODO: Update this to follow our convention for oneOf, whatever we decide it
	// to be.
	// +optional
	RollingUpdate *RollingUpdateDeployment `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate" bson:"rollingUpdate"`
}

type DaemonSetUpdateStrategyType string

const (
	// Replace the old daemons by new ones using rolling Update i.e replace them on each node one after the other.
	RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"

	// Replace the old daemons only when it's killed
	OnDeleteDaemonSetStrategyType DaemonSetUpdateStrategyType = "OnDelete"
)

// Spec to control the desired behavior of daemon set rolling Update.
type RollingUpdateDaemonSet struct {
	// The maximum number of DaemonSet pods that can be unavailable during the
	// Update. Value can be an absolute number (ex: 5) or a percentage of total
	// number of DaemonSet pods at the start of the Update (ex: 10%). Absolute
	// number is calculated from percentage by rounding up.
	// This cannot be 0.
	// Default value is 1.
	// Example: when this is set to 30%, at most 30% of the total number of nodes
	// that should be running the daemon pod (i.e. status.desiredNumberScheduled)
	// can have their pods stopped for an Update at any given
	// time. The Update starts by stopping at most 30% of those DaemonSet pods
	// and then brings up new DaemonSet pods in their place. Once the new pods
	// are available, it then proceeds onto other DaemonSet pods, thus ensuring
	// that at least 70% of original number of DaemonSet pods are available at
	// all times during the Update.
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty" protobuf:"bytes,1,opt,name=maxUnavailable" bson:"maxUnavailable"`
}

type DaemonSetSpec struct {
	// A label query over pods that are managed by the daemon set.
	// Must match in order to be controlled.
	// It must match the pod template's labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector LabelSelector `json:"selector" protobuf:"bytes,1,opt,name=selector" bson:"selector"`

	// An object that describes the pod that will be created.
	// The DaemonSet will create exactly one copy of this pod on every node
	// that matches the template's node selector (or on every node if no node
	// selector is specified).
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	Template PodTemplateSpec `json:"template" protobuf:"bytes,2,opt,name=template" bson:"template"`

	// An Update strategy to replace existing DaemonSet pods with new pods.
	// +optional
	UpdateStrategy DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,3,opt,name=updateStrategy" bson:"updateStrategy"`

	// The minimum number of seconds for which a newly created DaemonSet pod should
	// be ready without any of its container crashing, for it to be considered
	// available. Defaults to 0 (pod will be considered available as soon as it
	// is ready).
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds" bson:"minReadySeconds"`

	// The number of old history to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit" bson:"revisionHistoryLimit"`
}
type PodTemplateSpec struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`

	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec PodSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
}
type Volume struct {
	// Volume's name.
	// Must be a DNS_LABEL and unique within the pod.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// VolumeSource represents the location and type of the mounted volume.
	// If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	VolumeSource `json:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}
type HostPathType string

const (
	// For backwards compatible, leave it empty if unset
	HostPathUnset HostPathType = ""
	// If nothing exists at the given path, an empty directory will be created there
	// as needed with file mode 0755, having the same group and ownership with Kubelet.
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	// A directory must exist at the given path
	HostPathDirectory HostPathType = "Directory"
	// If nothing exists at the given path, an empty file will be created there
	// as needed with file mode 0644, having the same group and ownership with Kubelet.
	HostPathFileOrCreate HostPathType = "FileOrCreate"
	// A file must exist at the given path
	HostPathFile HostPathType = "File"
	// A UNIX socket must exist at the given path
	HostPathSocket HostPathType = "Socket"
	// A character device must exist at the given path
	HostPathCharDev HostPathType = "CharDevice"
	// A block device must exist at the given path
	HostPathBlockDev HostPathType = "BlockDevice"
)

type HostPathVolumeSource struct {
	// Path of the directory on the host.
	// If the path is a symlink, it will follow the link to the real path.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path string `json:"path" protobuf:"bytes,1,opt,name=path" bson:"path"`
	// Type for HostPath Volume
	// Defaults to ""
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	Type *HostPathType `json:"type,omitempty" protobuf:"bytes,2,opt,name=type" bson:"type"`
}
type StorageMedium string

const (
	StorageMediumDefault   StorageMedium = ""          // use whatever the default is for the node, assume anything we don't explicitly handle is this
	StorageMediumMemory    StorageMedium = "Memory"    // use memory (e.g. tmpfs on linux)
	StorageMediumHugePages StorageMedium = "HugePages" // use hugepages
)

// Protocol defines network protocols supported for things like container ports.
type Protocol string

const (
	// ProtocolTCP is the TCP protocol.
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP is the UDP protocol.
	ProtocolUDP Protocol = "UDP"
	// ProtocolSCTP is the SCTP protocol.
	ProtocolSCTP Protocol = "SCTP"
)

type EmptyDirVolumeSource struct {
	// What type of storage medium should back this directory.
	// The default is "" which means to use the node's default medium.
	// Must be an empty string (default) or Memory.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	Medium StorageMedium `json:"medium,omitempty" protobuf:"bytes,1,opt,name=medium,casttype=StorageMedium" bson:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	// The size limit is also applicable for memory medium.
	// The maximum usage on memory medium EmptyDir would be the minimum value between
	// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
	// The default is nil which means that the limit is undefined.
	// More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
	// +optional
	SizeLimit *string `json:"sizeLimit,omitempty" protobuf:"bytes,2,opt,name=sizeLimit" bson:"sizeLimit"`
}

type GCEPersistentDiskVolumeSource struct {
	// Unique name of the PD resource in GCE. Used to identify the disk in GCE.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	PDName string `json:"pdName" protobuf:"bytes,1,opt,name=pdName" bson:"pdName"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// The partition in the volume that you want to mount.
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	Partition int32 `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition" bson:"partition"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
}

type AWSElasticBlockStoreVolumeSource struct {
	// Unique ID of the persistent disk resource in AWS (Amazon EBS volume).
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID" bson:"volumeID"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// The partition in the volume that you want to mount.
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// +optional
	Partition int32 `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition" bson:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// If omitted, the default is "false".
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
}

type GitRepoVolumeSource struct {
	// Repository URL
	Repository string `json:"repository" protobuf:"bytes,1,opt,name=repository" bson:"repository"`
	// Commit hash for the specified revision.
	// +optional
	Revision string `json:"revision,omitempty" protobuf:"bytes,2,opt,name=revision" bson:"revision"`
	// Target directory name.
	// Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the
	// git repository.  Otherwise, if specified, the volume will contain the git repository in
	// the subdirectory with the given name.
	// +optional
	Directory string `json:"directory,omitempty" protobuf:"bytes,3,opt,name=directory" bson:"directory"`
}

type KeyToPath struct {
	// The key to project.
	Key string `json:"key" protobuf:"bytes,1,opt,name=key" bson:"key"`

	// The relative path of the file to map the key to.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path string `json:"path" protobuf:"bytes,2,opt,name=path" bson:"path"`
	// Optional: mode bits to use on this file, must be a value between 0
	// and 0777. If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,3,opt,name=mode" bson:"mode"`
}

type SecretVolumeSource struct {
	// Name of the secret in the pod's namespace to use.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,1,opt,name=secretName" bson:"secretName"`
	// If unspecified, each key-value pair in the Data field of the referenced
	// Secret will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the Secret,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"  bson:"items"`
	// Optional: mode bits to use on created files by default. Must be a
	// value between 0 and 0777. Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"bytes,3,opt,name=defaultMode" bson:"defaultMode"`
	// Specify whether the Secret or it's keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional" bson:"optional"`
}
type NFSVolumeSource struct {
	// Server is the hostname or IP address of the NFS server.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server string `json:"server" protobuf:"bytes,1,opt,name=server" bson:"server"`

	// Path that is exported by the NFS server.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path string `json:"path" protobuf:"bytes,2,opt,name=path" bson:"path"`

	// ReadOnly here will force
	// the NFS export to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
}

type LocalObjectReference struct {
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// TODO: Add other useful fields. apiVersion, kind, uid?
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`
}

type ISCSIVolumeSource struct {
	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal" bson:"targetPortal"`
	// Target iSCSI Qualified Name.
	IQN string `json:"iqn" protobuf:"bytes,2,opt,name=iqn" bson:"iqn"`
	// iSCSI Target Lun number.
	Lun int32 `json:"lun" protobuf:"varint,3,opt,name=lun" bson:"lun"`
	// iSCSI Interface Name that uses an iSCSI transport.
	// Defaults to 'default' (tcp).
	// +optional
	ISCSIInterface string `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface" bson:"iscsiInterface"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType" bson:"fsType"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly" bson:"readOnly"`
	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	// +optional
	Portals []string `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals" bson:"portals"`
	// whether support iSCSI Discovery CHAP authentication
	// +optional
	DiscoveryCHAPAuth bool `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery" bson:"chapAuthDiscovery"`
	// whether support iSCSI Session CHAP authentication
	// +optional
	SessionCHAPAuth bool `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession" bson:"chapAuthSession"`
	// CHAP Secret for iSCSI target and initiator authentication
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef" bson:"secretRef"`
	// Custom iSCSI Initiator Name.
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	// +optional
	InitiatorName *string `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName" bson:"initiatorName"`
}
type GlusterfsVolumeSource struct {
	// EndpointsName is the endpoint name that details Glusterfs topology.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	EndpointsName string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints" bson:"endpoints"`

	// Path is the Glusterfs volume path.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	Path string `json:"path" protobuf:"bytes,2,opt,name=path" bson:"path"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions.
	// Defaults to false.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
}
type PersistentVolumeClaimVolumeSource struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	ClaimName string `json:"claimName" protobuf:"bytes,1,opt,name=claimName" bson:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	// Default false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly" bson:"readOnly"`
}
type RBDVolumeSource struct {
	// A collection of Ceph monitors.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	CephMonitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors" bson:"monitors"`
	// The rados image name.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	RBDImage string `json:"image" protobuf:"bytes,2,opt,name=image" bson:"image"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType" bson:"fsType"`
	// The rados pool name.
	// Default is rbd.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	RBDPool string `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool" bson:"pool"`
	// The rados user name.
	// Default is admin.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	RadosUser string `json:"user,omitempty" protobuf:"bytes,5,opt,name=user" bson:"user"`
	// Keyring is the path to key ring for RBDUser.
	// Default is /etc/ceph/keyring.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	Keyring string `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring" bson:"keyring"`
	// SecretRef is name of the authentication secret for RBDUser. If provided
	// overrides keyring.
	// Default is nil.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef" bson:"secretRef"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly" bson:"readOnly"`
}

type FlexVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver" bson:"driver"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Optional: SecretRef is reference to the secret object containing
	// sensitive information to pass to the plugin scripts. This may be
	// empty if no secret object is specified. If the secret object
	// contains more than one secret, all secrets are passed to the plugin
	// scripts.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef" bson:"secretRef"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
	// Optional: Extra command options if any.
	// +optional
	Options map[string]string `json:"options,omitempty" protobuf:"bytes,5,rep,name=options" bson:"options"`
}

type CinderVolumeSource struct {
	// volume id used to identify the volume in cinder
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID" bson:"volumeID"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
	// Optional: points to a secret object containing parameters used to connect
	// to OpenStack.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef" bson:"secretRef"`
}

type CephFSVolumeSource struct {
	// Required: Monitors is a collection of Ceph monitors
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors" bson:"monitors"`
	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,2,opt,name=path" bson:"path"`
	// Optional: User is the rados user name, default is admin
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,3,opt,name=user" bson:"user"`
	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretFile string `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile" bson:"secretFile"`
	// Optional: SecretRef is reference to the authentication secret for User, default is empty.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef" bson:"secretRef"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly" bson:"readOnly"`
}

type FlockerVolumeSource struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker
	// should be considered as deprecated
	// +optional
	DatasetName string `json:"datasetName,omitempty" protobuf:"bytes,1,opt,name=datasetName" bson:"datasetName"`
	// UID of the dataset. This is unique identifier of a Flocker dataset
	// +optional
	DatasetUUID string `json:"datasetUUID,omitempty" protobuf:"bytes,2,opt,name=datasetUUID" bson:"datasetUUID"`
}
type ObjectFieldSelector struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion" bson:"apiVersion"`
	// Path of the field to select in the specified API version.
	FieldPath string `json:"fieldPath" protobuf:"bytes,2,opt,name=fieldPath" bson:"fieldPath"`
}

// ResourceFieldSelector represents container resources (cpu, memory) and their output format
type ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars
	// +optional
	ContainerName string `json:"containerName,omitempty" protobuf:"bytes,1,opt,name=containerName" bson:"containerName"`
	// Required: resource to select
	Resource string `json:"resource" protobuf:"bytes,2,opt,name=resource" bson:"resource"`
	// Specifies the output format of the exposed resources, defaults to "1"
	// +optional
	Divisor string `json:"divisor,omitempty" protobuf:"bytes,3,opt,name=divisor" bson:"divisor"`
}
type DownwardAPIVolumeFile struct {
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Path string `json:"path" protobuf:"bytes,1,opt,name=path" bson:"path"`
	// Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,2,opt,name=fieldRef" bson:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
	// +optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,3,opt,name=resourceFieldRef" bson:"resourceFieldRef"`
	// Optional: mode bits to use on this file, must be a value between 0
	// and 0777. If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,4,opt,name=mode" bson:"mode"`
}

type DownwardAPIVolumeSource struct {
	// Items is a list of downward API volume file
	// +optional
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items" bson:"items"`
	// Optional: mode bits to use on created files by default. Must be a
	// value between 0 and 0777. Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode" bson:"defaultMode"`
}

const (
	DownwardAPIVolumeSourceDefaultMode int32 = 0644
)

type FCVolumeSource struct {
	// Optional: FC target worldwide names (WWNs)
	// +optional
	TargetWWNs []string `json:"targetWWNs,omitempty" protobuf:"bytes,1,rep,name=targetWWNs" bson:"targetWWNs"`
	// Optional: FC target lun number
	// +optional
	Lun *int32 `json:"lun,omitempty" protobuf:"varint,2,opt,name=lun" bson:"lun"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType" bson:"fsType"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
	// Optional: FC volume world wide identifiers (wwids)
	// Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	// +optional
	WWIDs []string `json:"wwids,omitempty" protobuf:"bytes,5,rep,name=wwids" bson:"wwids"`
}
type AzureFileVolumeSource struct {
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName" protobuf:"bytes,1,opt,name=secretName" bson:"secretName"`
	// Share Name
	ShareName string `json:"shareName" protobuf:"bytes,2,opt,name=shareName" bson:"shareName"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
}
type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// If unspecified, each key-value pair in the Data field of the referenced
	// ConfigMap will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the ConfigMap,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items" bson:"items"`
	// Optional: mode bits to use on created files by default. Must be a
	// value between 0 and 0777. Defaults to 0644.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,3,opt,name=defaultMode" bson:"defaultMode"`
	// Specify whether the ConfigMap or it's keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional" bson:"optional"`
}

type VsphereVirtualDiskVolumeSource struct {
	// Path that identifies vSphere volume vmdk
	VolumePath string `json:"volumePath" protobuf:"bytes,1,opt,name=volumePath" bson:"volumePath"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Storage Policy Based Management (SPBM) profile name.
	// +optional
	StoragePolicyName string `json:"storagePolicyName,omitempty" protobuf:"bytes,3,opt,name=storagePolicyName" bson:"storagePolicyName"`
	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	// +optional
	StoragePolicyID string `json:"storagePolicyID,omitempty" protobuf:"bytes,4,opt,name=storagePolicyID" bson:"storagePolicyID"`
}
type QuobyteVolumeSource struct {
	// Registry represents a single or multiple Quobyte Registry services
	// specified as a string as host:port pair (multiple entries are separated with commas)
	// which acts as the central registry for volumes
	Registry string `json:"registry" protobuf:"bytes,1,opt,name=registry" bson:"registry"`

	// Volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume" protobuf:"bytes,2,opt,name=volume" bson:"volume"`

	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`

	// User to map volume access to
	// Defaults to serivceaccount user
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,4,opt,name=user" bson:"user"`

	// Group to map volume access to
	// Default is no group
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,5,opt,name=group" bson:"group"`
}

type AzureDataDiskCachingMode string
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "None"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ReadOnly"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "ReadWrite"

	AzureSharedBlobDisk    AzureDataDiskKind = "Shared"
	AzureDedicatedBlobDisk AzureDataDiskKind = "Dedicated"
	AzureManagedDisk       AzureDataDiskKind = "Managed"
)

type AzureDiskVolumeSource struct {
	// The Name of the data disk in the blob storage
	DiskName string `json:"diskName" protobuf:"bytes,1,opt,name=diskName" bson:"diskName"`
	// The URI the data disk in the blob storage
	DataDiskURI string `json:"diskURI" protobuf:"bytes,2,opt,name=diskURI" bson:"diskURI"`
	// Host Caching mode: None, Read Only, Read Write.
	// +optional
	CachingMode *AzureDataDiskCachingMode `json:"cachingMode,omitempty" protobuf:"bytes,3,opt,name=cachingMode,casttype=AzureDataDiskCachingMode" bson:"cachingMode"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly *bool `json:"readOnly,omitempty" protobuf:"varint,5,opt,name=readOnly" bson:"readOnly"`
	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind *AzureDataDiskKind `json:"kind,omitempty" protobuf:"bytes,6,opt,name=kind,casttype=AzureDataDiskKind" bson:"kind"`
}

type PhotonPersistentDiskVolumeSource struct {
	// ID that identifies Photon Controller persistent disk
	PdID string `json:"pdID" protobuf:"bytes,1,opt,name=pdID" bson:"pdID"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
}
type SecretProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// If unspecified, each key-value pair in the Data field of the referenced
	// Secret will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the Secret,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items" bson:"items"`
	// Specify whether the Secret or its key must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional" bson:"optional"`
}

type DownwardAPIProjection struct {
	// Items is a list of DownwardAPIVolume file
	// +optional
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items" bson:"items"`
}
type ConfigMapProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// If unspecified, each key-value pair in the Data field of the referenced
	// ConfigMap will be projected into the volume as a file whose name is the
	// key and content is the value. If specified, the listed keys will be
	// projected into the specified paths, and unlisted keys will not be
	// present. If a key is specified which is not present in the ConfigMap,
	// the volume setup will error unless it is marked optional. Paths must be
	// relative and may not contain the '..' path or start with '..'.
	// +optional
	Items []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items" bson:"items"`
	// Specify whether the ConfigMap or it's keys must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional" bson:"optional"`
}
type ServiceAccountTokenProjection struct {
	// Audience is the intended audience of the token. A recipient of a token
	// must identify itself with an identifier specified in the audience of the
	// token, and otherwise should reject the token. The audience defaults to the
	// identifier of the apiserver.
	//+optional
	Audience string `json:"audience,omitempty" protobuf:"bytes,1,rep,name=audience" bson:"audience"`
	// ExpirationSeconds is the requested duration of validity of the service
	// account token. As the token approaches expiration, the kubelet volume
	// plugin will proactively rotate the service account token. The kubelet will
	// start trying to rotate the token if the token is older than 80 percent of
	// its time to live or if the token is older than 24 hours.Defaults to 1 hour
	// and must be at least 10 minutes.
	//+optional
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty" protobuf:"varint,2,opt,name=expirationSeconds" bson:"expirationSeconds"`
	// Path is the path relative to the mount point of the file to project the
	// token into.
	Path string `json:"path" protobuf:"bytes,3,opt,name=path" bson:"path"`
}

type VolumeProjection struct {
	// all types below are the supported types for projection into the same volume

	// information about the secret data to project
	// +optional
	Secret *SecretProjection `json:"secret,omitempty" protobuf:"bytes,1,opt,name=secret" bson:"secret"`
	// information about the downwardAPI data to project
	// +optional
	DownwardAPI *DownwardAPIProjection `json:"downwardAPI,omitempty" protobuf:"bytes,2,opt,name=downwardAPI" bson:"downwardAPI"`
	// information about the configMap data to project
	// +optional
	ConfigMap *ConfigMapProjection `json:"configMap,omitempty" protobuf:"bytes,3,opt,name=configMap" bson:"configMap"`
	// information about the serviceAccountToken data to project
	// +optional
	ServiceAccountToken *ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty" protobuf:"bytes,4,opt,name=serviceAccountToken" bson:"serviceAccountToken"`
}

type ProjectedVolumeSource struct {
	// list of volume projections
	Sources []VolumeProjection `json:"sources" protobuf:"bytes,1,rep,name=sources" bson:"sources"`
	// Mode bits to use on created files by default. Must be a value between
	// 0 and 0777.
	// Directories within the path are not affected by this setting.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int32 `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode" bson:"defaultMode"`
}

type PortworxVolumeSource struct {
	// VolumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID" bson:"volumeID"`
	// FSType represents the filesystem type to mount
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
}

type ScaleIOVolumeSource struct {
	// The host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway" protobuf:"bytes,1,opt,name=gateway" bson:"gateway"`
	// The name of the storage system as configured in ScaleIO.
	System string `json:"system" protobuf:"bytes,2,opt,name=system" bson:"system"`
	// SecretRef references to the secret for ScaleIO user and other
	// sensitive information. If this is not provided, Login operation will fail.
	SecretRef *LocalObjectReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef" bson:"secretRef"`
	// Flag to enable/disable SSL communication with Gateway, default false
	// +optional
	SSLEnabled bool `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled" bson:"sslEnabled"`
	// The name of the ScaleIO Protection Domain for the configured storage.
	// +optional
	ProtectionDomain string `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain" bson:"protectionDomain"`
	// The ScaleIO Storage Pool associated with the protection domain.
	// +optional
	StoragePool string `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool" bson:"storagePool"`
	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	// Default is ThinProvisioned.
	// +optional
	StorageMode string `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode" bson:"storageMode"`
	// The name of a volume already created in the ScaleIO system
	// that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName" bson:"volumeName"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs".
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly" bson:"readOnly"`
}

type StorageOSVolumeSource struct {
	// VolumeName is the human-readable name of the StorageOS volume.  Volume
	// names are only unique within a namespace.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName" bson:"volumeName"`
	// VolumeNamespace specifies the scope of the volume within StorageOS.  If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	// +optional
	VolumeNamespace string `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace" bson:"volumeNamespace"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
	// SecretRef specifies the secret to use for obtaining the StorageOS API
	// credentials.  If not specified, default values will be attempted.
	// +optional
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef" bson:"secretRef"`
}

type VolumeSource struct {
	// HostPath represents a pre-existing file or directory on the host
	// machine that is directly exposed to the container. This is generally
	// used for system agents or other privileged things that are allowed
	// to see the host machine. Most containers will NOT need this.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// ---
	// TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not
	// mount host directories as read/write.
	// +optional
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" protobuf:"bytes,1,opt,name=hostPath" bson:"hostPath"`
	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty" protobuf:"bytes,2,opt,name=emptyDir"  bson:"emptyDir"`
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GCEPersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" protobuf:"bytes,3,opt,name=gcePersistentDisk" bson:"gcePersistentDisk"`
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,4,opt,name=awsElasticBlockStore" bson:"awsElasticBlockStore"`
	// GitRepo represents a git repository at a particular revision.
	// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an
	// EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir
	// into the Pod's container.
	// +optional
	GitRepo *GitRepoVolumeSource `json:"gitRepo,omitempty" protobuf:"bytes,5,opt,name=gitRepo" bson:"gitRepo"`
	// Secret represents a secret that should populate this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	Secret *SecretVolumeSource `json:"secret,omitempty" protobuf:"bytes,6,opt,name=secret" bson:"secret"`
	// NFS represents an NFS mount on the host that shares a pod's lifetime
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	NFS *NFSVolumeSource `json:"nfs,omitempty" protobuf:"bytes,7,opt,name=nfs" bson:"nfs"`
	// ISCSI represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md
	// +optional
	ISCSI *ISCSIVolumeSource `json:"iscsi,omitempty" protobuf:"bytes,8,opt,name=iscsi" bson:"iscsi"`
	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md
	// +optional
	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" protobuf:"bytes,9,opt,name=glusterfs" bson:"glusterfs"`
	// PersistentVolumeClaimVolumeSource represents a reference to a
	// PersistentVolumeClaim in the same namespace.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" protobuf:"bytes,10,opt,name=persistentVolumeClaim" bson:"persistentVolumeClaim"`
	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md
	// +optional
	RBD *RBDVolumeSource `json:"rbd,omitempty" protobuf:"bytes,11,opt,name=rbd" bson:"rbd"`
	// FlexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// +optional
	FlexVolume *FlexVolumeSource `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume" bson:"flexVolume"`
	// Cinder represents a cinder volume attached and mounted on kubelets host machine
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	Cinder *CinderVolumeSource `json:"cinder,omitempty" protobuf:"bytes,13,opt,name=cinder" bson:"cinder"`
	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// +optional
	CephFS *CephFSVolumeSource `json:"cephfs,omitempty" protobuf:"bytes,14,opt,name=cephfs" bson:"cephfs"`
	// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running
	// +optional
	Flocker *FlockerVolumeSource `json:"flocker,omitempty" protobuf:"bytes,15,opt,name=flocker" bson:"flocker"`
	// DownwardAPI represents downward API about the pod that should populate this volume
	// +optional
	DownwardAPI *DownwardAPIVolumeSource `json:"downwardAPI,omitempty" protobuf:"bytes,16,opt,name=downwardAPI" bson:"downwardAPI"`
	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	FC *FCVolumeSource `json:"fc,omitempty" protobuf:"bytes,17,opt,name=fc" bson:"fc"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// +optional
	AzureFile *AzureFileVolumeSource `json:"azureFile,omitempty" protobuf:"bytes,18,opt,name=azureFile" bson:"azureFile"`
	// ConfigMap represents a configMap that should populate this volume
	// +optional
	ConfigMap *ConfigMapVolumeSource `json:"configMap,omitempty" protobuf:"bytes,19,opt,name=configMap" bson:"configMap"`
	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// +optional
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty" protobuf:"bytes,20,opt,name=vsphereVolume" bson:"vsphereVolume"`
	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// +optional
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty" protobuf:"bytes,21,opt,name=quobyte" bson:"quobyte"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// +optional
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty" protobuf:"bytes,22,opt,name=azureDisk" bson:"azureDisk"`
	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" protobuf:"bytes,23,opt,name=photonPersistentDisk" bson:"photonPersistentDisk"`
	// Items for all in one resources secrets, configmaps, and downward API
	Projected *ProjectedVolumeSource `json:"projected,omitempty" protobuf:"bytes,26,opt,name=projected" bson:"projected"`
	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// +optional
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty" protobuf:"bytes,24,opt,name=portworxVolume" bson:"portworxVolume"`
	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// +optional
	ScaleIO *ScaleIOVolumeSource `json:"scaleIO,omitempty" protobuf:"bytes,25,opt,name=scaleIO" bson:"scaleIO"`
	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	// +optional
	StorageOS *StorageOSVolumeSource `json:"storageos,omitempty" protobuf:"bytes,27,opt,name=storageos" bson:"storageos"`
}

type ContainerPort struct {
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
	// named port in a pod must have a unique name. Name for the port that can be
	// referred to by services.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort" bson:"hostPort"`
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort" protobuf:"varint,3,opt,name=containerPort" bson:"containerPort"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol" bson:"Protocol"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP" bson:"hostIP"`
}
type ConfigMapEnvSource struct {
	// The ConfigMap to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// Specify whether the ConfigMap must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional" bson:"optional"`
}

// SecretEnvSource selects a Secret to populate the environment
// variables with.
//
// The contents of the target Secret's Data field will represent the
// key-value pairs as environment variables.
type SecretEnvSource struct {
	// The Secret to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// Specify whether the Secret must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional" bson:"optional"`
}

type EnvFromSource struct {
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	// +optional
	Prefix string `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix" bson:"prefix"`
	// The ConfigMap to select from
	// +optional
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty" protobuf:"bytes,2,opt,name=configMapRef" bson:"configMapRef"`
	// The Secret to select from
	// +optional
	SecretRef *SecretEnvSource `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef" bson:"secretRef"`
}
type EnvVar struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`

	// Optional: no more than one of the following may be specified.

	// Variable references $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any service environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value" bson:"value"`
	// Source for the environment variable's value. Cannot be used if value is not empty.
	// +optional
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty" protobuf:"bytes,3,opt,name=valueFrom" bson:"valueFrom"`
}

// EnvVarSource represents a source for the value of an EnvVar.
type EnvVarSource struct {
	// Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations,
	// spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.
	// +optional
	FieldRef *ObjectFieldSelector `json:"fieldRef,omitempty" protobuf:"bytes,1,opt,name=fieldRef" bson:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// +optional
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,2,opt,name=resourceFieldRef" bson:"resourceFieldRef"`
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *ConfigMapKeySelector `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef" bson:"configMapKeyRef"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef" bson:"secretKeyRef"`
}

// Selects a key from a ConfigMap.
type ConfigMapKeySelector struct {
	// The ConfigMap to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// The key to select.
	Key string `json:"key" protobuf:"bytes,2,opt,name=key" bson:"key"`
	// Specify whether the ConfigMap or it's key must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional" bson:"optional"`
}

// SecretKeySelector selects a key of a Secret.
type SecretKeySelector struct {
	// The name of the secret in the pod's namespace to select from.
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key" protobuf:"bytes,2,opt,name=key" bson:"key"`
	// Specify whether the Secret or it's key must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional" bson:"optional"`
}
type ResourceList map[ResourceName]string

type ResourceName string

// Resource names must be not more than 63 characters, consisting of upper- or lower-case alphanumeric characters,
// with the -, _, and . characters allowed anywhere, except the first or last character.
// The default convention, matching that for annotations, is to use lower-case names, with dashes, rather than
// camel case, separating compound words.
// Fully-qualified resource typenames are constructed from a DNS-style subdomain, followed by a slash `/` and a name.
const (
	// CPU, in cores. (500m = .5 cores)
	ResourceCPU ResourceName = "cpu"
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	ResourceMemory ResourceName = "memory"
	// Volume size, in bytes (e,g. 5Gi = 5GiB = 5 * 1024 * 1024 * 1024)
	ResourceStorage ResourceName = "storage"
	// Local ephemeral storage, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	// The resource name for ResourceEphemeralStorage is alpha and it can change across releases.
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Limits ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName" bson:"limits"`
	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Requests ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName" bson:"requests"`
}

type MountPropagationMode string

const (
	// MountPropagationNone means that the volume in a container will
	// not receive new mounts from the host or other containers, and filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode corresponds to "private" in Linux terminology.
	MountPropagationNone MountPropagationMode = "None"
	// MountPropagationHostToContainer means that the volume in a container will
	// receive new mounts from the host or other containers, but filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rslave" in Linux terminology).
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	// MountPropagationBidirectional means that the volume in a container will
	// receive new mounts from the host or other containers, and its own mounts
	// will be propagated from the container to the host or other containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rshared" in Linux terminology).
	MountPropagationBidirectional MountPropagationMode = "Bidirectional"
)

type VolumeMount struct {
	// This must match the Name of a Volume.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly" bson:"readOnly"`
	// Path within the container at which the volume should be mounted.  Must
	// not contain ':'.
	MountPath string `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath" bson:"mountPath"`
	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	// +optional
	SubPath string `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath" bson:"subPath"`
	// mountPropagation determines how mounts are propagated from the host
	// to container and the other way around.
	// When not set, MountPropagationNone is used.
	// This field is beta in 1.10.
	// +optional
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty" protobuf:"bytes,5,opt,name=mountPropagation,casttype=MountPropagationMode" bson:"mountPropagation"`
}
type VolumeDevice struct {
	// name must match the name of a persistentVolumeClaim in the pod
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// devicePath is the path inside of the container that the device will be mapped to.
	DevicePath string `json:"devicePath" protobuf:"bytes,2,opt,name=devicePath" bson:"devicePath"`
}
type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
	// a shell, you need to explicitly call out to that shell.
	// Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	// +optional
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command" bson:"command"`
}
type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "HTTPS"
)

type HTTPHeader struct {
	// The header field name
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// The header field value
	Value string `json:"value" protobuf:"bytes,2,opt,name=value" bson:"value"`
}

type HTTPGetAction struct {
	// Path to access on the HTTP server.
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path" bson:"path"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port intstr.IntOrString `json:"port" protobuf:"bytes,2,opt,name=port" bson:"port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,3,opt,name=host" bson:"host"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme URIScheme `json:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme" bson:"scheme"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty" protobuf:"bytes,5,rep,name=httpHeaders" bson:"httpHeaders"`
}
type TCPSocketAction struct {
	// Number or name of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port intstr.IntOrString `json:"port" protobuf:"bytes,1,opt,name=port" bson:"port"`
	// Optional: Host name to connect to, defaults to the pod IP.
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,2,opt,name=host" bson:"host"`
}

type Handler struct {
	// One and only one of the following should be specified.
	// Exec specifies the action to take.
	// +optional
	Exec *ExecAction `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec" bson:"exec"`
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet" bson:"httpGet"`
	// TCPSocket specifies an action involving a TCP port.
	// TCP hooks not yet supported
	// TODO: implement a realistic TCP lifecycle hook
	// +optional
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcpSocket" bson:"tcpSocket"`
}
type Probe struct {
	// The action taken to determine the health of a container
	Handler `json:",inline" protobuf:"bytes,1,opt,name=handler"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds" bson:"initialDelaySeconds"`
	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds" bson:"timeoutSeconds"`
	// How often (in seconds) to perform the probe.
	// Default to 10 seconds. Minimum value is 1.
	// +optional
	PeriodSeconds int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds" bson:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	// Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	// +optional
	SuccessThreshold int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold" bson:"successThreshold"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold" bson:"failureThreshold"`
}
type Lifecycle struct {
	// PostStart is called immediately after a container is created. If the handler fails,
	// the container is terminated and restarted according to its restart policy.
	// Other management of the container blocks until the hook completes.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PostStart *Handler `json:"postStart,omitempty" protobuf:"bytes,1,opt,name=postStart" bson:"postStart"`
	// PreStop is called immediately before a container is terminated.
	// The container is terminated after the handler completes.
	// The reason for termination is passed to the handler.
	// Regardless of the outcome of the handler, the container is eventually terminated.
	// Other management of the container blocks until the hook completes.
	// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
	// +optional
	PreStop *Handler `json:"preStop,omitempty" protobuf:"bytes,2,opt,name=preStop" bson:"preStop"`
}

type TerminationMessagePolicy string

const (
	// TerminationMessageReadFile is the default behavior and will set the container status message to
	// the contents of the container's terminationMessagePath when the container exits.
	TerminationMessageReadFile TerminationMessagePolicy = "File"
	// TerminationMessageFallbackToLogsOnError will read the most recent contents of the container logs
	// for the container status message when the container exits with an error and the
	// terminationMessagePath has no contents.
	TerminationMessageFallbackToLogsOnError TerminationMessagePolicy = "FallbackToLogsOnError"
)

type PullPolicy string

const (
	// PullAlways means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.
	PullAlways PullPolicy = "Always"
	// PullNever means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn't present
	PullNever PullPolicy = "Never"
	// PullIfNotPresent means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

type Capability string

// Adds and removes POSIX capabilities from running containers.
type Capabilities struct {
	// Added capabilities
	// +optional
	Add []Capability `json:"add,omitempty" protobuf:"bytes,1,rep,name=add,casttype=Capability" bson:"add"`
	// Removed capabilities
	// +optional
	Drop []Capability `json:"drop,omitempty" protobuf:"bytes,2,rep,name=drop,casttype=Capability" bson:"drop"`
}

type SELinuxOptions struct {
	// User is a SELinux user label that applies to the container.
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,1,opt,name=user" bson:"user"`
	// Role is a SELinux role label that applies to the container.
	// +optional
	Role string `json:"role,omitempty" protobuf:"bytes,2,opt,name=role" bson:"role"`
	// Type is a SELinux type label that applies to the container.
	// +optional
	Type string `json:"type,omitempty" protobuf:"bytes,3,opt,name=type" bson:"type"`
	// Level is SELinux level label that applies to the container.
	// +optional
	Level string `json:"level,omitempty" protobuf:"bytes,4,opt,name=level" bson:"level"`
}

type ProcMountType string

const (
	// DefaultProcMount uses the container runtime defaults for readonly and masked
	// paths for /proc.  Most container runtimes mask certain paths in /proc to avoid
	// accidental security exposure of special devices or information.
	DefaultProcMount ProcMountType = "Default"

	// UnmaskedProcMount bypasses the default masking behavior of the container
	// runtime and ensures the newly created /proc the container stays in tact with
	// no modifications.
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type SecurityContext struct {
	// The capabilities to add/drop when running containers.
	// Defaults to the default set of capabilities granted by the container runtime.
	// +optional
	Capabilities *Capabilities `json:"capabilities,omitempty" protobuf:"bytes,1,opt,name=capabilities" bson:"capabilities"`
	// Run container in privileged mode.
	// Processes in privileged containers are essentially equivalent to root on the host.
	// Defaults to false.
	// +optional
	Privileged *bool `json:"privileged,omitempty" protobuf:"varint,2,opt,name=privileged" bson:"privileged"`
	// The SELinux context to be applied to the container.
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	SELinuxOptions *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,3,opt,name=seLinuxOptions" bson:"seLinuxOptions"`
	// The UID to run the entrypoint of the container process.
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty" protobuf:"varint,4,opt,name=runAsUser" bson:"runAsUser"`
	// The GID to run the entrypoint of the container process.
	// Uses runtime default if unset.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsGroup *int64 `json:"runAsGroup,omitempty" protobuf:"varint,8,opt,name=runAsGroup" bson:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" protobuf:"varint,5,opt,name=runAsNonRoot" bson:"runAsNonRoot"`
	// Whether this container has a read-only root filesystem.
	// Default is false.
	// +optional
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty" protobuf:"varint,6,opt,name=readOnlyRootFilesystem" bson:"readOnlyRootFilesystem"`
	// AllowPrivilegeEscalation controls whether a process can gain more
	// privileges than its parent process. This bool directly controls if
	// the no_new_privs flag will be set on the container process.
	// AllowPrivilegeEscalation is true always when the container is:
	// 1) run as Privileged
	// 2) has CAP_SYS_ADMIN
	// +optional
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty" protobuf:"varint,7,opt,name=allowPrivilegeEscalation" bson:"allowPrivilegeEscalation"`
	// procMount denotes the type of proc mount to use for the containers.
	// The default is DefaultProcMount which uses the container runtime defaults for
	// readonly paths and masked paths.
	// This requires the ProcMountType feature flag to be enabled.
	// +optional
	ProcMount *ProcMountType `json:"procMount,omitempty" protobuf:"bytes,9,opt,name=procMount" bson:"procMount"`
}
type Container struct {
	// Name of the container specified as a DNS_LABEL.
	// Each container in a pod must have a unique name (DNS_LABEL).
	// Cannot be updated.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Docker image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image" bson:"image"`
	// Entrypoint array. Not executed within a shell.
	// The docker image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax
	// can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Command []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command" bson:"command"`
	// Arguments to the entrypoint.
	// The docker image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax
	// can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Args []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args" bson:"args"`
	// Container's working directory.
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// Cannot be updated.
	// +optional
	WorkingDir string `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir" bson:"workingDir"`
	// List of ports to expose from the container. Exposing a port here gives
	// the system additional information about the network connections a
	// container uses, but is primarily informational. Not specifying a port here
	// DOES NOT prevent that port from being exposed. Any port which is
	// listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=containerPort
	// +patchStrategy=merge
	Ports []ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports" bson:"ports"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	EnvFrom []EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom" bson:"envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Env []EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env" bson:"env"`
	// Compute Resources required by this container.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources" bson:"resources"`
	// Pod volumes to mount into the container's filesystem.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts" bson:"volumeMounts"`
	// volumeDevices is the list of block devices to be used by the container.
	// This is an alpha feature and may change in the future.
	// +patchMergeKey=devicePath
	// +patchStrategy=merge
	// +optional
	VolumeDevices []VolumeDevice `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices" bson:"volumeDevices"`
	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe" bson:"livenessProbe"`
	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe" bson:"readinessProbe"`
	// Actions that the management system should take in response to container lifecycle events.
	// Cannot be updated.
	// +optional
	Lifecycle *Lifecycle `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle" bson:"lifecycle"`
	// Optional: Path at which the file to which the container's termination message
	// will be written is mounted into the container's filesystem.
	// Message written is intended to be brief final status, such as an assertion failure message.
	// Will be truncated by the node if greater than 4096 bytes. The total message length across
	// all containers will be limited to 12kb.
	// Defaults to /dev/termination-log.
	// Cannot be updated.
	// +optional
	TerminationMessagePath string `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath" bson:"terminationMessagePath"`
	// Indicate how the termination message should be populated. File will use the contents of
	// terminationMessagePath to populate the container status message on both success and failure.
	// FallbackToLogsOnError will use the last chunk of container log output if the termination
	// message file is empty and the container exited with an error.
	// The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
	// Defaults to File.
	// Cannot be updated.
	// +optional
	TerminationMessagePolicy TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy" bson:"terminationMessagePolicy"`
	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy" bson:"imagePullPolicy"`
	// Security options the pod should run with.
	// More info: https://kubernetes.io/docs/concepts/policy/security-context/
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *SecurityContext `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext" bson:"securityContext"`

	// Variables for interactive containers, these have very specialized use-cases (e.g. debugging)
	// and shouldn't be used for general purpose containers.

	// Whether this container should allocate a buffer for stdin in the container runtime. If this
	// is not set, reads from stdin in the container will always result in EOF.
	// Default is false.
	// +optional
	Stdin bool `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin" bson:"stdin"`
	// Whether the container runtime should close the stdin channel after it has been opened by
	// a single attach. When stdin is true the stdin stream will remain open across multiple attach
	// sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
	// first client attaches to stdin, and then remains open and accepts data until the client disconnects,
	// at which time stdin is closed and remains closed until the container is restarted. If this
	// flag is false, a container processes that reads from stdin will never receive an EOF.
	// Default is false
	// +optional
	StdinOnce bool `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce" bson:"stdinOnce"`
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.
	// Default is false.
	// +optional
	TTY bool `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty" bson:"tty"`
}
type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

type DNSPolicy string

const (
	// DNSClusterFirstWithHostNet indicates that the pod should use cluster DNS
	// first, if it is available, then fall back on the default
	// (as determined by kubelet) DNS settings.
	DNSClusterFirstWithHostNet DNSPolicy = "ClusterFirstWithHostNet"

	// DNSClusterFirst indicates that the pod should use cluster DNS
	// first unless hostNetwork is true, if it is available, then
	// fall back on the default (as determined by kubelet) DNS settings.
	DNSClusterFirst DNSPolicy = "ClusterFirst"

	// DNSDefault indicates that the pod should use the default (as
	// determined by kubelet) DNS settings.
	DNSDefault DNSPolicy = "Default"

	// DNSNone indicates that the pod should use empty DNS settings. DNS
	// parameters such as nameservers and search paths should be defined via
	// DNSConfig.
	DNSNone DNSPolicy = "None"
)

type Sysctl struct {
	// Name of a property to set
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Value of a property to set
	Value string `json:"value" protobuf:"bytes,2,opt,name=value" bson:"value"`
}

type PodSecurityContext struct {
	// The SELinux context to be applied to all containers.
	// If unspecified, the container runtime will allocate a random SELinux context for each
	// container.  May also be set in SecurityContext.  If set in
	// both SecurityContext and PodSecurityContext, the value specified in SecurityContext
	// takes precedence for that container.
	// +optional
	SELinuxOptions *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,1,opt,name=seLinuxOptions" bson:"seLinuxOptions"`
	// The UID to run the entrypoint of the container process.
	// Defaults to user specified in image metadata if unspecified.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence
	// for that container.
	// +optional
	RunAsUser *int64 `json:"runAsUser,omitempty" protobuf:"varint,2,opt,name=runAsUser" bson:"runAsUser"`
	// The GID to run the entrypoint of the container process.
	// Uses runtime default if unset.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence
	// for that container.
	// +optional
	RunAsGroup *int64 `json:"runAsGroup,omitempty" protobuf:"varint,6,opt,name=runAsGroup" bson:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// May also be set in SecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" protobuf:"varint,3,opt,name=runAsNonRoot" bson:"runAsNonRoot"`
	// A list of groups applied to the first process run in each container, in addition
	// to the container's primary GID.  If unspecified, no groups will be added to
	// any container.
	// +optional
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty" protobuf:"varint,4,rep,name=supplementalGroups" bson:"supplementalGroups"`
	// A special supplemental group that applies to all containers in a pod.
	// Some volume types allow the Kubelet to change the ownership of that volume
	// to be owned by the pod:
	//
	// 1. The owning GID will be the FSGroup
	// 2. The setgid bit is set (new files created in the volume will be owned by FSGroup)
	// 3. The permission bits are OR'd with rw-rw----
	//
	// If unset, the Kubelet will not modify the ownership and permissions of any volume.
	// +optional
	FSGroup *int64 `json:"fsGroup,omitempty" protobuf:"varint,5,opt,name=fsGroup" bson:"fsGroup"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported
	// sysctls (by the container runtime) might fail to launch.
	// +optional
	Sysctls []Sysctl `json:"sysctls,omitempty" protobuf:"bytes,7,rep,name=sysctls" bson:"sysctls"`
}
type NodeSelectorOperator string

const (
	NodeSelectorOpIn           NodeSelectorOperator = "In"
	NodeSelectorOpNotIn        NodeSelectorOperator = "NotIn"
	NodeSelectorOpExists       NodeSelectorOperator = "Exists"
	NodeSelectorOpDoesNotExist NodeSelectorOperator = "DoesNotExist"
	NodeSelectorOpGt           NodeSelectorOperator = "Gt"
	NodeSelectorOpLt           NodeSelectorOperator = "Lt"
)

type NodeSelectorRequirement struct {
	// The label key that the selector applies to.
	Key string `json:"key" protobuf:"bytes,1,opt,name=key" bson:"key"`
	// Represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator NodeSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=NodeSelectorOperator" bson:"operator"`
	// An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. If the operator is Gt or Lt, the values
	// array must have a single element, which will be interpreted as an integer.
	// This array is replaced during a strategic merge patch.
	// +optional
	Values []string `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// +optional
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,1,rep,name=matchExpressions"  bson:"matchExpressions"`
	// A list of node selector requirements by node's fields.
	// +optional
	MatchFields []NodeSelectorRequirement `json:"matchFields,omitempty" protobuf:"bytes,2,rep,name=matchFields" bson:"matchFields"`
}

type NodeSelector struct {
	//Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms" protobuf:"bytes,1,rep,name=nodeSelectorTerms" bson:"nodeSelectorTerms"`
}

type PreferredSchedulingTerm struct {
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight" protobuf:"varint,1,opt,name=weight" bson:"weight"`
	// A node selector term, associated with the corresponding weight.
	Preference NodeSelectorTerm `json:"preference" protobuf:"bytes,2,opt,name=preference" bson:"preference"`
}

type NodeAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to an Update), the system
	// will try to eventually evict the pod from its node.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution *NodeSelector `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to an Update), the system
	// may or may not try to eventually evict the pod from its node.
	// +optional
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelector `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,opt,name=requiredDuringSchedulingIgnoredDuringExecution" bson:"requiredDuringSchedulingIgnoredDuringExecution"`
	// The scheduler will prefer to schedule pods to nodes that satisfy
	// the affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node matches the corresponding matchExpressions; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution" bson:"preferredDuringSchedulingIgnoredDuringExecution"`
}
type PodAffinityTerm struct {
	// A label query over a set of resources, in this case pods.
	// +optional
	LabelSelector LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,1,opt,name=labelSelector" bson:"labelSelector"`
	// namespaces specifies which namespaces the labelSelector applies to (matches against);
	// null or empty list means "this pod's namespace"
	// +optional
	Namespaces []string `json:"namespaces,omitempty" protobuf:"bytes,2,rep,name=namespaces" bson:"namespaces"`
	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
	// the labelSelector in the specified namespaces, where co-located is defined as running on a node
	// whose value of the label with key topologyKey matches that of any node on which any of the
	// selected pods is running.
	// Empty topologyKey is not allowed.
	TopologyKey string `json:"topologyKey" protobuf:"bytes,3,opt,name=topologyKey" bson:"topologyKey"`
}
type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `json:"weight" protobuf:"varint,1,opt,name=weight" bson:"weight"`
	// Required. A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm" protobuf:"bytes,2,opt,name=podAffinityTerm" bson:"podAffinityTerm"`
}
type PodAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label Update), the
	// system will try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution []PodAffinityTerm  `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label Update), the
	// system may or may not try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	RequiredDuringSchedulingIgnoredDuringExecution []PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution" bson:"requiredDuringSchedulingIgnoredDuringExecution"`
	// The scheduler will prefer to schedule pods to nodes that satisfy
	// the affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution" bson:"preferredDuringSchedulingIgnoredDuringExecution"`
}
type PodAntiAffinity struct {
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// If the anti-affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the anti-affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label Update), the
	// system will try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	// RequiredDuringSchedulingRequiredDuringExecution []PodAffinityTerm  `json:"requiredDuringSchedulingRequiredDuringExecution,omitempty"`

	// If the anti-affinity requirements specified by this field are not met at
	// scheduling time, the pod will not be scheduled onto the node.
	// If the anti-affinity requirements specified by this field cease to be met
	// at some point during pod execution (e.g. due to a pod label Update), the
	// system may or may not try to eventually evict the pod from its node.
	// When there are multiple elements, the lists of nodes corresponding to each
	// podAffinityTerm are intersected, i.e. all terms must be satisfied.
	// +optional
	RequiredDuringSchedulingIgnoredDuringExecution []PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution" bson:"requiredDuringSchedulingIgnoredDuringExecution"`
	// The scheduler will prefer to schedule pods to nodes that satisfy
	// the anti-affinity expressions specified by this field, but it may choose
	// a node that violates one or more of the expressions. The node that is
	// most preferred is the one with the greatest sum of weights, i.e.
	// for each node that meets all of the scheduling requirements (resource
	// request, requiredDuringScheduling anti-affinity expressions, etc.),
	// compute a sum by iterating through the elements of this field and adding
	// "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
	// node(s) with the highest sum are the most preferred.
	// +optional
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution" bson:"preferredDuringSchedulingIgnoredDuringExecution"`
}

type Affinity struct {
	// Describes node affinity scheduling rules for the pod.
	// +optional
	NodeAffinity *NodeAffinity `json:"nodeAffinity,omitempty" protobuf:"bytes,1,opt,name=nodeAffinity" bson:"nodeAffinity"`
	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
	// +optional
	PodAffinity *PodAffinity `json:"podAffinity,omitempty" protobuf:"bytes,2,opt,name=podAffinity" bson:"podAffinity"`
	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
	// +optional
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty" protobuf:"bytes,3,opt,name=podAntiAffinity" bson:"podAntiAffinity"`
}

type TaintEffect string

const (
	// Do not allow new pods to schedule onto the node unless they tolerate the taint,
	// but allow all pods submitted to Kubelet without going through the scheduler
	// to start, and allow all already-running pods to continue running.
	// Enforced by the scheduler.
	TaintEffectNoSchedule TaintEffect = "NoSchedule"
	// Like TaintEffectNoSchedule, but the scheduler tries not to schedule
	// new pods onto the node, rather than prohibiting new pods from scheduling
	// onto the node entirely. Enforced by the scheduler.
	TaintEffectPreferNoSchedule TaintEffect = "PreferNoSchedule"
	// NOT YET IMPLEMENTED. TODO: Uncomment field once it is implemented.
	// Like TaintEffectNoSchedule, but additionally do not allow pods submitted to
	// Kubelet without going through the scheduler to start.
	// Enforced by Kubelet and the scheduler.
	// TaintEffectNoScheduleNoAdmit TaintEffect = "NoScheduleNoAdmit"

	// Evict any already-running pods that do not tolerate the taint.
	// Currently enforced by NodeController.
	TaintEffectNoExecute TaintEffect = "NoExecute"
)

type Toleration struct {
	// Key is the taint key that the toleration applies to. Empty means match all taint keys.
	// If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	// +optional
	Key string `json:"key,omitempty" protobuf:"bytes,1,opt,name=key" bson:"key"`
	// Operator represents a key's relationship to the value.
	// Valid operators are Exists and Equal. Defaults to Equal.
	// Exists is equivalent to wildcard for value, so that a pod can
	// tolerate all taints of a particular category.
	// +optional
	Operator TolerationOperator `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator" bson:"operator"`
	// Value is the taint value the toleration matches to.
	// If the operator is Exists, the value should be empty, otherwise just a regular string.
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value" bson:"value"`
	// Effect indicates the taint effect to match. Empty means match all taint effects.
	// When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	// +optional
	Effect TaintEffect `json:"effect,omitempty" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect" bson:"effect"`
	// TolerationSeconds represents the period of time the toleration (which must be
	// of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,
	// it is not set, which means tolerate the taint forever (do not evict). Zero and
	// negative values will be treated as 0 (evict immediately) by the system.
	// +optional
	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty" protobuf:"varint,5,opt,name=tolerationSeconds" bson:"tolerationSeconds"`
}

// A toleration operator is the set of operators that can be used in a toleration.
type TolerationOperator string

const (
	TolerationOpExists TolerationOperator = "Exists"
	TolerationOpEqual  TolerationOperator = "Equal"
)

type PodDNSConfig struct {
	// A list of DNS name server IP addresses.
	// This will be appended to the base nameservers generated from DNSPolicy.
	// Duplicated nameservers will be removed.
	// +optional
	Nameservers []string `json:"nameservers,omitempty" protobuf:"bytes,1,rep,name=nameservers" bson:"nameservers"`
	// A list of DNS search domains for host-name lookup.
	// This will be appended to the base search paths generated from DNSPolicy.
	// Duplicated search paths will be removed.
	// +optional
	Searches []string `json:"searches,omitempty" protobuf:"bytes,2,rep,name=searches" bson:"searches"`
	// A list of DNS resolver options.
	// This will be merged with the base options generated from DNSPolicy.
	// Duplicated entries will be removed. Resolution options given in Options
	// will override those that appear in the base DNSPolicy.
	// +optional
	Options []PodDNSConfigOption `json:"options,omitempty" protobuf:"bytes,3,rep,name=options" bson:"options"`
}

// PodDNSConfigOption defines DNS resolver options of a pod.
type PodDNSConfigOption struct {
	// Required.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// +optional
	Value *string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value" bson:"value"`
}
type HostAlias struct {
	// IP address of the host file entry.
	IP string `json:"ip,omitempty" protobuf:"bytes,1,opt,name=ip" bson:"ip"`
	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames,omitempty" protobuf:"bytes,2,rep,name=hostnames" bson:"hostnames"`
}

type PodSpec struct {
	// List of volumes that can be mounted by containers belonging to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	Volumes []Volume `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes" bson:"volumes"`
	// List of initialization containers belonging to the pod.
	// Init containers are executed in order prior to containers being started. If any
	// init container fails, the pod is considered to have failed and is handled according
	// to its restartPolicy. The name for an init container or normal container must be
	// unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes.
	// The resourceRequirements of an init container are taken into account during scheduling
	// by finding the highest request/limit for each resource type, and then using the max of
	// of that value or the sum of the normal containers. Limits are applied to init containers
	// in a similar fashion.
	// Init containers cannot currently be added or removed.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	// +patchMergeKey=name
	// +patchStrategy=merge
	InitContainers []Container `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers" bson:"initContainers"`
	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Cannot be updated.
	// +patchMergeKey=name
	// +patchStrategy=merge
	Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers" bson:"containers"`
	// Restart policy for all containers within the pod.
	// One of Always, OnFailure, Never.
	// Default to Always.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	// +optional
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy" bson:"RestartPolicy"`
	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in Delete request.
	// Value must be non-negative integer. The value zero indicates Delete immediately.
	// If this value is nil, the default grace period will be used instead.
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds" bson:"terminationGracePeriodSeconds"`
	// Optional duration in seconds the pod may be active on the node relative to
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer.
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds" bson:"activeDeadlineSeconds"`
	// Set DNS policy for the pod.
	// Defaults to "ClusterFirst".
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
	// DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
	// To have DNS options set along with hostNetwork, you have to specify DNS policy
	// explicitly to 'ClusterFirstWithHostNet'.
	// +optional
	DNSPolicy DNSPolicy `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy" bson:"DNSPolicy"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector" bson:"nodeSelector"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName" bson:"serviceAccountName"`
	// DeprecatedServiceAccount is a depreciated alias for ServiceAccountName.
	// Deprecated: Use serviceAccountName instead.
	// +k8s:conversion-gen=false
	// +optional
	DeprecatedServiceAccount string `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount" bson:"serviceAccount"`
	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	// +optional
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken" bson:"automountServiceAccountToken"`

	// NodeName is a request to schedule this pod onto a specific node. If it is non-empty,
	// the scheduler simply schedules this pod onto that node, assuming that it fits resource
	// requirements.
	// +optional
	NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName" bson:"nodeName"`
	// Host networking requested for this pod. Use the host's network namespace.
	// If this option is set, the ports that will be used must be specified.
	// Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork" bson:"hostNetwork"`
	// Use the host's pid namespace.
	// Optional: Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID" bson:"hostPID"`
	// Use the host's ipc namespace.
	// Optional: Default to false.
	// +k8s:conversion-gen=false
	// +optional
	HostIPC bool `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC" bson:"hostIPC"`
	// Share a single process namespace between all of the containers in a pod.
	// When this is set containers will be able to view and signal processes from other containers
	// in the same pod, and the first process in each container will not be assigned PID 1.
	// HostPID and ShareProcessNamespace cannot both be set.
	// Optional: Default to false.
	// This field is beta-level and may be disabled with the PodShareProcessNamespace feature.
	// +k8s:conversion-gen=false
	// +optional
	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace" bson:"shareProcessNamespace"`
	// SecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext" bson:"securityContext"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets" bson:"imagePullSecrets"`
	// Specifies the hostname of the Pod
	// If not specified, the pod's hostname will be set to a system-defined value.
	// +optional
	Hostname string `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname" bson:"hostname"`
	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
	// If not specified, the pod will not have a domainname at all.
	// +optional
	Subdomain string `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain" bson:"subdomain"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity" bson:"affinity"`
	// If specified, the pod will be dispatched by specified scheduler.
	// If not specified, the pod will be dispatched by default scheduler.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName" bson:"schedulerName"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations" bson:"tolerations"`
	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts
	// file if specified. This is only valid for non-hostNetwork pods.
	// +optional
	// +patchMergeKey=ip
	// +patchStrategy=merge
	HostAliases []HostAlias `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases" bson:"hostAliases"`
	// If specified, indicates the pod's priority. "system-node-critical" and
	// "system-cluster-critical" are two special keywords which indicate the
	// highest priorities with the former being the highest priority. Any other
	// name must be defined by creating a PriorityClass object with that name.
	// If not specified, the pod priority will be default or zero if there is no
	// default.
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName" bson:"priorityClassName"`
	// The priority value. Various system components use this field to find the
	// priority of the pod. When Priority Admission Controller is enabled, it
	// prevents users from setting this field. The admission controller populates
	// this field from PriorityClassName.
	// The higher the value, the higher the priority.
	// +optional
	Priority *int32 `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority" bson:"priority"`
	// Specifies the DNS parameters of a pod.
	// Parameters specified here will be merged to the generated DNS
	// configuration based on DNSPolicy.
	// +optional
	DNSConfig *PodDNSConfig `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig" bson:"dnsConfig"`

	// If specified, all readiness gates will be evaluated for pod readiness.
	// A pod is ready when all its containers are ready AND
	// all conditions specified in the readiness gates have status equal to "True"
	// More info: https://github.com/kubernetes/community/blob/master/keps/sig-network/0007-pod-ready%2B%2B.md
	// +optional
	ReadinessGates []PodReadinessGate `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates" bson:"readinessGates"`
	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
	// to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
	// If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an
	// empty definition that uses the default runtime handler.
	// More info: https://github.com/kubernetes/community/blob/master/keps/sig-node/0014-runtime-class.md
	// This is an alpha feature and may change in the future.
	// +optional
	RuntimeClassName *string `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"  bson:"runtimeClassName"`
}
type PodConditionType string

// These are valid conditions of pod.
const (
	// PodScheduled represents status of the scheduling process for this pod.
	PodScheduled PodConditionType = "PodScheduled"
	// PodReady means the pod is able to service requests and should be added to the
	// load balancing pools of all matching services.
	PodReady PodConditionType = "Ready"
	// PodInitialized means that all init containers in the pod have started successfully.
	PodInitialized PodConditionType = "Initialized"
	// PodReasonUnschedulable reason in PodScheduled PodCondition means that the scheduler
	// can't schedule the pod right now, for example due to insufficient resources in the cluster.
	PodReasonUnschedulable = "Unschedulable"
	// ContainersReady indicates whether all containers in the pod are ready.
	ContainersReady PodConditionType = "ContainersReady"
)

type PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type.
	ConditionType PodConditionType `json:"conditionType" protobuf:"bytes,1,opt,name=conditionType,casttype=PodConditionType" bson:"conditionType"`
}

type DaemonSetStatus struct {
	// The number of nodes that are running at least 1
	// daemon pod and are supposed to run the daemon pod.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	CurrentNumberScheduled int32 `json:"currentNumberScheduled" protobuf:"varint,1,opt,name=currentNumberScheduled" bson:"currentNumberScheduled"`

	// The number of nodes that are running the daemon pod, but are
	// not supposed to run the daemon pod.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	NumberMisscheduled int32 `json:"numberMisscheduled" protobuf:"varint,2,opt,name=numberMisscheduled" bson:"numberMisscheduled"`

	// The total number of nodes that should be running the daemon
	// pod (including nodes correctly running the daemon pod).
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled" protobuf:"varint,3,opt,name=desiredNumberScheduled" bson:"desiredNumberScheduled"`

	// The number of nodes that should be running the daemon pod and have one
	// or more of the daemon pod running and ready.
	NumberReady int32 `json:"numberReady" protobuf:"varint,4,opt,name=numberReady" bson:"numberReady"`

	// The most recent generation observed by the daemon set controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,5,opt,name=observedGeneration" bson:"observedGeneration"`

	// The total number of nodes that are running updated daemon pod
	// +optional
	UpdatedNumberScheduled int32 `json:"updatedNumberScheduled,omitempty" protobuf:"varint,6,opt,name=updatedNumberScheduled" bson:"updatedNumberScheduled"`

	// The number of nodes that should be running the
	// daemon pod and have one or more of the daemon pod running and
	// available (ready for at least spec.minReadySeconds)
	// +optional
	NumberAvailable int32 `json:"numberAvailable,omitempty" protobuf:"varint,7,opt,name=numberAvailable" bson:"numberAvailable"`

	// The number of nodes that should be running the
	// daemon pod and have none of the daemon pod running and available
	// (ready for at least spec.minReadySeconds)
	// +optional
	NumberUnavailable int32 `json:"numberUnavailable,omitempty" protobuf:"varint,8,opt,name=numberUnavailable" bson:"numberUnavailable"`

	// CountByClusterId of hash collisions for the DaemonSet. The DaemonSet controller
	// uses this field as a collision avoidance mechanism when it needs to
	// create the name for the newest ControllerRevision.
	// +optional
	CollisionCount *int32 `json:"collisionCount,omitempty" protobuf:"varint,9,opt,name=collisionCount" bson:"collisionCount"`

	// Represents the latest available observations of a DaemonSet's current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []DaemonSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,10,rep,name=conditions" bson:"conditions"`
}

type DaemonSetConditionType string

// TODO: Add valid condition types of a DaemonSet.
type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// DaemonSetCondition describes the state of a DaemonSet at a certain point.
type DaemonSetCondition struct {
	// Type of DaemonSet condition.
	Type DaemonSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DaemonSetConditionType" bson:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus" bson:"status"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason" bson:"reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message" bson:"message"`
}
type DeploymentStrategyType string

const (
	// Kill all existing pods before creating new ones.
	RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

	// Replace the old ReplicaSets by new one using rolling Update i.e gradually scale down the old ReplicaSets and scale up the new one.
	RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
	// The maximum number of pods that can be unavailable during the Update.
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding down.
	// This can not be 0 if MaxSurge is 0.
	// Defaults to 25%.
	// Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods
	// immediately when the rolling Update starts. Once new pods are ready, old ReplicaSet
	// can be scaled down further, followed by scaling up the new ReplicaSet, ensuring
	// that the total number of pods available at all times during the Update is at
	// least 70% of desired pods.
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty" protobuf:"bytes,1,opt,name=maxUnavailable"`

	// The maximum number of pods that can be scheduled above the desired number of
	// pods.
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// This can not be 0 if MaxUnavailable is 0.
	// Absolute number is calculated from percentage by rounding up.
	// Defaults to 25%.
	// Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when
	// the rolling Update starts, such that the total number of old and new pods do not exceed
	// 130% of desired pods. Once old pods have been killed,
	// new ReplicaSet can be scaled up further, ensuring that total number of pods running
	// at any time during the Update is at most 130% of desired pods.
	// +optional
	MaxSurge *intstr.IntOrString `json:"maxSurge,omitempty" protobuf:"bytes,2,opt,name=maxSurge"`
}

type DeploymentSpec struct {
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas" bson:"replicas"`

	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this deployment.
	// It must match the pod template's labels.
	Selector LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector" bson:"selector"`

	// Template describes the pods that will be created.
	Template PodTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template" bson:"template"`

	// The deployment strategy to use to replace existing pods with new ones.
	// +optional
	Strategy DeploymentStrategy `json:"strategy,omitempty" protobuf:"bytes,4,opt,name=strategy" bson:"strategy"`

	// Minimum number of seconds for which a newly created pod should be ready
	// without any of its container crashing, for it to be considered available.
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,5,opt,name=minReadySeconds" bson:"minReadySeconds"`

	// The number of old ReplicaSets to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit" bson:"revisionHistoryLimit"`

	// Indicates that the deployment is paused.
	// +optional
	Paused bool `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused" bson:"paused"`

	// The maximum time in seconds for a deployment to make progress before it
	// is considered to be failed. The deployment controller will continue to
	// process failed deployments and a condition with a ProgressDeadlineExceeded
	// reason will be surfaced in the deployment status. Note that progress will
	// not be estimated during the time a deployment is paused. Defaults to 600s.
	ProgressDeadlineSeconds *int32 `json:"progressDeadlineSeconds,omitempty" protobuf:"varint,9,opt,name=progressDeadlineSeconds" bson:"progressDeadlineSeconds"`
}

type DeploymentConditionType string

// These are valid conditions of a deployment.
const (
	// Available means the deployment is available, ie. at least the minimum available
	// replicas required are up and running for at least minReadySeconds.
	DeploymentAvailable DeploymentConditionType = "Available"
	// Progressing means the deployment is progressing. Progress for a deployment is
	// considered when a new replica set is created or adopted, and when new pods scale
	// up or old pods scale down. Progress is not estimated for paused deployments or
	// when progressDeadlineSeconds is not specified.
	DeploymentProgressing DeploymentConditionType = "Progressing"
	// ReplicaFailure is added in a deployment when one of its pods fails to be created
	// or deleted.
	DeploymentReplicaFailure DeploymentConditionType = "ReplicaFailure"
)

type DeploymentCondition struct {
	// Type of deployment condition.
	Type DeploymentConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DeploymentConditionType" bson:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus" bson:"status"`
	// The last time this condition was updated.
	LastUpdateTime Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime" bson:"lastUpdateTime"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason" bson:"reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message" bson:"message"`
}

type DeploymentStatus struct {
	// The generation observed by the deployment controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration" bson:"observedGeneration"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas" bson:"replicas"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" protobuf:"varint,3,opt,name=updatedReplicas" bson:"updatedReplicas"`

	// Total number of ready pods targeted by this deployment.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,7,opt,name=readyReplicas" bson:"readyReplicas"`

	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty" protobuf:"varint,4,opt,name=availableReplicas" bson:"availableReplicas"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of
	// pods that are still required for the deployment to have 100% available capacity. They may
	// either be pods that are running but not yet available or pods that still have not been created.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty" protobuf:"varint,5,opt,name=unavailableReplicas" bson:"unavailableReplicas"`

	// Represents the latest available observations of a deployment's current state.
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions" bson:"conditions"`

	// CountByClusterId of hash collisions for the Deployment. The Deployment controller uses this
	// field as a collision avoidance mechanism when it needs to create the name for the
	// newest ReplicaSet.
	// +optional
	CollisionCount *int32 `json:"collisionCount,omitempty" protobuf:"varint,8,opt,name=collisionCount" bson:"collisionCount"`
}

type IngressBackend struct {
	// Specifies the name of the referenced service.
	ServiceName string `json:"serviceName" protobuf:"bytes,1,opt,name=serviceName" bson:"serviceName"`

	// Specifies the port of the referenced service.
	ServicePort intstr.IntOrString `json:"servicePort" protobuf:"bytes,2,opt,name=servicePort" bson:"servicePort"`
}
type IngressTLS struct {
	// Hosts are a list of hosts included in the TLS certificate. The values in
	// this list must match the name/s used in the tlsSecret. Defaults to the
	// wildcard host setting for the loadbalancer controller fulfilling this
	// Ingress, if left unspecified.
	// +optional
	Hosts []string `json:"hosts,omitempty" protobuf:"bytes,1,rep,name=hosts" bson:"hosts"`
	// SecretName is the name of the secret used to terminate SSL traffic on 443.
	// Field is left optional to allow SSL routing based on SNI hostname alone.
	// If the SNI host in a listener conflicts with the "Host" header field used
	// by an IngressRule, the SNI host is used for termination and value of the
	// Host header is used for routing.
	// +optional
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,2,opt,name=secretName" bson:"secretName"`
	// TODO: Consider specifying different modes of termination, protocols etc.
}
type HTTPIngressRuleValue struct {
	// A collection of paths that map requests to backends.
	Paths []HTTPIngressPath `json:"paths" protobuf:"bytes,1,rep,name=paths" bson:"paths"`
	// TODO: Consider adding fields for ingress-type specific global
	// options usable by a loadbalancer, like http keep-alive.
}

// HTTPIngressPath associates a path regex with a backend. Incoming urls matching
// the path are forwarded to the backend.
type HTTPIngressPath struct {
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1,
	// (i.e this follows the egrep/unix syntax, not the perl syntax)
	// matched against the path of an incoming request. Currently it can
	// contain characters disallowed from the conventional "path"
	// part of a URL as defined by RFC 3986. Paths must begin with
	// a '/'. If unspecified, the path defaults to a catch all sending
	// traffic to the backend.
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path" bson:"path"`

	// Backend defines the referenced service endpoint to which the traffic
	// will be forwarded to.
	Backend IngressBackend `json:"backend" protobuf:"bytes,2,opt,name=backend" bson:"backend"`
}

type IngressRuleValue struct {
	//TODO:
	// 1. Consider renaming this resource and the associated rules so they
	// aren't tied to Ingress. They can be used to route intra-cluster traffic.
	// 2. Consider adding fields for ingress-type specific global options
	// usable by a loadbalancer, like http keep-alive.

	// +optional
	HTTP *HTTPIngressRuleValue `json:"http,omitempty" protobuf:"bytes,1,opt,name=http" bson:"http"`
}

type IngressRule struct {
	// Host is the fully qualified domain name of a network host, as defined
	// by RFC 3986. Note the following deviations from the "host" part of the
	// URI as defined in the RFC:
	// 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the
	//	  IP in the Spec of the parent Ingress.
	// 2. The `:` delimiter is not respected because ports are not allowed.
	//	  Currently the port of an Ingress is implicitly :80 for http and
	//	  :443 for https.
	// Both these may change in the future.
	// Incoming requests are matched against the host before the IngressRuleValue.
	// If the host is unspecified, the Ingress routes all traffic based on the
	// specified IngressRuleValue.
	// +optional
	Host string `json:"host,omitempty" protobuf:"bytes,1,opt,name=host" bson:"host"`
	// IngressRuleValue represents a rule to route requests for this IngressRule.
	// If unspecified, the rule defaults to a http catch-all. Whether that sends
	// just traffic matching the host to the default backend or all traffic to the
	// default backend, is left to the controller fulfilling the Ingress. Http is
	// currently the only supported IngressRuleValue.
	// +optional
	IngressRuleValue `json:",inline,omitempty" protobuf:"bytes,2,opt,name=ingressRuleValue"`
}

type IngressSpec struct {
	// A default backend capable of servicing requests that don't match any
	// rule. At least one of 'backend' or 'rules' must be specified. This field
	// is optional to allow the loadbalancer controller or defaulting logic to
	// specify a global default.
	// +optional
	Backend *IngressBackend `json:"backend,omitempty" protobuf:"bytes,1,opt,name=backend" bson:"backend"`

	// TLS configuration. Currently the Ingress only supports a single TLS
	// port, 443. If multiple members of this list specify different hosts, they
	// will be multiplexed on the same port according to the hostname specified
	// through the SNI TLS extension, if the ingress controller fulfilling the
	// ingress supports SNI.
	// +optional
	TLS []IngressTLS `json:"tls,omitempty" protobuf:"bytes,2,rep,name=tls" bson:"tls"`

	// A list of host rules used to configure the Ingress. If unspecified, or
	// no rule matches, all traffic is sent to the default backend.
	// +optional
	Rules []IngressRule `json:"rules,omitempty" protobuf:"bytes,3,rep,name=rules" bson:"rules"`
	// TODO: Add the ability to specify load-balancer IP through claims
}

// LoadBalancerStatus represents the status of a load-balancer.
type LoadBalancerStatus struct {
	// Ingress is a list containing ingress points for the load-balancer.
	// Traffic intended for the service should be sent to these ingress points.
	// +optional
	Ingress []LoadBalancerIngress `json:"ingress,omitempty" protobuf:"bytes,1,rep,name=ingress" bson:"ingress"`
}

// LoadBalancerIngress represents the status of a load-balancer ingress point:
// traffic intended for the service should be sent to an ingress point.
type LoadBalancerIngress struct {
	// IP is set for load-balancer ingress points that are IP based
	// (typically GCE or OpenStack load-balancers)
	// +optional
	IP string `json:"ip,omitempty" protobuf:"bytes,1,opt,name=ip" bson:"ip"`

	// Hostname is set for load-balancer ingress points that are DNS based
	// (typically AWS load-balancers)
	// +optional
	Hostname string `json:"hostname,omitempty" protobuf:"bytes,2,opt,name=hostname" bson:"hostname"`
}

type IngressStatus struct {
	// LoadBalancer contains the current status of the load-balancer.
	// +optional
	LoadBalancer LoadBalancerStatus `json:"loadBalancer,omitempty" protobuf:"bytes,1,opt,name=loadBalancer" bson:"loadBalancer"`
}
type NetworkPolicyPort struct {
	// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this
	// field defaults to TCP.
	// +optional
	Protocol *Protocol `json:"protocol,omitempty" protobuf:"bytes,1,opt,name=protocol,casttype=k8s.io/api/core/v1.Protocol" bson:"protocol"`

	// The port on the given protocol. This can either be a numerical or named port on
	// a pod. If this field is not provided, this matches all port names and numbers.
	// +optional
	Port *intstr.IntOrString `json:"port,omitempty" protobuf:"bytes,2,opt,name=port" bson:"port"`
}
type IPBlock struct {
	// CIDR is a string representing the IP Block
	// Valid examples are "192.168.1.1/24"
	CIDR string `json:"cidr" protobuf:"bytes,1,name=cidr" bson:"cidr"`
	// Except is a slice of CIDRs that should not be included within an IP Block
	// Valid examples are "192.168.1.1/24"
	// Except values will be rejected if they are outside the CIDR range
	// +optional
	Except []string `json:"except,omitempty" protobuf:"bytes,2,rep,name=except" bson:"except"`
}

type NetworkPolicyPeer struct {
	// This is a label selector which selects Pods. This field follows standard label
	// selector semantics; if present but empty, it selects all pods.
	//
	// If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects
	// the Pods matching PodSelector in the Namespaces selected by NamespaceSelector.
	// Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.
	// +optional
	PodSelector *LabelSelector `json:"podSelector,omitempty" protobuf:"bytes,1,opt,name=podSelector" bson:"podSelector"`

	// Selects Namespaces using cluster-scoped labels. This field follows standard label
	// selector semantics; if present but empty, it selects all namespaces.
	//
	// If PodSelector is also set, then the NetworkPolicyPeer as a whole selects
	// the Pods matching PodSelector in the Namespaces selected by NamespaceSelector.
	// Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.
	// +optional
	NamespaceSelector *LabelSelector `json:"namespaceSelector,omitempty" protobuf:"bytes,2,opt,name=namespaceSelector" bson:"namespaceSelector"`

	// IPBlock defines policy on a particular IPBlock. If this field is set then
	// neither of the other fields can be.
	// +optional
	IPBlock *IPBlock `json:"ipBlock,omitempty" protobuf:"bytes,3,rep,name=ipBlock" bson:"ipBlock"`
}

type NetworkPolicyIngressRule struct {
	// List of ports which should be made accessible on the pods selected for this
	// rule. Each item in this list is combined using a logical OR. If this field is
	// empty or missing, this rule matches all ports (traffic not restricted by port).
	// If this field is present and contains at least one item, then this rule allows
	// traffic only if the traffic matches at least one port in the list.
	// +optional
	Ports []NetworkPolicyPort `json:"ports,omitempty" protobuf:"bytes,1,rep,name=ports" bson:"ports"`

	// List of sources which should be able to access the pods selected for this rule.
	// Items in this list are combined using a logical OR operation. If this field is
	// empty or missing, this rule matches all sources (traffic not restricted by
	// source). If this field is present and contains at least on item, this rule
	// allows traffic only if the traffic matches at least one item in the from list.
	// +optional
	From []NetworkPolicyPeer `json:"from,omitempty" protobuf:"bytes,2,rep,name=from" bson:"from"`
}
type PolicyType string

const (
	// PolicyTypeIngress is a NetworkPolicy that affects ingress traffic on selected pods
	PolicyTypeIngress PolicyType = "Ingress"
	// PolicyTypeEgress is a NetworkPolicy that affects egress traffic on selected pods
	PolicyTypeEgress PolicyType = "Egress"
)

type NetworkPolicyEgressRule struct {
	// List of destination ports for outgoing traffic.
	// Each item in this list is combined using a logical OR. If this field is
	// empty or missing, this rule matches all ports (traffic not restricted by port).
	// If this field is present and contains at least one item, then this rule allows
	// traffic only if the traffic matches at least one port in the list.
	// +optional
	Ports []NetworkPolicyPort `json:"ports,omitempty" protobuf:"bytes,1,rep,name=ports" bson:"ports"`

	// List of destinations for outgoing traffic of pods selected for this rule.
	// Items in this list are combined using a logical OR operation. If this field is
	// empty or missing, this rule matches all destinations (traffic not restricted by
	// destination). If this field is present and contains at least one item, this rule
	// allows traffic only if the traffic matches at least one item in the to list.
	// +optional
	To []NetworkPolicyPeer `json:"to,omitempty" protobuf:"bytes,2,rep,name=to" bson:"to"`
}
type NetworkPolicySpec struct {
	// Selects the pods to which this NetworkPolicy object applies. The array of
	// ingress rules is applied to any pods selected by this field. Multiple network
	// policies can select the same set of pods. In this case, the ingress rules for
	// each are combined additively. This field is NOT optional and follows standard
	// label selector semantics. An empty podSelector matches all pods in this
	// namespace.
	PodSelector LabelSelector `json:"podSelector" protobuf:"bytes,1,opt,name=podSelector" bson:"podSelector"`

	// List of ingress rules to be applied to the selected pods. Traffic is allowed to
	// a pod if there are no NetworkPolicies selecting the pod
	// (and cluster policy otherwise allows the traffic), OR if the traffic source is
	// the pod's local node, OR if the traffic matches at least one ingress rule
	// across all of the NetworkPolicy objects whose podSelector matches the pod. If
	// this field is empty then this NetworkPolicy does not allow any traffic (and serves
	// solely to ensure that the pods it selects are isolated by default)
	// +optional
	Ingress []NetworkPolicyIngressRule `json:"ingress,omitempty" protobuf:"bytes,2,rep,name=ingress" bson:"ingress"`

	// List of egress rules to be applied to the selected pods. Outgoing traffic is
	// allowed if there are no NetworkPolicies selecting the pod (and cluster policy
	// otherwise allows the traffic), OR if the traffic matches at least one egress rule
	// across all of the NetworkPolicy objects whose podSelector matches the pod. If
	// this field is empty then this NetworkPolicy limits all outgoing traffic (and serves
	// solely to ensure that the pods it selects are isolated by default).
	// This field is beta-level in 1.8
	// +optional
	Egress []NetworkPolicyEgressRule `json:"egress,omitempty" protobuf:"bytes,3,rep,name=egress" bson:"egress"`

	// List of rule types that the NetworkPolicy relates to.
	// Valid options are Ingress, Egress, or Ingress,Egress.
	// If this field is not specified, it will default based on the existence of Ingress or Egress rules;
	// policies that contain an Egress section are assumed to affect Egress, and all policies
	// (whether or not they contain an Ingress section) are assumed to affect Ingress.
	// If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ].
	// Likewise, if you want to write a policy that specifies that no egress is allowed,
	// you must specify a policyTypes value that include "Egress" (since such a policy would not include
	// an Egress section and would otherwise default to just [ "Ingress" ]).
	// This field is beta-level in 1.8
	// +optional
	PolicyTypes []PolicyType `json:"policyTypes,omitempty" protobuf:"bytes,4,rep,name=policyTypes,casttype=PolicyType" bson:"policyTypes"`
}

type PodPhase string

// These are the valid statuses of pods.
const (
	// PodPending means the pod has been accepted by the system, but one or more of the containers
	// has not been started. This includes time before being bound to a node, as well as time spent
	// pulling images onto the host.
	PodPending PodPhase = "Pending"
	// PodRunning means the pod has been bound to a node and all of the containers have been started.
	// At least one container is still running or is in the process of being restarted.
	PodRunning PodPhase = "Running"
	// PodSucceeded means that all containers in the pod have voluntarily terminated
	// with a container exit code of 0, and the system is not going to restart any of these containers.
	PodSucceeded PodPhase = "Succeeded"
	// PodFailed means that all containers in the pod have terminated, and at least one container has
	// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
	PodFailed PodPhase = "Failed"
	// PodUnknown means that for some reason the state of the pod could not be obtained, typically due
	// to an error in communicating with the host of the pod.
	PodUnknown PodPhase = "Unknown"
)

// PodCondition contains details for the current condition of this pod.
type PodCondition struct {
	// Type is the type of the condition.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type PodConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PodConditionType" bson:"type"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus" bson:"status"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime" bson:"lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason" bson:"reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message" bson:"message"`
}

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.

// ContainerStateWaiting is a waiting state of a container.
type ContainerStateWaiting struct {
	// (brief) reason the container is not yet running.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,1,opt,name=reason" bson:"reason"`
	// Message regarding why the container is not yet running.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message" bson:"message"`
}

// ContainerStateRunning is a running state of a container.
type ContainerStateRunning struct {
	// Time at which the container was last (re-)started
	// +optional
	StartedAt Time `json:"startedAt,omitempty" protobuf:"bytes,1,opt,name=startedAt" bson:"startedAt"`
}
type ContainerStateTerminated struct {
	// Exit status from the last termination of the container
	ExitCode int32 `json:"exitCode" protobuf:"varint,1,opt,name=exitCode" bson:"exitCode"`
	// Signal from the last termination of the container
	// +optional
	Signal int32 `json:"signal,omitempty" protobuf:"varint,2,opt,name=signal" bson:"signal"`
	// (brief) reason from the last termination of the container
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason" bson:"reason"`
	// Message regarding the last termination of the container
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message" bson:"message"`
	// Time at which previous execution of the container started
	// +optional
	StartedAt Time `json:"startedAt,omitempty" protobuf:"bytes,5,opt,name=startedAt" bson:"startedAt"`
	// Time at which the container last terminated
	// +optional
	FinishedAt Time `json:"finishedAt,omitempty" protobuf:"bytes,6,opt,name=finishedAt" bson:"finishedAt"`
	// Container's ID in the format 'docker://<container_id>'
	// +optional
	ContainerID string `json:"containerID,omitempty" protobuf:"bytes,7,opt,name=containerID" bson:"containerID"`
}
type ContainerState struct {
	// Details about a waiting container
	// +optional
	Waiting *ContainerStateWaiting `json:"waiting,omitempty" protobuf:"bytes,1,opt,name=waiting" bson:"waiting"`
	// Details about a running container
	// +optional
	Running *ContainerStateRunning `json:"running,omitempty" protobuf:"bytes,2,opt,name=running" bson:"running"`
	// Details about a terminated container
	// +optional
	Terminated *ContainerStateTerminated `json:"terminated,omitempty" protobuf:"bytes,3,opt,name=terminated" bson:"terminated"`
}

type ContainerStatus struct {
	// This must be a DNS_LABEL. Each container in a pod must have a unique name.
	// Cannot be updated.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Details about the container's current condition.
	// +optional
	State ContainerState `json:"state,omitempty" protobuf:"bytes,2,opt,name=state" bson:"state"`
	// Details about the container's last termination condition.
	// +optional
	LastTerminationState ContainerState `json:"lastState,omitempty" protobuf:"bytes,3,opt,name=lastState" bson:"lastState"`
	// Specifies whether the container has passed its readiness probe.
	Ready bool `json:"ready" protobuf:"varint,4,opt,name=ready" bson:"ready"`
	// The number of times the container has been restarted, currently based on
	// the number of dead containers that have not yet been removed.
	// Note that this is calculated from dead containers. But those containers are subject to
	// garbage collection. This value will get capped at 5 by GC.
	RestartCount int32 `json:"restartCount" protobuf:"varint,5,opt,name=restartCount" bson:"restartCount"`
	// The image the container is running.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// TODO(dchen1107): Which image the container is running with?
	Image string `json:"image" protobuf:"bytes,6,opt,name=image" bson:"image"`
	// ImageID of the container's image.
	ImageID string `json:"imageID" protobuf:"bytes,7,opt,name=imageID" bson:"imageID"`
	// Container's ID in the format 'docker://<container_id>'.
	// +optional
	ContainerID string `json:"containerID,omitempty" protobuf:"bytes,8,opt,name=containerID" bson:"containerID"`
}

type PodQOSClass string

const (
	// PodQOSGuaranteed is the Guaranteed qos class.
	PodQOSGuaranteed PodQOSClass = "Guaranteed"
	// PodQOSBurstable is the Burstable qos class.
	PodQOSBurstable PodQOSClass = "Burstable"
	// PodQOSBestEffort is the BestEffort qos class.
	PodQOSBestEffort PodQOSClass = "BestEffort"
)

type PodStatus struct {
	// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle.
	// The conditions array, the reason and message fields, and the individual container status
	// arrays contain more detail about the pod's status.
	// There are five possible phase values:
	//
	// Pending: The pod has been accepted by the Kubernetes system, but one or more of the
	// container images has not been created. This includes time before being scheduled as
	// well as time spent downloading images over the network, which could take a while.
	// Running: The pod has been bound to a node, and all of the containers have been created.
	// At least one container is still running, or is in the process of starting or restarting.
	// Succeeded: All containers in the pod have terminated in success, and will not be restarted.
	// Failed: All containers in the pod have terminated, and at least one container has
	// terminated in failure. The container either exited with non-zero status or was terminated
	// by the system.
	// Unknown: For some reason the state of the pod could not be obtained, typically due to an
	// error in communicating with the host of the pod.
	//
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase
	// +optional
	Phase PodPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PodPhase" bson:"phase"`
	// Current service state of pod.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []PodCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions" bson:"conditions"`
	// A human readable message indicating details about why the pod is in this condition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message" bson:"message"`
	// A brief CamelCase message indicating details about why the pod is in this state.
	// e.g. 'Evicted'
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason" bson:"reason"`
	// nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be
	// scheduled right away as preemption victims receive their graceful termination periods.
	// This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide
	// to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to
	// give the resources on this node to a higher priority pod that is created after preemption.
	// As a result, this field may be different than PodSpec.nodeName when the pod is
	// scheduled.
	// +optional
	NominatedNodeName string `json:"nominatedNodeName,omitempty" protobuf:"bytes,11,opt,name=nominatedNodeName" bson:"nominatedNodeName"`

	// IP address of the host to which the pod is assigned. Empty if not yet scheduled.
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP" bson:"hostIP"`
	// IP address allocated to the pod. Routable at least within the cluster.
	// Empty if not yet allocated.
	// +optional
	PodIP string `json:"podIP,omitempty" protobuf:"bytes,6,opt,name=podIP" bson:"podIP"`

	// RFC 3339 date and time at which the object was acknowledged by the Kubelet.
	// This is before the Kubelet pulled the container image(s) for the pod.
	// +optional
	StartTime *Time `json:"startTime,omitempty" protobuf:"bytes,7,opt,name=startTime" bson:"startTime"`

	// The list has one entry per init container in the manifest. The most recent successful
	// init container will have ready = true, the most recently started container will have
	// startTime set.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	InitContainerStatuses []ContainerStatus `json:"initContainerStatuses,omitempty" protobuf:"bytes,10,rep,name=initContainerStatuses" bson:"initContainerStatuses"`

	// The list has one entry per container in the manifest. Each entry is currently the output
	// of `docker inspect`.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status
	// +optional
	ContainerStatuses []ContainerStatus `json:"containerStatuses,omitempty" protobuf:"bytes,8,rep,name=containerStatuses" bson:"containerStatuses"`
	// The Quality of Service (QOS) classification assigned to the pod based on resource requirements
	// See PodQOSClass type for available QOS classes
	// More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md
	// +optional
	QOSClass PodQOSClass `json:"qosClass,omitempty" protobuf:"bytes,9,rep,name=qosClass" bson:"qosClass"`
}

// Represents a Rados Block Device mount that lasts the lifetime of a pod.
// RBD volumes support ownership management and SELinux relabeling.
type RBDPersistentVolumeSource struct {
	// A collection of Ceph monitors.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	CephMonitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors" bson:"monitors"`
	// The rados image name.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	RBDImage string `json:"image" protobuf:"bytes,2,opt,name=image" bson:"image"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType" bson:"fsType"`
	// The rados pool name.
	// Default is rbd.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	RBDPool string `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool" bson:"pool"`
	// The rados user name.
	// Default is admin.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	RadosUser string `json:"user,omitempty" protobuf:"bytes,5,opt,name=user" bson:"user"`
	// Keyring is the path to key ring for RBDUser.
	// Default is /etc/ceph/keyring.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	Keyring string `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring" bson:"keyring"`
	// SecretRef is name of the authentication secret for RBDUser. If provided
	// overrides keyring.
	// Default is nil.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef" bson:"secretRef"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly" bson:"readOnly"`
}
type SecretReference struct {
	// Name is unique within a namespace to reference a secret resource.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`
	// Namespace defines the space within which the secret name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace" bson:"namespace"`
}

// Represents a cinder volume resource in Openstack.
// A Cinder volume must exist before mounting to a container.
// The volume must also be in the same region as the kubelet.
// Cinder volumes support ownership management and SELinux relabeling.
type CinderPersistentVolumeSource struct {
	// volume id used to identify the volume in cinder
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID" bson:"volumeID"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
	// Optional: points to a secret object containing parameters used to connect
	// to OpenStack.
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef" bson:"secretRef"`
}

type ISCSIPersistentVolumeSource struct {
	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal" bson:"targetPortal"`
	// Target iSCSI Qualified Name.
	IQN string `json:"iqn" protobuf:"bytes,2,opt,name=iqn" bson:"iqn"`
	// iSCSI Target Lun number.
	Lun int32 `json:"lun" protobuf:"varint,3,opt,name=lun" bson:"lun"`
	// iSCSI Interface Name that uses an iSCSI transport.
	// Defaults to 'default' (tcp).
	// +optional
	ISCSIInterface string `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface" bson:"iscsiInterface"`
	// Filesystem type of the volume that you want to mount.
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType" bson:"fsType"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts.
	// Defaults to false.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly" bson:"readOnly"`
	// iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	// +optional
	Portals []string `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals" bson:"portals"`
	// whether support iSCSI Discovery CHAP authentication
	// +optional
	DiscoveryCHAPAuth bool `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery" bson:"chapAuthDiscovery"`
	// whether support iSCSI Session CHAP authentication
	// +optional
	SessionCHAPAuth bool `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession" bson:"chapAuthSession"`
	// CHAP Secret for iSCSI target and initiator authentication
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef" bson:"secretRef"`
	// Custom iSCSI Initiator Name.
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	// +optional
	InitiatorName *string `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName" bson:"initiatorName"`
}

type CephFSPersistentVolumeSource struct {
	// Required: Monitors is a collection of Ceph monitors
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	Monitors []string `json:"monitors" protobuf:"bytes,1,rep,name=monitors" bson:"monitors"`
	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,2,opt,name=path" bson:"path"`
	// Optional: User is the rados user name, default is admin
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,3,opt,name=user" bson:"user"`
	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretFile string `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile" bson:"secretFile"`
	// Optional: SecretRef is reference to the authentication secret for User, default is empty.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef" bson:"secretRef"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly" bson:"readOnly"`
}

type FlexPersistentVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver" bson:"driver"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
	// Optional: SecretRef is reference to the secret object containing
	// sensitive information to pass to the plugin scripts. This may be
	// empty if no secret object is specified. If the secret object
	// contains more than one secret, all secrets are passed to the plugin
	// scripts.
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef" bson:"secretRef"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
	// Optional: Extra command options if any.
	// +optional
	Options map[string]string `json:"options,omitempty" protobuf:"bytes,5,rep,name=options" bson:"options"`
}

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFilePersistentVolumeSource struct {
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName" protobuf:"bytes,1,opt,name=secretName" bson:"secretName"`
	// Share Name
	ShareName string `json:"shareName" protobuf:"bytes,2,opt,name=shareName" bson:"shareName"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`
	// the namespace of the secret that contains Azure Storage Account Name and Key
	// default is the same as the Pod
	// +optional
	SecretNamespace *string `json:"secretNamespace" protobuf:"bytes,4,opt,name=secretNamespace" bson:"secretNamespace"`
}

type ScaleIOPersistentVolumeSource struct {
	// The host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway" protobuf:"bytes,1,opt,name=gateway" bson:"gateway"`
	// The name of the storage system as configured in ScaleIO.
	System string `json:"system" protobuf:"bytes,2,opt,name=system" bson:"system"`
	// SecretRef references to the secret for ScaleIO user and other
	// sensitive information. If this is not provided, Login operation will fail.
	SecretRef *SecretReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef" bson:"secretRef"`
	// Flag to enable/disable SSL communication with Gateway, default false
	// +optional
	SSLEnabled bool `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled" bson:"sslEnabled"`
	// The name of the ScaleIO Protection Domain for the configured storage.
	// +optional
	ProtectionDomain string `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain" bson:"protectionDomain"`
	// The ScaleIO Storage Pool associated with the protection domain.
	// +optional
	StoragePool string `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool" bson:"storagePool"`
	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
	// Default is ThinProvisioned.
	// +optional
	StorageMode string `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode" bson:"storageMode"`
	// The name of a volume already created in the ScaleIO system
	// that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName" bson:"volumeName"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// Default is "xfs"
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly" bson:"readOnly"`
}

type LocalVolumeSource struct {
	// The full path to the volume on the node.
	// It can be either a directory or block device (disk, partition, ...).
	Path string `json:"path" protobuf:"bytes,1,opt,name=path" bson:"path"`

	// Filesystem type to mount.
	// It applies only when the Path is a block device.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a fileystem if unspecified.
	// +optional
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType" bson:"fsType"`
}
type ObjectReference struct {
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind" bson:"kind"`
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace" bson:"namespace"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name" bson:"name"`
	// UID of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	// +optional
	UID UID `json:"uid,omitempty" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID" bson:"uid"`
	// API version of the referent.
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,5,opt,name=apiVersion" bson:"apiVersion"`
	// Specific resourceVersion to which this reference is made, if any.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion" bson:"resourceVersion"`

	// If referring to a piece of an object instead of an entire object, this string
	// should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
	// For example, if the object reference is to a container within a pod, this would take on a value like:
	// "spec.containers{name}" (where "name" refers to the name of the container that triggered
	// the event) or if no container name is specified "spec.containers[2]" (container with
	// index 2 in this pod). This syntax is chosen only to have some well-defined way of
	// referencing a part of an object.
	// TODO: this design is not final and this field is subject to change in the future.
	// +optional
	FieldPath string `json:"fieldPath,omitempty" protobuf:"bytes,7,opt,name=fieldPath"`
}

type StorageOSPersistentVolumeSource struct {
	// VolumeName is the human-readable name of the StorageOS volume.  Volume
	// names are only unique within a namespace.
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName" bson:"volumeName"`
	// VolumeNamespace specifies the scope of the volume within StorageOS.  If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	// +optional
	VolumeNamespace string `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace" bson:"volumeNamespace"`
	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType" bson:"fsType"`
	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly" bson:"readOnly"`
	// SecretRef specifies the secret to use for obtaining the StorageOS API
	// credentials.  If not specified, default values will be attempted.
	// +optional
	SecretRef *ObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef" bson:"secretRef"`
}
type CSIPersistentVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	// Required.
	Driver string `json:"driver" protobuf:"bytes,1,opt,name=driver" bson:"driver"`

	// VolumeHandle is the unique volume name returned by the CSI volume
	// plugin’s CreateVolume to refer to the volume on all subsequent calls.
	// Required.
	VolumeHandle string `json:"volumeHandle" protobuf:"bytes,2,opt,name=volumeHandle" bson:"volumeHandle"`

	// Optional: The value to pass to ControllerPublishVolumeRequest.
	// Defaults to false (read/write).
	// +optional
	ReadOnly bool `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly" bson:"readOnly"`

	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs".
	// +optional
	FSType string `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType" bson:"fsType"`

	// Attributes of the volume to publish.
	// +optional
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty" protobuf:"bytes,5,rep,name=volumeAttributes" bson:"volumeAttributes"`

	// ControllerPublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// ControllerPublishVolume and ControllerUnpublishVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	ControllerPublishSecretRef *SecretReference `json:"controllerPublishSecretRef,omitempty" protobuf:"bytes,6,opt,name=controllerPublishSecretRef" bson:"controllerPublishSecretRef"`

	// NodeStageSecretRef is a reference to the secret object containing sensitive
	// information to pass to the CSI driver to complete the CSI NodeStageVolume
	// and NodeStageVolume and NodeUnstageVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodeStageSecretRef *SecretReference `json:"nodeStageSecretRef,omitempty" protobuf:"bytes,7,opt,name=nodeStageSecretRef" bson:"nodeStageSecretRef"`

	// NodePublishSecretRef is a reference to the secret object containing
	// sensitive information to pass to the CSI driver to complete the CSI
	// NodePublishVolume and NodeUnpublishVolume calls.
	// This field is optional, and  may be empty if no secret is required. If the
	// secret object contains more than one secret, all secrets are passed.
	// +optional
	NodePublishSecretRef *SecretReference `json:"nodePublishSecretRef,omitempty" protobuf:"bytes,8,opt,name=nodePublishSecretRef" bson:"nodePublishSecretRef"`
}
type PersistentVolumeSource struct {
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GCEPersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" protobuf:"bytes,1,opt,name=gcePersistentDisk" bson:"gcePersistentDisk"`
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,2,opt,name=awsElasticBlockStore" bson:"awsElasticBlockStore"`
	// HostPath represents a directory on the host.
	// Provisioned by a developer or tester.
	// This is useful for single-node development and testing only!
	// On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" protobuf:"bytes,3,opt,name=hostPath" bson:"hostPath"`
	// Glusterfs represents a Glusterfs volume that is attached to a host and
	// exposed to the pod. Provisioned by an admin.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md
	// +optional
	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" protobuf:"bytes,4,opt,name=glusterfs" bson:"glusterfs"`
	// NFS represents an NFS mount on the host. Provisioned by an admin.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	NFS *NFSVolumeSource `json:"nfs,omitempty" protobuf:"bytes,5,opt,name=nfs" bson:"nfs"`
	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md
	// +optional
	RBD *RBDPersistentVolumeSource `json:"rbd,omitempty" protobuf:"bytes,6,opt,name=rbd" bson:"rbd"`
	// ISCSI represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// +optional
	ISCSI *ISCSIPersistentVolumeSource `json:"iscsi,omitempty" protobuf:"bytes,7,opt,name=iscsi" bson:"iscsi"`
	// Cinder represents a cinder volume attached and mounted on kubelets host machine
	// More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	// +optional
	Cinder *CinderPersistentVolumeSource `json:"cinder,omitempty" protobuf:"bytes,8,opt,name=cinder" bson:"cinder"`
	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// +optional
	CephFS *CephFSPersistentVolumeSource `json:"cephfs,omitempty" protobuf:"bytes,9,opt,name=cephfs" bson:"cephfs"`
	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	FC *FCVolumeSource `json:"fc,omitempty" protobuf:"bytes,10,opt,name=fc" bson:"fc"`
	// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
	// +optional
	Flocker *FlockerVolumeSource `json:"flocker,omitempty" protobuf:"bytes,11,opt,name=flocker" bson:"flocker"`
	// FlexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// +optional
	FlexVolume *FlexPersistentVolumeSource `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume" bson:"flexVolume"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// +optional
	AzureFile *AzureFilePersistentVolumeSource `json:"azureFile,omitempty" protobuf:"bytes,13,opt,name=azureFile" bson:"azureFile"`
	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// +optional
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty" protobuf:"bytes,14,opt,name=vsphereVolume" bson:"vsphereVolume"`
	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// +optional
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty" protobuf:"bytes,15,opt,name=quobyte" bson:"quobyte"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// +optional
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty" protobuf:"bytes,16,opt,name=azureDisk" bson:"azureDisk"`
	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" protobuf:"bytes,17,opt,name=photonPersistentDisk" bson:"photonPersistentDisk"`
	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// +optional
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty" protobuf:"bytes,18,opt,name=portworxVolume" bson:"portworxVolume"`
	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// +optional
	ScaleIO *ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty" protobuf:"bytes,19,opt,name=scaleIO" bson:"scaleIO"`
	// Local represents directly-attached storage with node affinity
	// +optional
	Local *LocalVolumeSource `json:"local,omitempty" protobuf:"bytes,20,opt,name=local" bson:"local"`
	// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod
	// More info: https://releases.k8s.io/HEAD/examples/volumes/storageos/README.md
	// +optional
	StorageOS *StorageOSPersistentVolumeSource `json:"storageos,omitempty" protobuf:"bytes,21,opt,name=storageos" bson:"storageos"`
	// CSI represents storage that handled by an external CSI driver (Beta feature).
	// +optional
	CSI *CSIPersistentVolumeSource `json:"csi,omitempty" protobuf:"bytes,22,opt,name=csi" bson:"csi"`
}

type PersistentVolumeAccessMode string

const (
	// can be mounted read/write mode to exactly 1 host
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	// can be mounted in read-only mode to many hosts
	ReadOnlyMany PersistentVolumeAccessMode = "ReadOnlyMany"
	// can be mounted in read/write mode to many hosts
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
)

type PersistentVolumePhase string

const (
	// used for PersistentVolumes that are not available
	VolumePending PersistentVolumePhase = "Pending"
	// used for PersistentVolumes that are not yet bound
	// Available volumes are held by the binder and matched to PersistentVolumeClaims
	VolumeAvailable PersistentVolumePhase = "Available"
	// used for PersistentVolumes that are bound
	VolumeBound PersistentVolumePhase = "Bound"
	// used for PersistentVolumes where the bound PersistentVolumeClaim was deleted
	// released volumes must be recycled before becoming available again
	// this phase is used by the persistent volume claim binder to signal to another process to reclaim the resource
	VolumeReleased PersistentVolumePhase = "Released"
	// used for PersistentVolumes that failed to be correctly recycled or deleted after being released from a claim
	VolumeFailed PersistentVolumePhase = "Failed"
)

type PersistentVolumeClaimPhase string

const (
	// used for PersistentVolumeClaims that are not yet bound
	ClaimPending PersistentVolumeClaimPhase = "Pending"
	// used for PersistentVolumeClaims that are bound
	ClaimBound PersistentVolumeClaimPhase = "Bound"
	// used for PersistentVolumeClaims that lost their underlying
	// PersistentVolume. The claim was bound to a PersistentVolume and this
	// volume does not exist any longer and all data on it was lost.
	ClaimLost PersistentVolumeClaimPhase = "Lost"
)

type PersistentVolumeReclaimPolicy string

const (
	// PersistentVolumeReclaimRecycle means the volume will be recycled back into the pool of unbound persistent volumes on release from its claim.
	// The volume plugin must support Recycling.
	PersistentVolumeReclaimRecycle PersistentVolumeReclaimPolicy = "Recycle"
	// PersistentVolumeReclaimDelete means the volume will be deleted from Kubernetes on release from its claim.
	// The volume plugin must support Deletion.
	PersistentVolumeReclaimDelete PersistentVolumeReclaimPolicy = "Delete"
	// PersistentVolumeReclaimRetain means the volume will be left in its current phase (Released) for manual reclamation by the administrator.
	// The default policy is Retain.
	PersistentVolumeReclaimRetain PersistentVolumeReclaimPolicy = "Retain"
)

// PersistentVolumeMode describes how a volume is intended to be consumed, either Block or Filesystem.
type PersistentVolumeMode string

const (
	// PersistentVolumeBlock means the volume will not be formatted with a filesystem and will remain a raw block device.
	PersistentVolumeBlock PersistentVolumeMode = "Block"
	// PersistentVolumeFilesystem means the volume will be or is formatted with a filesystem.
	PersistentVolumeFilesystem PersistentVolumeMode = "Filesystem"
)

type VolumeNodeAffinity struct {
	// Required specifies hard node constraints that must be met.
	Required *NodeSelector `json:"required,omitempty" protobuf:"bytes,1,opt,name=required" bson:"required"`
}

type PersistentVolumeSpec struct {
	// A description of the persistent volume's resources and capacity.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// +optional
	Capacity ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName" bson:"capacity"`
	// The actual volume backing the persistent volume.
	PersistentVolumeSource `json:",inline" protobuf:"bytes,2,opt,name=persistentVolumeSource"`
	// AccessModes contains all ways the volume can be mounted.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	// +optional
	AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode" bson:"accessModes"`
	// ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	// Expected to be non-nil when bound.
	// claim.VolumeName is the authoritative bind between PV and PVC.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	// +optional
	ClaimRef *ObjectReference `json:"claimRef,omitempty" protobuf:"bytes,4,opt,name=claimRef" bson:"claimRef"`
	// What happens to a persistent volume when released from its claim.
	// Valid options are Retain (default for manually created PersistentVolumes), Delete (default
	// for dynamically provisioned PersistentVolumes), and Recycle (deprecated).
	// Recycle must be supported by the volume plugin underlying this PersistentVolume.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	// +optional
	PersistentVolumeReclaimPolicy PersistentVolumeReclaimPolicy `json:"persistentVolumeReclaimPolicy,omitempty" protobuf:"bytes,5,opt,name=persistentVolumeReclaimPolicy,casttype=PersistentVolumeReclaimPolicy" bson:"persistentVolumeReclaimPolicy"`
	// Name of StorageClass to which this persistent volume belongs. Empty value
	// means that this volume does not belong to any StorageClass.
	// +optional
	StorageClassName string `json:"storageClassName,omitempty" protobuf:"bytes,6,opt,name=storageClassName" bson:"storageClassName"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will
	// simply fail if one is invalid.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	// +optional
	MountOptions []string `json:"mountOptions,omitempty" protobuf:"bytes,7,opt,name=mountOptions" bson:"mountOptions"`
	// volumeMode defines if a volume is intended to be used with a formatted filesystem
	// or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	// This is an alpha feature and may change in the future.
	// +optional
	VolumeMode *PersistentVolumeMode `json:"volumeMode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode" bson:"volumeMode"`
	// NodeAffinity defines constraints that limit what nodes this volume can be accessed from.
	// This field influences the scheduling of pods that use this volume.
	// +optional
	NodeAffinity *VolumeNodeAffinity `json:"nodeAffinity,omitempty" protobuf:"bytes,9,opt,name=nodeAffinity" bson:"nodeAffinity"`
}

type PersistentVolumeStatus struct {
	// Phase indicates if a volume is available, bound to a claim, or released by a claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	// +optional
	Phase PersistentVolumePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumePhase" bson:"phase"`
	// A human-readable message indicating details about why the volume is in this state.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message" bson:"message"`
	// Reason is a brief CamelCase string that describes any failure and is meant
	// for machine parsing and tidy display in the CLI.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason" bson:"reason"`
}
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup" bson:"apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind" bson:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name" protobuf:"bytes,3,opt,name=name" bson:"name"`
}

type PersistentVolumeClaimSpec struct {
	// AccessModes contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode" bson:"accessModes"`
	// A label query over volumes to consider for binding.
	// +optional
	Selector *LabelSelector `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector" bson:"selector"`
	// Resources represents the minimum resources the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources" bson:"resources"`
	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	// +optional
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,3,opt,name=volumeName" bson:"volumeName"`
	// Name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty" protobuf:"bytes,5,opt,name=storageClassName" bson:"storageClassName"`
	// volumeMode defines what type of volume is required by the claim.
	// Value of Filesystem is implied when not included in claim spec.
	// This is an alpha feature and may change in the future.
	// +optional
	VolumeMode *PersistentVolumeMode `json:"volumeMode,omitempty" protobuf:"bytes,6,opt,name=volumeMode,casttype=PersistentVolumeMode" bson:"volumeMode"`
	// This field requires the VolumeSnapshotDataSource alpha feature gate to be
	// enabled and currently VolumeSnapshot is the only supported data source.
	// If the provisioner can support VolumeSnapshot data source, it will create
	// a new volume and data will be restored to the volume at the same time.
	// If the provisioner does not support VolumeSnapshot data source, volume will
	// not be created and the failure will be reported as an event.
	// In the future, we plan to support more data source types and the behavior
	// of the provisioner may change.
	// +optional
	DataSource *TypedLocalObjectReference `json:"dataSource" protobuf:"bytes,7,opt,name=dataSource" bson:"dataSource"`
}

// PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type
type PersistentVolumeClaimConditionType string

const (
	// PersistentVolumeClaimResizing - a user trigger resize of pvc has been started
	PersistentVolumeClaimResizing PersistentVolumeClaimConditionType = "Resizing"
	// PersistentVolumeClaimFileSystemResizePending - controller resize is finished and a file system resize is pending on node
	PersistentVolumeClaimFileSystemResizePending PersistentVolumeClaimConditionType = "FileSystemResizePending"
)

type PersistentVolumeClaimStatus struct {
	// Phase represents the current phase of PersistentVolumeClaim.
	// +optional
	Phase PersistentVolumeClaimPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumeClaimPhase" bson:"phase"`
	// AccessModes contains the actual access modes the volume backing the PVC has.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,2,rep,name=accessModes,casttype=PersistentVolumeAccessMode" bson:"accessModes"`
	// Represents the actual resources of the underlying volume.
	// +optional
	Capacity ResourceList `json:"capacity,omitempty" protobuf:"bytes,3,rep,name=capacity,casttype=ResourceList,castkey=ResourceName" bson:"capacity"`
	// Current Condition of persistent volume claim. If underlying persistent volume is being
	// resized then the Condition will be set to 'ResizeStarted'.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []PersistentVolumeClaimCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions" bson:"conditions"`
}

type PersistentVolumeClaimCondition struct {
	Type   PersistentVolumeClaimConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PersistentVolumeClaimConditionType" bson:"type"`
	Status ConditionStatus                    `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus" bson:"status"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime" bson:"lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// Unique, this should be a short, machine understandable string that gives the reason
	// for condition's last transition. If it reports "ResizeStarted" that means the underlying
	// persistent volume is being resized.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason" bson:"reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message" bson:"message"`
}

type ReplicaSetSpec struct {
	// Replicas is the number of desired replicas.
	// This is a pointer to distinguish between explicit zero and unspecified.
	// Defaults to 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas" bson:"replicas"`

	// Minimum number of seconds for which a newly created pod should be ready
	// without any of its container crashing, for it to be considered available.
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds" bson:"minReadySeconds"`

	// Selector is a label query over pods that should match the replica count.
	// Label keys and values that must match in order to be controlled by this replica set.
	// It must match the pod template's labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector" bson:"selector"`

	// Template is the object that describes the pod that will be created if
	// insufficient replicas are detected.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// +optional
	Template PodTemplateSpec `json:"template,omitempty" protobuf:"bytes,3,opt,name=template" bson:"template"`
}

type ReplicaSetStatus struct {
	// Replicas is the most recently oberved number of replicas.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas int32 `json:"replicas" protobuf:"varint,1,opt,name=replicas" bson:"replicas"`

	// The number of pods that have labels matching the labels of the pod template of the replicaset.
	// +optional
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty" protobuf:"varint,2,opt,name=fullyLabeledReplicas" bson:"fullyLabeledReplicas"`

	// The number of ready replicas for this replica set.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,4,opt,name=readyReplicas" bson:"readyReplicas"`

	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty" protobuf:"varint,5,opt,name=availableReplicas" bson:"availableReplicas"`

	// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration" bson:"observedGeneration"`

	// Represents the latest available observations of a replica set's current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []ReplicaSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions" bson:"conditions"`
}

type ReplicaSetConditionType string

// These are valid conditions of a replica set.
const (
	// ReplicaSetReplicaFailure is added in a replica set when one of its pods fails to be created
	// due to insufficient quota, limit ranges, pod security policy, node selectors, etc. or deleted
	// due to kubelet being down or finalizers are failing.
	ReplicaSetReplicaFailure ReplicaSetConditionType = "ReplicaFailure"
)

type ReplicaSetCondition struct {
	// Type of replica set condition.
	Type ReplicaSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ReplicaSetConditionType" bson:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus" bson:"status"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason" bson:"reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message" bson:"message"`
}

type SecretType string

const (
	// SecretTypeOpaque is the default. Arbitrary user-defined data
	SecretTypeOpaque SecretType = "Opaque"

	// SecretTypeServiceAccountToken contains a token that identifies a service account to the API
	//
	// Required fields:
	// - Secret.Annotations["kubernetes.io/service-account.name"] - the name of the ServiceAccount the token identifies
	// - Secret.Annotations["kubernetes.io/service-account.uid"] - the UID of the ServiceAccount the token identifies
	// - Secret.Data["token"] - a token that identifies the service account to the API
	SecretTypeServiceAccountToken SecretType = "kubernetes.io/service-account-token"

	// ServiceAccountNameKey is the key of the required annotation for SecretTypeServiceAccountToken secrets
	ServiceAccountNameKey = "kubernetes.io/service-account.name"
	// ServiceAccountUIDKey is the key of the required annotation for SecretTypeServiceAccountToken secrets
	ServiceAccountUIDKey = "kubernetes.io/service-account.uid"
	// ServiceAccountTokenKey is the key of the required data for SecretTypeServiceAccountToken secrets
	ServiceAccountTokenKey = "token"
	// ServiceAccountKubeconfigKey is the key of the optional kubeconfig data for SecretTypeServiceAccountToken secrets
	ServiceAccountKubeconfigKey = "kubernetes.kubeconfig"
	// ServiceAccountRootCAKey is the key of the optional root certificate authority for SecretTypeServiceAccountToken secrets
	ServiceAccountRootCAKey = "ca.crt"
	// ServiceAccountNamespaceKey is the key of the optional namespace to use as the default for namespaced API calls
	ServiceAccountNamespaceKey = "namespace"

	// SecretTypeDockercfg contains a dockercfg file that follows the same format rules as ~/.dockercfg
	//
	// Required fields:
	// - Secret.Data[".dockercfg"] - a serialized ~/.dockercfg file
	SecretTypeDockercfg SecretType = "kubernetes.io/dockercfg"

	// DockerConfigKey is the key of the required data for SecretTypeDockercfg secrets
	DockerConfigKey = ".dockercfg"

	// SecretTypeDockerConfigJson contains a dockercfg file that follows the same format rules as ~/.docker/config.json
	//
	// Required fields:
	// - Secret.Data[".dockerconfigjson"] - a serialized ~/.docker/config.json file
	SecretTypeDockerConfigJson SecretType = "kubernetes.io/dockerconfigjson"

	// DockerConfigJsonKey is the key of the required data for SecretTypeDockerConfigJson secrets
	DockerConfigJsonKey = ".dockerconfigjson"

	// SecretTypeBasicAuth contains data needed for basic authentication.
	//
	// Required at least one of fields:
	// - Secret.Data["username"] - username used for authentication
	// - Secret.Data["password"] - password or token needed for authentication
	SecretTypeBasicAuth SecretType = "kubernetes.io/basic-auth"

	// BasicAuthUsernameKey is the key of the username for SecretTypeBasicAuth secrets
	BasicAuthUsernameKey = "username"
	// BasicAuthPasswordKey is the key of the password or token for SecretTypeBasicAuth secrets
	BasicAuthPasswordKey = "password"

	// SecretTypeSSHAuth contains data needed for SSH authetication.
	//
	// Required field:
	// - Secret.Data["ssh-privatekey"] - private SSH key needed for authentication
	SecretTypeSSHAuth SecretType = "kubernetes.io/ssh-auth"

	// SSHAuthPrivateKey is the key of the required SSH private key for SecretTypeSSHAuth secrets
	SSHAuthPrivateKey = "ssh-privatekey"
	// SecretTypeTLS contains information about a TLS client or server secret. It
	// is primarily used with TLS termination of the Ingress resource, but may be
	// used in other types.
	//
	// Required fields:
	// - Secret.Data["tls.key"] - TLS private key.
	//   Secret.Data["tls.crt"] - TLS certificate.
	// TODO: Consider supporting different formats, specifying CA/destinationCA.
	SecretTypeTLS SecretType = "kubernetes.io/tls"

	// TLSCertKey is the key for tls certificates in a TLS secert.
	TLSCertKey = "tls.crt"
	// TLSPrivateKeyKey is the key for the private key field in a TLS secret.
	TLSPrivateKeyKey = "tls.key"
)

type ServiceType string

const (
	// ServiceTypeClusterIP means a service will only be accessible inside the
	// cluster, via the cluster IP.
	ServiceTypeClusterIP ServiceType = "ClusterIP"

	// ServiceTypeNodePort means a service will be exposed on one port of
	// every node, in addition to 'ClusterIP' type.
	ServiceTypeNodePort ServiceType = "NodePort"

	// ServiceTypeLoadBalancer means a service will be exposed via an
	// external load balancer (if the cloud provider supports it), in addition
	// to 'NodePort' type.
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"

	// ServiceTypeExternalName means a service consists of only a reference to
	// an external name that kubedns or equivalent will return as a CNAME
	// record, with no exposing or proxying of any pods involved.
	ServiceTypeExternalName ServiceType = "ExternalName"
)

type ServiceAffinity string

const (
	// ServiceAffinityClientIP is the Client IP based.
	ServiceAffinityClientIP ServiceAffinity = "ClientIP"

	// ServiceAffinityNone - no session affinity.
	ServiceAffinityNone ServiceAffinity = "None"
)

type ServiceExternalTrafficPolicyType string

const (
	// ServiceExternalTrafficPolicyTypeLocal specifies node-local endpoints behavior.
	ServiceExternalTrafficPolicyTypeLocal ServiceExternalTrafficPolicyType = "Local"
	// ServiceExternalTrafficPolicyTypeCluster specifies node-global (legacy) behavior.
	ServiceExternalTrafficPolicyTypeCluster ServiceExternalTrafficPolicyType = "Cluster"
)

type SessionAffinityConfig struct {
	// clientIP contains the configurations of Client IP based session affinity.
	// +optional
	ClientIP *ClientIPConfig `json:"clientIP,omitempty" protobuf:"bytes,1,opt,name=clientIP" bson:"clientIP"`
}

// ClientIPConfig represents the configurations of Client IP based session affinity.
type ClientIPConfig struct {
	// timeoutSeconds specifies the seconds of ClientIP type session sticky time.
	// The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP".
	// Default value is 10800(for 3 hours).
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,1,opt,name=timeoutSeconds" bson:"timeoutSeconds"`
}

type ServiceSpec struct {
	// The list of ports that are exposed by this service.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +patchMergeKey=port
	// +patchStrategy=merge
	Ports []ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports" bson:"ports"`

	// Route service traffic to pods with label keys and values matching this
	// selector. If empty or not present, the service is assumed to have an
	// external process managing its endpoints, which Kubernetes will not
	// modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
	// Ignored if type is ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/
	// +optional
	Selector map[string]string `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector" bson:"selector"`

	// clusterIP is the IP address of the service and is usually assigned
	// randomly by the master. If an address is specified manually and is not in
	// use by others, it will be allocated to the service; otherwise, creation
	// of the service will fail. This field can not be changed through updates.
	// Valid values are "None", empty string (""), or a valid IP address. "None"
	// can be specified for headless services when proxying is not required.
	// Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if
	// type is ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	ClusterIP string `json:"clusterIP,omitempty" protobuf:"bytes,3,opt,name=clusterIP" bson:"clusterIP"`

	// type determines how the Service is exposed. Defaults to ClusterIP. Valid
	// options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
	// "ExternalName" maps to the specified externalName.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object. If clusterIP is
	// "None", no virtual IP is allocated and the endpoints are published as a
	// set of endpoints rather than a stable IP.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an
	// external load-balancer (if supported in the current cloud) which routes
	// to the clusterIP.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services---service-types
	// +optional
	Type ServiceType `json:"type,omitempty" protobuf:"bytes,4,opt,name=type,casttype=ServiceType" bson:"type"`

	// externalIPs is a list of IP addresses for which nodes in the cluster
	// will also accept traffic for this service.  These IPs are not managed by
	// Kubernetes.  The user is responsible for ensuring that traffic arrives
	// at a node with this IP.  A common example is external load-balancers
	// that are not part of the Kubernetes system.
	// +optional
	ExternalIPs []string `json:"externalIPs,omitempty" protobuf:"bytes,5,rep,name=externalIPs" bson:"externalIPs"`

	// Supports "ClientIP" and "None". Used to maintain session affinity.
	// Enable client IP based session affinity.
	// Must be ClientIP or None.
	// Defaults to None.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	SessionAffinity ServiceAffinity `json:"sessionAffinity,omitempty" protobuf:"bytes,7,opt,name=sessionAffinity,casttype=ServiceAffinity" bson:"sessionAffinity"`

	// Only applies to Service Type: LoadBalancer
	// LoadBalancer will get created with the IP specified in this field.
	// This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created.
	// This field will be ignored if the cloud-provider does not support the feature.
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty" protobuf:"bytes,8,opt,name=loadBalancerIP" bson:"loadBalancerIP"`

	// If specified and supported by the platform, this will restrict traffic through the cloud-provider
	// load-balancer will be restricted to the specified client IPs. This field will be ignored if the
	// cloud-provider does not support the feature."
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty" protobuf:"bytes,9,opt,name=loadBalancerSourceRanges" bson:"loadBalancerSourceRanges"`

	// externalName is the external reference that kubedns or equivalent will
	// return as a CNAME record for this service. No proxying will be involved.
	// Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123)
	// and requires Type to be ExternalName.
	// +optional
	ExternalName string `json:"externalName,omitempty" protobuf:"bytes,10,opt,name=externalName" bson:"externalName"`

	// externalTrafficPolicy denotes if this Service desires to route external
	// traffic to node-local or cluster-wide endpoints. "Local" preserves the
	// client source IP and avoids a second hop for LoadBalancer and Nodeport
	// type services, but risks potentially imbalanced traffic spreading.
	// "Cluster" obscures the client source IP and may cause a second hop to
	// another node, but should have good overall load-spreading.
	// +optional
	ExternalTrafficPolicy ServiceExternalTrafficPolicyType `json:"externalTrafficPolicy,omitempty" protobuf:"bytes,11,opt,name=externalTrafficPolicy" bson:"externalTrafficPolicy"`

	// healthCheckNodePort specifies the healthcheck nodePort for the service.
	// If not specified, HealthCheckNodePort is created by the service api
	// backend with the allocated nodePort. Will use user-specified nodePort value
	// if specified by the client. Only effects when Type is set to LoadBalancer
	// and ExternalTrafficPolicy is set to Local.
	// +optional
	HealthCheckNodePort int32 `json:"healthCheckNodePort,omitempty" protobuf:"bytes,12,opt,name=healthCheckNodePort" bson:"healthCheckNodePort"`

	// publishNotReadyAddresses, when set to true, indicates that DNS implementations
	// must publish the notReadyAddresses of subsets for the Endpoints associated with
	// the Service. The default value is false.
	// The primary use case for setting this field is to use a StatefulSet's Headless Service
	// to propagate SRV records for its Pods without respect to their readiness for purpose
	// of peer discovery.
	// +optional
	PublishNotReadyAddresses bool `json:"publishNotReadyAddresses,omitempty" protobuf:"varint,13,opt,name=publishNotReadyAddresses" bson:"publishNotReadyAddresses"`
	// sessionAffinityConfig contains the configurations of session affinity.
	// +optional
	SessionAffinityConfig *SessionAffinityConfig `json:"sessionAffinityConfig,omitempty" protobuf:"bytes,14,opt,name=sessionAffinityConfig" bson:"sessionAffinityConfig"`
}

// ServicePort contains information on service's port.
type ServicePort struct {
	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names. This maps to
	// the 'Name' field in EndpointPort objects.
	// Optional if only one ServicePort is defined on this service.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name" bson:"name"`

	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	// Default is TCP.
	// +optional
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol,casttype=Protocol" bson:"protocol"`

	// The port that will be exposed by this service.
	Port int32 `json:"port" protobuf:"varint,3,opt,name=port" bson:"port"`

	// Number or name of the port to access on the pods targeted by the service.
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// If this is a string, it will be looked up as a named port in the
	// target Pod's container ports. If this is not specified, the value
	// of the 'port' field is used (an identity map).
	// This field is ignored for services with clusterIP=None, and should be
	// omitted or set equal to the 'port' field.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	// +optional
	TargetPort intstr.IntOrString `json:"targetPort,omitempty" protobuf:"bytes,4,opt,name=targetPort" bson:"targetPort"`

	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	// Usually assigned by the system. If specified, it will be allocated to the service
	// if unused or else creation of the service will fail.
	// Default is to auto-allocate a port if the ServiceType of this Service requires one.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	// +optional
	NodePort int32 `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort" bson:"nodePort"`
}
type ServiceStatus struct {
	// LoadBalancer contains the current status of the load-balancer,
	// if one is present.
	// +optional
	LoadBalancer LoadBalancerStatus `json:"loadBalancer,omitempty" protobuf:"bytes,1,opt,name=loadBalancer" bson:"loadBalancer"`
}
type PodManagementPolicyType string

const (
	// OrderedReadyPodManagement will create pods in strictly increasing order on
	// scale up and strictly decreasing order on scale down, progressing only when
	// the previous pod is ready or terminated. At most one pod will be changed
	// at any time.
	OrderedReadyPodManagement PodManagementPolicyType = "OrderedReady"
	// ParallelPodManagement will create and Delete pods as soon as the stateful set
	// replica count is changed, and will not wait for pods to be ready or complete
	// termination.
	ParallelPodManagement = "Parallel"
)

type StatefulSetUpdateStrategyType string

const (
	// RollingUpdateStatefulSetStrategyType indicates that Update will be
	// applied to all Pods in the StatefulSet with respect to the StatefulSet
	// ordering constraints. When a scale operation is performed with this
	// strategy, new Pods will be created from the specification version indicated
	// by the StatefulSet's updateRevision.
	RollingUpdateStatefulSetStrategyType = "RollingUpdate"
	// OnDeleteStatefulSetStrategyType triggers the legacy behavior. Version
	// tracking and ordered rolling restarts are disabled. Pods are recreated
	// from the StatefulSetSpec when they are manually deleted. When a scale
	// operation is performed with this strategy,specification version indicated
	// by the StatefulSet's currentRevision.
	OnDeleteStatefulSetStrategyType = "OnDelete"
)

type RollingUpdateStatefulSetStrategy struct {
	// Partition indicates the ordinal at which the StatefulSet should be
	// partitioned.
	// Default value is 0.
	// +optional
	Partition *int32 `json:"partition,omitempty" protobuf:"varint,1,opt,name=partition" bson:"partition"`
}
type StatefulSetUpdateStrategy struct {
	// Type indicates the type of the StatefulSetUpdateStrategy.
	// Default is RollingUpdate.
	// +optional
	Type StatefulSetUpdateStrategyType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=StatefulSetStrategyType" bson:"type"`
	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	// +optional
	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate" bson:"rollingUpdate"`
}
type PersistentVolumeClaimStruct struct {
	TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	Spec PersistentVolumeClaimSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`

	// Status represents the current information/status of a persistent volume claim.
	// Read-only.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	Status PersistentVolumeClaimStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type StatefulSetSpec struct {
	// replicas is the desired number of replicas of the given Template.
	// These are replicas in the sense that they are instantiations of the
	// same Template, but individual replicas also have a consistent identity.
	// If unspecified, defaults to 1.
	// TODO: Consider a rename of this field.
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas" bson:"replicas"`

	// selector is a label query over pods that should match the replica count.
	// It must match the pod template's labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector" bson:"selector"`

	// template is the object that describes the pod that will be created if
	// insufficient replicas are detected. Each pod stamped out by the StatefulSet
	// will fulfill this Template, but have a unique identity from the rest
	// of the StatefulSet.
	Template PodTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template" bson:"template"`

	// volumeClaimTemplates is a list of claims that pods are allowed to reference.
	// The StatefulSet controller is responsible for mapping network identities to
	// claims in a way that maintains the identity of a pod. Every claim in
	// this list must have at least one matching (by name) volumeMount in one
	// container in the template. A claim in this list takes precedence over
	// any volumes in the template, with the same name.
	// TODO: Define the behavior if a claim already exists with the same name.
	// +optional
	VolumeClaimTemplates []PersistentVolumeClaimStruct `json:"volumeClaimTemplates,omitempty" protobuf:"bytes,4,rep,name=volumeClaimTemplates" bson:"volumeClaimTemplates"`

	// serviceName is the name of the service that governs this StatefulSet.
	// This service must exist before the StatefulSet, and is responsible for
	// the network identity of the set. Pods get DNS/hostnames that follow the
	// pattern: pod-specific-string.serviceName.default.svc.cluster.local
	// where "pod-specific-string" is managed by the StatefulSet controller.
	ServiceName string `json:"serviceName" protobuf:"bytes,5,opt,name=serviceName" bson:"serviceName"`

	// podManagementPolicy controls how pods are created during initial scale up,
	// when replacing pods on nodes, or when scaling down. The default policy is
	// `OrderedReady`, where pods are created in increasing order (pod-0, then
	// pod-1, etc) and the controller will wait until each pod is ready before
	// continuing. When scaling down, the pods are removed in the opposite order.
	// The alternative policy is `Parallel` which will create pods in parallel
	// to match the desired scale without waiting, and on scale down will Delete
	// all pods at once.
	// +optional
	PodManagementPolicy PodManagementPolicyType `json:"podManagementPolicy,omitempty" protobuf:"bytes,6,opt,name=podManagementPolicy,casttype=PodManagementPolicyType" bson:"podManagementPolicy"`

	// updateStrategy indicates the StatefulSetUpdateStrategy that will be
	// employed to Update Pods in the StatefulSet when a revision is made to
	// Template.
	UpdateStrategy StatefulSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,7,opt,name=updateStrategy" bson:"updateStrategy"`

	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the StatefulSet's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// StatefulSetSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,8,opt,name=revisionHistoryLimit" bson:"revisionHistoryLimit"`
}

type StatefulSetStatus struct {
	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the
	// StatefulSet's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration" bson:"observedGeneration"`

	// replicas is the number of Pods created by the StatefulSet controller.
	Replicas int32 `json:"replicas" protobuf:"varint,2,opt,name=replicas" bson:"replicas"`

	// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,3,opt,name=readyReplicas" bson:"readyReplicas"`

	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version
	// indicated by currentRevision.
	CurrentReplicas int32 `json:"currentReplicas,omitempty" protobuf:"varint,4,opt,name=currentReplicas" bson:"currentReplicas"`

	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version
	// indicated by updateRevision.
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" protobuf:"varint,5,opt,name=updatedReplicas" bson:"updatedReplicas"`

	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the
	// sequence [0,currentReplicas).
	CurrentRevision string `json:"currentRevision,omitempty" protobuf:"bytes,6,opt,name=currentRevision" bson:"currentRevision"`

	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence
	// [replicas-updatedReplicas,replicas)
	UpdateRevision string `json:"updateRevision,omitempty" protobuf:"bytes,7,opt,name=updateRevision" bson:"updateRevision"`

	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller
	// uses this field as a collision avoidance mechanism when it needs to create the name for the
	// newest ControllerRevision.
	// +optional
	CollisionCount *int32 `json:"collisionCount,omitempty" protobuf:"varint,9,opt,name=collisionCount" bson:"collisionCount"`

	// Represents the latest available observations of a statefulset's current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []StatefulSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,10,rep,name=conditions" bson:"conditions"`
}

type StatefulSetConditionType string

// StatefulSetCondition describes the state of a statefulset at a certain point.
type StatefulSetCondition struct {
	// Type of statefulset condition.
	Type StatefulSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=StatefulSetConditionType" bson:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus" bson:"status"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason" bson:"reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message" bson:"message"`
}

type NamespaceStatus struct {
	Phase NamespacePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=NamespacePhase" bson:"NamespacePhase"`
}
type NamespaceSpec struct {
	Finalizers []FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,1,rep,name=finalizers,casttype=FinalizerName" bson:"finalizers"`
}

type NodeSpec struct {
	// PodCIDR represents the pod IP range assigned to the node.
	// +optional
	PodCIDR string `json:"podCIDR,omitempty" protobuf:"bytes,1,opt,name=podCIDR" bson:"podCIDR"`

	// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this
	// field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for
	// each of IPv4 and IPv6.
	// +optional
	// +patchStrategy=merge
	PodCIDRs []string `json:"podCIDRs,omitempty" protobuf:"bytes,7,opt,name=podCIDRs" patchStrategy:"merge" bson:"podCIDRs"`

	// ID of the node assigned by the cloud provider in the format: <ProviderName>://<ProviderSpecificNodeID>
	// +optional
	ProviderID string `json:"providerID,omitempty" protobuf:"bytes,3,opt,name=providerID" bson:"providerID"`
	// Unschedulable controls node schedulability of new pods. By default, node is schedulable.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
	// +optional
	Unschedulable bool `json:"unschedulable,omitempty" protobuf:"varint,4,opt,name=unschedulable" bson:"unschedulable"`
	// If specified, the node's taints.
	// +optional
	Taints []Taint `json:"taints,omitempty" protobuf:"bytes,5,opt,name=taints" bson:"taints"`
	// If specified, the source to get node configuration from
	// The DynamicKubeletConfig feature gate must be enabled for the Kubelet to use this field
	// +optional
	ConfigSource *NodeConfigSource `json:"configSource,omitempty" protobuf:"bytes,6,opt,name=configSource" bson:"configSource"`

	// Deprecated. Not all kubelets will set this field. Remove field after 1.13.
	// see: https://issues.k8s.io/61966
	// +optional
	DoNotUseExternalID string `json:"externalID,omitempty" protobuf:"bytes,2,opt,name=externalID" bson:"externalID"`
}

// NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
type NodeConfigSource struct {
	// For historical context, regarding the below kind, apiVersion, and configMapRef deprecation tags:
	// 1. kind/apiVersion were used by the kubelet to persist this struct to disk (they had no protobuf tags)
	// 2. configMapRef and proto tag 1 were used by the API to refer to a configmap,
	//    but used a generic ObjectReference type that didn't really have the fields we needed
	// All uses/persistence of the NodeConfigSource struct prior to 1.11 were gated by alpha feature flags,
	// so there was no persisted data for these fields that needed to be migrated/handled.

	// +k8s:deprecated=kind
	// +k8s:deprecated=apiVersion
	// +k8s:deprecated=configMapRef,protobuf=1

	// ConfigMap is a reference to a Node's ConfigMap
	ConfigMap *ConfigMapNodeConfigSource `json:"configMap,omitempty" protobuf:"bytes,2,opt,name=configMap" bson:"configMap"`
}

// ConfigMapNodeConfigSource contains the information to reference a ConfigMap as a config source for the Node.
type ConfigMapNodeConfigSource struct {
	// Namespace is the metadata.namespace of the referenced ConfigMap.
	// This field is required in all cases.
	Namespace string `json:"namespace" protobuf:"bytes,1,opt,name=namespace" bson:"namespace"`

	// Name is the metadata.name of the referenced ConfigMap.
	// This field is required in all cases.
	Name string `json:"name" protobuf:"bytes,2,opt,name=name" bson:"name"`

	// UID is the metadata.UID of the referenced ConfigMap.
	// This field is forbidden in Node.Spec, and required in Node.Status.
	// +optional
	UID UID `json:"uid,omitempty" protobuf:"bytes,3,opt,name=uid" bson:"uid"`

	// ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap.
	// This field is forbidden in Node.Spec, and required in Node.Status.
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,4,opt,name=resourceVersion" bson:"resourceVersion"`

	// KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure
	// This field is required in all cases.
	KubeletConfigKey string `json:"kubeletConfigKey" protobuf:"bytes,5,opt,name=kubeletConfigKey" bson:"kubeletConfigKey"`
}

type Taint struct {
	// Required. The taint key to be applied to a node.
	Key string `json:"key" protobuf:"bytes,1,opt,name=key" bson:"key"`
	// Required. The taint value corresponding to the taint key.
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value" bson:"value"`
	// Required. The effect of the taint on pods
	// that do not tolerate the taint.
	// Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	Effect TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect" bson:"effect"`
	// TimeAdded represents the time at which the taint was added.
	// It is only written for NoExecute taints.
	// +optional
	TimeAdded *Time `json:"timeAdded,omitempty" protobuf:"bytes,4,opt,name=timeAdded" bson:"timeAdded"`
}

type NodePhase string

// These are the valid phases of node.
const (
	// NodePending means the node has been created/added by the system, but not configured.
	NodePending NodePhase = "Pending"
	// NodeRunning means the node has been configured and has Kubernetes components running.
	NodeRunning NodePhase = "Running"
	// NodeTerminated means the node has been removed from the cluster.
	NodeTerminated NodePhase = "Terminated"
)

type NodeConditionType string

// These are valid conditions of node. Currently, we don't have enough information to decide
// node condition. In the future, we will add more. The proposed set of conditions are:
// NodeReachable, NodeLive, NodeReady, NodeSchedulable, NodeRunnable.
const (
	// NodeReady means kubelet is healthy and ready to accept pods.
	NodeReady NodeConditionType = "Ready"
	// NodeMemoryPressure means the kubelet is under pressure due to insufficient available memory.
	NodeMemoryPressure NodeConditionType = "MemoryPressure"
	// NodeDiskPressure means the kubelet is under pressure due to insufficient available disk.
	NodeDiskPressure NodeConditionType = "DiskPressure"
	// NodePIDPressure means the kubelet is under pressure due to insufficient available PID.
	NodePIDPressure NodeConditionType = "PIDPressure"
	// NodeNetworkUnavailable means that network for the node is not correctly configured.
	NodeNetworkUnavailable NodeConditionType = "NetworkUnavailable"
)

// NodeCondition contains condition information for a node.
type NodeCondition struct {
	// Type of node condition.
	Type NodeConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeConditionType" bson:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus" bson:"status"`
	// Last time we got an Update on a given condition.
	// +optional
	LastHeartbeatTime Time `json:"lastHeartbeatTime,omitempty" protobuf:"bytes,3,opt,name=lastHeartbeatTime" bson:"lastHeartbeatTime"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime" bson:"lastTransitionTime"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason" bson:"reason"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message" bson:"message"`
}

type NodeAddressType string

// These are valid address type of node.
const (
	NodeHostName    NodeAddressType = "Hostname"
	NodeExternalIP  NodeAddressType = "ExternalIP"
	NodeInternalIP  NodeAddressType = "InternalIP"
	NodeExternalDNS NodeAddressType = "ExternalDNS"
	NodeInternalDNS NodeAddressType = "InternalDNS"
)

// NodeAddress contains information for the node's address.
type NodeAddress struct {
	// Node address type, one of Hostname, ExternalIP or InternalIP.
	Type NodeAddressType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeAddressType" bson:"type"`
	// The node address.
	Address string `json:"address" protobuf:"bytes,2,opt,name=address" bson:"address"`
}

const (
	// Default namespace prefix.
	ResourceDefaultNamespacePrefix = "kubernetes.io/"
	// Name prefix for huge page resources (alpha).
	ResourceHugePagesPrefix = "hugepages-"
	// Name prefix for storage resource limits
	ResourceAttachableVolumesPrefix = "attachable-volumes-"
)

type DaemonEndpoint struct {
	/*
		The port tag was not properly in quotes in earlier releases, so it must be
		uppercased for backwards compat (since it was falling back to var name of
		'Port').
	*/

	// Port number of the given endpoint.
	Port int32 `json:"Port" protobuf:"varint,1,opt,name=Port" bson:"Port"`
}

// NodeDaemonEndpoints lists ports opened by daemons running on the Node.
type NodeDaemonEndpoints struct {
	// Endpoint on which Kubelet is listening.
	// +optional
	KubeletEndpoint DaemonEndpoint `json:"kubeletEndpoint,omitempty" protobuf:"bytes,1,opt,name=kubeletEndpoint" bson:"kubeletEndpoint"`
}

// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type NodeSystemInfo struct {
	// MachineID reported by the node. For unique machine identification
	// in the cluster this field is preferred. Learn more from man(5)
	// machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID string `json:"machineID" protobuf:"bytes,1,opt,name=machineID" bson:"machineID"`
	// SystemUUID reported by the node. For unique machine identification
	// MachineID is preferred. This field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	SystemUUID string `json:"systemUUID" protobuf:"bytes,2,opt,name=systemUUID" bson:"systemUUID"`
	// Boot ID reported by the node.
	BootID string `json:"bootID" protobuf:"bytes,3,opt,name=bootID" bson:"bootID"`
	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `json:"kernelVersion" protobuf:"bytes,4,opt,name=kernelVersion" bson:"kernelVersion"`
	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OSImage string `json:"osImage" protobuf:"bytes,5,opt,name=osImage" bson:"osImage"`
	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `json:"containerRuntimeVersion" protobuf:"bytes,6,opt,name=containerRuntimeVersion" bson:"containerRuntimeVersion"`
	// Kubelet Version reported by the node.
	KubeletVersion string `json:"kubeletVersion" protobuf:"bytes,7,opt,name=kubeletVersion" bson:"kubeletVersion"`
	// KubeProxy Version reported by the node.
	KubeProxyVersion string `json:"kubeProxyVersion" protobuf:"bytes,8,opt,name=kubeProxyVersion" bson:"kubeProxyVersion"`
	// The Operating System reported by the node
	OperatingSystem string `json:"operatingSystem" protobuf:"bytes,9,opt,name=operatingSystem" bson:"operatingSystem"`
	// The Architecture reported by the node
	Architecture string `json:"architecture" protobuf:"bytes,10,opt,name=architecture" bson:"architecture"`
}

// NodeConfigStatus describes the status of the config assigned by Node.Spec.ConfigSource.
type NodeConfigStatus struct {
	// Assigned reports the checkpointed config the node will try to use.
	// When Node.Spec.ConfigSource is updated, the node checkpoints the associated
	// config payload to local disk, along with a record indicating intended
	// config. The node refers to this record to choose its config checkpoint, and
	// reports this record in Assigned. Assigned only updates in the status after
	// the record has been checkpointed to disk. When the Kubelet is restarted,
	// it tries to make the Assigned config the Active config by loading and
	// validating the checkpointed payload identified by Assigned.
	// +optional
	Assigned *NodeConfigSource `json:"assigned,omitempty" protobuf:"bytes,1,opt,name=assigned" bson:"assigned"`
	// Active reports the checkpointed config the node is actively using.
	// Active will represent either the current version of the Assigned config,
	// or the current LastKnownGood config, depending on whether attempting to use the
	// Assigned config results in an error.
	// +optional
	Active *NodeConfigSource `json:"active,omitempty" protobuf:"bytes,2,opt,name=active" bson:"active"`
	// LastKnownGood reports the checkpointed config the node will fall back to
	// when it encounters an error attempting to use the Assigned config.
	// The Assigned config becomes the LastKnownGood config when the node determines
	// that the Assigned config is stable and correct.
	// This is currently implemented as a 10-minute soak period starting when the local
	// record of Assigned config is updated. If the Assigned config is Active at the end
	// of this period, it becomes the LastKnownGood. Note that if Spec.ConfigSource is
	// reset to nil (use local defaults), the LastKnownGood is also immediately reset to nil,
	// because the local default config is always assumed good.
	// You should not make assumptions about the node's method of determining config stability
	// and correctness, as this may change or become configurable in the future.
	// +optional
	LastKnownGood *NodeConfigSource `json:"lastKnownGood,omitempty" protobuf:"bytes,3,opt,name=lastKnownGood" bson:"lastKnownGood"`
	// Error describes any problems reconciling the Spec.ConfigSource to the Active config.
	// Errors may occur, for example, attempting to checkpoint Spec.ConfigSource to the local Assigned
	// record, attempting to checkpoint the payload associated with Spec.ConfigSource, attempting
	// to load or validate the Assigned config, etc.
	// Errors may occur at different points while syncing config. Earlier errors (e.g. download or
	// checkpointing errors) will not result in a rollback to LastKnownGood, and may resolve across
	// Kubelet retries. Later errors (e.g. loading or validating a checkpointed config) will result in
	// a rollback to LastKnownGood. In the latter case, it is usually possible to resolve the error
	// by fixing the config assigned in Spec.ConfigSource.
	// You can find additional information for debugging by searching the error message in the Kubelet log.
	// Error is a human-readable description of the error state; machines can check whether or not Error
	// is empty, but should not rely on the stability of the Error text across Kubelet versions.
	// +optional
	Error string `json:"error,omitempty" protobuf:"bytes,4,opt,name=error" bson:"error"`
}

type ContainerImage struct {
	// Names by which this image is known.
	// e.g. ["k8s.gcr.io/hyperkube:v1.0.7", "dockerhub.io/google_containers/hyperkube:v1.0.7"]
	Names []string `json:"names" protobuf:"bytes,1,rep,name=names" bson:"names"`
	// The size of the image in bytes.
	// +optional
	SizeBytes int64 `json:"sizeBytes,omitempty" protobuf:"varint,2,opt,name=sizeBytes" bson:"sizeBytes"`
}

type UniqueVolumeName string

// AttachedVolume describes a volume attached to a node
type AttachedVolume struct {
	// Name of the attached volume
	Name UniqueVolumeName `json:"name" protobuf:"bytes,1,rep,name=name" bson:"name"`

	// DevicePath represents the device path where the volume should be available
	DevicePath string `json:"devicePath" protobuf:"bytes,2,rep,name=devicePath" bson:"devicePath"`
}
type NodeStatus struct {
	// Capacity represents the total resources of a node.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// +optional
	Capacity ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName" bson:"capacity"`
	// Allocatable represents the resources of a node that are available for scheduling.
	// Defaults to Capacity.
	// +optional
	Allocatable ResourceList `json:"allocatable,omitempty" protobuf:"bytes,2,rep,name=allocatable,casttype=ResourceList,castkey=ResourceName" bson:"allocatable"`
	// NodePhase is the recently observed lifecycle phase of the node.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#phase
	// The field is never populated, and now is deprecated.
	// +optional
	Phase NodePhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=NodePhase" bson:"phase"`
	// Conditions is an array of current observed node conditions.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []NodeCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions" bson:"conditions"`
	// List of addresses reachable to the node.
	// Queried from cloud provider, if available.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	// Note: This field is declared as mergeable, but the merge key is not sufficiently
	// unique, which can cause data corruption when it is merged. Callers should instead
	// use a full-replacement patch. See http://pr.k8s.io/79391 for an example.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Addresses []NodeAddress `json:"addresses,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,5,rep,name=addresses" bson:"addresses"`
	// Endpoints of daemons running on the Node.
	// +optional
	DaemonEndpoints NodeDaemonEndpoints `json:"daemonEndpoints,omitempty" protobuf:"bytes,6,opt,name=daemonEndpoints" bson:"daemonEndpoints"`
	// Set of ids/uuids to uniquely identify the node.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#info
	// +optional
	NodeInfo NodeSystemInfo `json:"nodeInfo,omitempty" protobuf:"bytes,7,opt,name=nodeInfo" bson:"nodeInfo"`
	// List of container images on this node
	// +optional
	Images []ContainerImage `json:"images,omitempty" protobuf:"bytes,8,rep,name=images" bson:"images"`
	// List of attachable volumes in use (mounted) by the node.
	// +optional
	VolumesInUse []UniqueVolumeName `json:"volumesInUse,omitempty" protobuf:"bytes,9,rep,name=volumesInUse" bson:"volumesInUse"`
	// List of volumes that are attached to the node.
	// +optional
	VolumesAttached []AttachedVolume `json:"volumesAttached,omitempty" protobuf:"bytes,10,rep,name=volumesAttached" bson:"volumesAttached"`
	// Status of the config assigned to the node via the dynamic Kubelet config feature.
	// +optional
	Config *NodeConfigStatus `json:"config,omitempty" protobuf:"bytes,11,opt,name=config" bson:"config"`
}

/// /// ///
// +kubebuilder:validation:Enum=RSA;ECDSA;Ed25519
type PrivateKeyAlgorithm string

const (
	// Denotes the RSA private key type.
	RSAKeyAlgorithm PrivateKeyAlgorithm = "RSA"

	// Denotes the ECDSA private key type.
	ECDSAKeyAlgorithm PrivateKeyAlgorithm = "ECDSA"

	// Denotes the Ed25519 private key type.
	Ed25519KeyAlgorithm PrivateKeyAlgorithm = "Ed25519"
)

// +kubebuilder:validation:Enum=PKCS1;PKCS8
type PrivateKeyEncoding string

const (
	// PKCS1 key encoding will produce PEM files that include the type of
	// private key as part of the PEM header, e.g. `BEGIN RSA PRIVATE KEY`.
	// If the keyAlgorithm is set to 'ECDSA', this will produce private keys
	// that use the `BEGIN EC PRIVATE KEY` header.
	PKCS1 PrivateKeyEncoding = "PKCS1"

	// PKCS8 key encoding will produce PEM files with the `BEGIN PRIVATE KEY`
	// header. It encodes the keyAlgorithm of the private key as part of the
	// DER encoded PEM block.
	PKCS8 PrivateKeyEncoding = "PKCS8"
)

type CertificateObjectReference struct {
	// Name of the resource being referred to.
	Name string `json:"name" bson:"name"`
	// Kind of the resource being referred to.
	// +optional
	Kind string `json:"kind,omitempty" bson:"kind"`
	// Group of the resource being referred to.
	// +optional
	Group string `json:"group,omitempty" bson:"group"`
}

// CertificateSpec defines the desired state of Certificate.
// A valid Certificate requires at least one of a CommonName, DNSName, or
// URISAN to be valid.
type CertificateSpec struct {
	// Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).
	// +optional
	Subject *X509Subject `json:"subject,omitempty" bson:"subject"`

	// CommonName is a common name to be used on the Certificate.
	// The CommonName should have a length of 64 characters or fewer to avoid
	// generating invalid CSRs.
	// This value is ignored by TLS clients when any subject alt name is set.
	// This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4
	// +optional
	CommonName string `json:"commonName,omitempty" bson:"commonName"`

	// The requested 'duration' (i.e. lifetime) of the Certificate. This option
	// may be ignored/overridden by some issuer types. If unset this defaults to
	// 90 days. Certificate will be renewed either 2/3 through its duration or
	// `renewBefore` period before its expiry, whichever is later. Minimum
	// accepted duration is 1 hour. Value must be in units accepted by Go
	// time.ParseDuration https://golang.org/pkg/time/#ParseDuration
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty" bson:"duration"`

	// How long before the currently issued certificate's expiry
	// cert-manager should renew the certificate. The default is 2/3 of the
	// issued certificate's duration. Minimum accepted value is 5 minutes.
	// Value must be in units accepted by Go time.ParseDuration
	// https://golang.org/pkg/time/#ParseDuration
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty" bson:"renewBefore"`

	// DNSNames is a list of DNS subjectAltNames to be set on the Certificate.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty" bson:"dnsNames"`

	// IPAddresses is a list of IP address subjectAltNames to be set on the Certificate.
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty" bson:"ipAddresses"`

	// URIs is a list of URI subjectAltNames to be set on the Certificate.
	// +optional
	URIs []string `json:"uris,omitempty" bson:"uris"`

	// EmailAddresses is a list of email subjectAltNames to be set on the Certificate.
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty" bson:"emailAddresses"`

	// SecretName is the name of the secret resource that will be automatically
	// created and managed by this Certificate resource.
	// It will be populated with a private key and certificate, signed by the
	// denoted issuer.
	SecretName string `json:"secretName" bson:"secretName"`

	// SecretTemplate defines annotations and labels to be propagated
	// to the Kubernetes Secret when it is created or updated. Once created,
	// labels and annotations are not yet removed from the Secret when they are
	// removed from the template. See https://github.com/jetstack/cert-manager/issues/4292
	// +optional
	SecretTemplate *CertificateSecretTemplate `json:"secretTemplate,omitempty" bson:"secretTemplate"`

	// Keystores configures additional keystore output formats stored in the
	// `secretName` Secret resource.
	// +optional
	Keystores *CertificateKeystores `json:"keystores,omitempty" bson:"keystores"`

	// IssuerRef is a reference to the issuer for this certificate.
	// If the `kind` field is not set, or set to `Issuer`, an Issuer resource
	// with the given name in the same namespace as the Certificate will be used.
	// If the `kind` field is set to `ClusterIssuer`, a ClusterIssuer with the
	// provided name will be used.
	// The `name` field in this stanza is required at all times.
	IssuerRef CertificateObjectReference `json:"issuerRef"  bson:"issuerRef"`

	// IsCA will mark this Certificate as valid for certificate signing.
	// This will automatically add the `cert sign` usage to the list of `usages`.
	// +optional
	IsCA bool `json:"isCA,omitempty"  bson:"isCA"`

	// Usages is the set of x509 usages that are requested for the certificate.
	// Defaults to `digital signature` and `key encipherment` if not specified.
	// +optional
	Usages []KeyUsage `json:"usages,omitempty"  bson:"usages"`

	// Options to control private keys used for the Certificate.
	// +optional
	PrivateKey *CertificatePrivateKey `json:"privateKey,omitempty"  bson:"privateKey"`

	// EncodeUsagesInRequest controls whether key usages should be present
	// in the CertificateRequest
	// +optional
	EncodeUsagesInRequest *bool `json:"encodeUsagesInRequest,omitempty"  bson:"encodeUsagesInRequest"`

	// revisionHistoryLimit is the maximum number of CertificateRequest revisions
	// that are maintained in the Certificate's history. Each revision represents
	// a single `CertificateRequest` created by this Certificate, either when it
	// was created, renewed, or Spec was changed. Revisions will be removed by
	// oldest first if the number of revisions exceeds this number. If set,
	// revisionHistoryLimit must be a value of `1` or greater. If unset (`nil`),
	// revisions will not be garbage collected. Default value is `nil`.
	// +kubebuilder:validation:ExclusiveMaximum=false
	// +optional
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"  bson:"revisionHistoryLimit"` // Validated by the validating webhook.
}

// CertificatePrivateKey contains configuration options for private keys
// used by the Certificate controller.
// This allows control of how private keys are rotated.
type CertificatePrivateKey struct {
	// RotationPolicy controls how private keys should be regenerated when a
	// re-issuance is being processed.
	// If set to Never, a private key will only be generated if one does not
	// already exist in the target `spec.secretName`. If one does exists but it
	// does not have the correct algorithm or size, a warning will be raised
	// to await user intervention.
	// If set to Always, a private key matching the specified requirements
	// will be generated whenever a re-issuance occurs.
	// Default is 'Never' for backward compatibility.
	// +optional
	RotationPolicy PrivateKeyRotationPolicy `json:"rotationPolicy,omitempty" bson:"rotationPolicy"`

	// The private key cryptography standards (PKCS) encoding for this
	// certificate's private key to be encoded in.
	// If provided, allowed values are `PKCS1` and `PKCS8` standing for PKCS#1
	// and PKCS#8, respectively.
	// Defaults to `PKCS1` if not specified.
	// +optional
	Encoding PrivateKeyEncoding `json:"encoding,omitempty" bson:"encoding"`

	// Algorithm is the private key algorithm of the corresponding private key
	// for this certificate. If provided, allowed values are either `RSA`,`Ed25519` or `ECDSA`
	// If `algorithm` is specified and `size` is not provided,
	// key size of 256 will be used for `ECDSA` key algorithm and
	// key size of 2048 will be used for `RSA` key algorithm.
	// key size is ignored when using the `Ed25519` key algorithm.
	// +optional
	Algorithm PrivateKeyAlgorithm `json:"algorithm,omitempty" bson:"algorithm"`

	// Size is the key bit size of the corresponding private key for this certificate.
	// If `algorithm` is set to `RSA`, valid values are `2048`, `4096` or `8192`,
	// and will default to `2048` if not specified.
	// If `algorithm` is set to `ECDSA`, valid values are `256`, `384` or `521`,
	// and will default to `256` if not specified.
	// If `algorithm` is set to `Ed25519`, Size is ignored.
	// No other values are allowed.
	// +optional
	Size int `json:"size,omitempty" bson:"size"` // Validated by webhook. Be mindful of adding OpenAPI validation- see https://github.com/jetstack/cert-manager/issues/3644
}

// Denotes how private keys should be generated or sourced when a Certificate
// is being issued.
type PrivateKeyRotationPolicy string

var (
	// RotationPolicyNever means a private key will only be generated if one
	// does not already exist in the target `spec.secretName`.
	// If one does exists but it does not have the correct algorithm or size,
	// a warning will be raised to await user intervention.
	RotationPolicyNever PrivateKeyRotationPolicy = "Never"

	// RotationPolicyAlways means a private key matching the specified
	// requirements will be generated whenever a re-issuance occurs.
	RotationPolicyAlways PrivateKeyRotationPolicy = "Always"
)

// X509Subject Full X509 name specification
type X509Subject struct {
	// Organizations to be used on the Certificate.
	// +optional
	Organizations []string `json:"organizations,omitempty" bson:"organizations"`
	// Countries to be used on the Certificate.
	// +optional
	Countries []string `json:"countries,omitempty" bson:"countries"`
	// Organizational Units to be used on the Certificate.
	// +optional
	OrganizationalUnits []string `json:"organizationalUnits,omitempty" bson:"organizationalUnits"`
	// Cities to be used on the Certificate.
	// +optional
	Localities []string `json:"localities,omitempty" bson:"localities"`
	// State/Provinces to be used on the Certificate.
	// +optional
	Provinces []string `json:"provinces,omitempty" bson:"provinces"`
	// Street addresses to be used on the Certificate.
	// +optional
	StreetAddresses []string `json:"streetAddresses,omitempty" bson:"streetAddresses"`
	// Postal codes to be used on the Certificate.
	// +optional
	PostalCodes []string `json:"postalCodes,omitempty" bson:"postalCodes"`
	// Serial number to be used on the Certificate.
	// +optional
	SerialNumber string `json:"serialNumber,omitempty" bson:"serialNumber"`
}

// CertificateKeystores configures additional keystore output formats to be
// created in the Certificate's output Secret.
type CertificateKeystores struct {
	// JKS configures options for storing a JKS keystore in the
	// `spec.secretName` Secret resource.
	// +optional
	JKS *JKSKeystore `json:"jks,omitempty" bson:"jks"`

	// PKCS12 configures options for storing a PKCS12 keystore in the
	// `spec.secretName` Secret resource.
	// +optional
	PKCS12 *PKCS12Keystore `json:"pkcs12,omitempty" bson:"pkcs12"`
}

// JKS configures options for storing a JKS keystore in the `spec.secretName`
// Secret resource.
type JKSKeystore struct {
	// Create enables JKS keystore creation for the Certificate.
	// If true, a file named `keystore.jks` will be created in the target
	// Secret resource, encrypted using the password stored in
	// `passwordSecretRef`.
	// The keystore file will only be updated upon re-issuance.
	// A file named `truststore.jks` will also be created in the target
	// Secret resource, encrypted using the password stored in
	// `passwordSecretRef` containing the issuing Certificate Authority
	Create bool `json:"create" bson:"create"`

	// PasswordSecretRef is a reference to a key in a Secret resource
	// containing the password used to encrypt the JKS keystore.
	PasswordSecretRef CertificateSecretKeySelector `json:"passwordSecretRef" bson:"passwordSecretRef"`
}

type CertificateSecretKeySelector struct {
	// The name of the Secret resource being referred to.
	LocalObjectReference `json:",inline" bson:"inline"`

	// The key of the entry in the Secret resource's `data` field to be used.
	// Some instances of this field may be defaulted, in others it may be
	// required.
	// +optional
	Key string `json:"key,omitempty" bson:"key"`
}

// PKCS12 configures options for storing a PKCS12 keystore in the
// `spec.secretName` Secret resource.
type PKCS12Keystore struct {
	// Create enables PKCS12 keystore creation for the Certificate.
	// If true, a file named `keystore.p12` will be created in the target
	// Secret resource, encrypted using the password stored in
	// `passwordSecretRef`.
	// The keystore file will only be updated upon re-issuance.
	// A file named `truststore.p12` will also be created in the target
	// Secret resource, encrypted using the password stored in
	// `passwordSecretRef` containing the issuing Certificate Authority
	Create bool `json:"create" bson:"create"`

	// PasswordSecretRef is a reference to a key in a Secret resource
	// containing the password used to encrypt the PKCS12 keystore.
	PasswordSecretRef CertificateSecretKeySelector `json:"passwordSecretRef" bson:"passwordSecretRef"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	// List of status conditions to indicate the status of certificates.
	// Known condition types are `Ready` and `Issuing`.
	// +optional
	Conditions []CertificateCondition `json:"conditions,omitempty"  bson:"conditions"`

	// LastFailureTime is the time as recorded by the Certificate controller
	// of the most recent failure to complete a CertificateRequest for this
	// Certificate resource.
	// If set, cert-manager will not re-request another Certificate until
	// 1 hour has elapsed from this time.
	// +optional
	LastFailureTime *metav1.Time `json:"lastFailureTime,omitempty" bson:"lastFailureTime"`

	// The time after which the certificate stored in the secret named
	// by this resource in spec.secretName is valid.
	// +optional
	NotBefore *metav1.Time `json:"notBefore,omitempty" bson:"notBefore"`

	// The expiration time of the certificate stored in the secret named
	// by this resource in `spec.secretName`.
	// +optional
	NotAfter *metav1.Time `json:"notAfter,omitempty" bson:"notAfter"`

	// RenewalTime is the time at which the certificate will be next
	// renewed.
	// If not set, no upcoming renewal is scheduled.
	// +optional
	RenewalTime *metav1.Time `json:"renewalTime,omitempty" bson:"renewalTime"`

	// The current 'revision' of the certificate as issued.
	//
	// When a CertificateRequest resource is created, it will have the
	// `cert-manager.io/certificate-revision` set to one greater than the
	// current value of this field.
	//
	// Upon issuance, this field will be set to the value of the annotation
	// on the CertificateRequest resource used to issue the certificate.
	//
	// Persisting the value on the CertificateRequest resource allows the
	// certificates controller to know whether a request is part of an old
	// issuance or if it is part of the ongoing revision's issuance by
	// checking if the revision value in the annotation is greater than this
	// field.
	// +optional
	Revision *int `json:"revision,omitempty" bson:"revision"`

	// The name of the Secret resource containing the private key to be used
	// for the next certificate iteration.
	// The keymanager controller will automatically set this field if the
	// `Issuing` condition is set to `True`.
	// It will automatically unset this field when the Issuing condition is
	// not set or False.
	// +optional
	NextPrivateKeySecretName *string `json:"nextPrivateKeySecretName,omitempty" bson:"nextPrivateKeySecretName"`
}

// CertificateCondition contains condition information for an Certificate.
type CertificateCondition struct {
	// Type of the condition, known values are (`Ready`, `Issuing`).
	Type CertificateConditionType `json:"type"  bson:"type"`

	// Status of the condition, one of (`True`, `False`, `Unknown`).
	Status cmmeta.ConditionStatus `json:"status"  bson:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"  bson:"lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	// +optional
	Reason string `json:"reason,omitempty"  bson:"reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	// +optional
	Message string `json:"message,omitempty" bson:"message"`

	// If set, this represents the .metadata.generation that the condition was
	// set based upon.
	// For instance, if .metadata.generation is currently 12, but the
	// .status.condition[x].observedGeneration is 9, the condition is out of date
	// with respect to the current state of the Certificate.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" bson:"observedGeneration"`
}

// CertificateConditionType represents an Certificate condition value.
type CertificateConditionType string

const (
	// CertificateConditionReady indicates that a certificate is ready for use.
	// This is defined as:
	// - The target secret exists
	// - The target secret contains a certificate that has not expired
	// - The target secret contains a private key valid for the certificate
	// - The commonName and dnsNames attributes match those specified on the Certificate
	CertificateConditionReady CertificateConditionType = "Ready"

	// A condition added to Certificate resources when an issuance is required.
	// This condition will be automatically added and set to true if:
	//   * No keypair data exists in the target Secret
	//   * The data stored in the Secret cannot be decoded
	//   * The private key and certificate do not have matching public keys
	//   * If a CertificateRequest for the current revision exists and the
	//     certificate data stored in the Secret does not match the
	//    `status.certificate` on the CertificateRequest.
	//   * If no CertificateRequest resource exists for the current revision,
	//     the options on the Certificate resource are compared against the
	//     x509 data in the Secret, similar to what's done in earlier versions.
	//     If there is a mismatch, an issuance is triggered.
	// This condition may also be added by external API consumers to trigger
	// a re-issuance manually for any other reason.
	//
	// It will be removed by the 'issuing' controller upon completing issuance.
	CertificateConditionIssuing CertificateConditionType = "Issuing"
)

// CertificateSecretTemplate defines the default labels and annotations
// to be copied to the Kubernetes Secret resource named in `CertificateSpec.secretName`.
type CertificateSecretTemplate struct {
	// Annotations is a key value map to be copied to the target Kubernetes Secret.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" bson:"annotations"`

	// Labels is a key value map to be copied to the target Kubernetes Secret.
	// +optional
	Labels map[string]string `json:"labels,omitempty" bson:"labels"`
}

type KeyUsage string

const (
	UsageSigning           KeyUsage = "signing"
	UsageDigitalSignature  KeyUsage = "digital signature"
	UsageContentCommitment KeyUsage = "content commitment"
	UsageKeyEncipherment   KeyUsage = "key encipherment"
	UsageKeyAgreement      KeyUsage = "key agreement"
	UsageDataEncipherment  KeyUsage = "data encipherment"
	UsageCertSign          KeyUsage = "cert sign"
	UsageCRLSign           KeyUsage = "crl sign"
	UsageEncipherOnly      KeyUsage = "encipher only"
	UsageDecipherOnly      KeyUsage = "decipher only"
	UsageAny               KeyUsage = "any"
	UsageServerAuth        KeyUsage = "server auth"
	UsageClientAuth        KeyUsage = "client auth"
	UsageCodeSigning       KeyUsage = "code signing"
	UsageEmailProtection   KeyUsage = "email protection"
	UsageSMIME             KeyUsage = "s/mime"
	UsageIPsecEndSystem    KeyUsage = "ipsec end system"
	UsageIPsecTunnel       KeyUsage = "ipsec tunnel"
	UsageIPsecUser         KeyUsage = "ipsec user"
	UsageTimestamping      KeyUsage = "timestamping"
	UsageOCSPSigning       KeyUsage = "ocsp signing"
	UsageMicrosoftSGC      KeyUsage = "microsoft sgc"
	UsageNetscapeSGC       KeyUsage = "netscape sgc"
)

// DefaultKeyUsages contains the default list of key usages
func DefaultKeyUsages() []KeyUsage {
	// The serverAuth EKU is required as of Mac OS Catalina: https://support.apple.com/en-us/HT210176
	// Without this usage, certificates will _always_ flag a warning in newer Mac OS browsers.
	// We don't explicitly add it here as it leads to strange behaviour when a user sets isCA: true
	// (in which case, 'serverAuth' on the CA can break a lot of clients).
	// CAs can (and often do) opt to automatically add usages.
	return []KeyUsage{UsageDigitalSignature, UsageKeyEncipherment}
}
