package v1

import (
	"context"
	"encoding/json"
	"github.com/go-bongo/bongo"
	"github.com/klovercloud/lighthouse-command/core/v1/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

const (
	EventCollection = "eventCollection"
)

type K8sEvent struct {
	TypeMeta `json:",inline" bson:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	ObjectMeta `json:"objectMeta" protobuf:"bytes,1,opt,name=metadata" bson:"objectMeta"`

	// The object that this event is about.
	InvolvedObject ObjectReference `json:"involvedObject" protobuf:"bytes,2,opt,name=involvedObject" bson:"involvedObject"`

	// This should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status.
	// TODO: provide exact specification for format.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason" bson:"reason"`

	// A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message" bson:"message"`

	// The component reporting this event. Should be a short machine understandable string.
	// +optional
	Source EventSource `json:"source,omitempty" protobuf:"bytes,5,opt,name=source" bson:"source"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	// +optional
	FirstTimestamp metav1.Time `json:"firstTimestamp,omitempty" protobuf:"bytes,6,opt,name=firstTimestamp" bson:"firstTimestamp"`

	// The time at which the most recent occurrence of this event was recorded.
	// +optional
	LastTimestamp metav1.Time `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp" bson:"lastTimestamp"`

	// The number of times this event has occurred.
	// +optional
	Count int32 `json:"count,omitempty" protobuf:"varint,8,opt,name=count" bson:"count"`

	// Type of this event (Normal, Warning), new types could be added in the future
	// +optional
	Type string `json:"type,omitempty" protobuf:"bytes,9,opt,name=type" bson:"type"`

	// Time when this Event was first observed.
	// +optional
	EventTime metav1.MicroTime `json:"eventTime,omitempty" protobuf:"bytes,10,opt,name=eventTime" bson:"eventTime"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	// +optional
	Series *EventSeries `json:"series,omitempty" protobuf:"bytes,11,opt,name=series" bson:"series"`

	// What action was taken/failed regarding to the Regarding object.
	// +optional
	Action string `json:"action,omitempty" protobuf:"bytes,12,opt,name=action" bson:"action"`

	// Optional secondary object for more complex actions.
	// +optional
	Related *ObjectReference `json:"related,omitempty" protobuf:"bytes,13,opt,name=related" bson:"related"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	// +optional
	ReportingController string `json:"reportingComponent" protobuf:"bytes,14,opt,name=reportingComponent" bson:"reportingController"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	// +optional
	ReportingInstance string `json:"reportingInstance" protobuf:"bytes,15,opt,name=reportingInstance" bson:"reportingInstance"`
}

type Event struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sEvent `bson:"obj" json:"obj"`
	AgentName          string   `bson:"agent_name" json:"agent_name"`
}

func (e Event) Save(extra map[string]string) error {
	e.AgentName = extra["agent_name"]
	if e.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(EventCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, e)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := e.Update(Event{Obj:e.findByNameAndNamespace(), AgentName: e.AgentName},e.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Event) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": e.Obj.Name},
			{"obj.metadata.namespace": e.Obj.Namespace},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (e Event) Update(oldObj interface{},agent string) error {
	var oldObject Ingress
	body, _ := json.Marshal(oldObj)
	errorOfUnmarshal := json.Unmarshal(body, &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
	}
	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": oldObject.Obj.Name},
			{"obj.metadata.namespace": oldObject.Obj.Namespace},
			{"agent_name":agent},
		},
	}
	update := bson.M{
		"$set": e,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (e Event) findByNameAndNamespace() K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": e.Obj.Name},
			{"obj.metadata.namespace": e.Obj.Namespace},
			{"agent_name": e.AgentName},
		},
	}
	temp := new(Event)
	coll := db.GetDmManager().Db.Collection(EventCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (e Event) findById() K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": e.Obj.UID},
			{"agent_name": e.AgentName},
		},
	}
	temp := new(Event)
	coll := db.GetDmManager().Db.Collection(EventCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (e Event) saveAll(objs []Event) error {
	coll := db.GetDmManager().Db.Collection(EventCollection)
	if len(objs) > 0 {
		var data []interface{}
		data = append(data, objs)
		_, err := coll.InsertMany(db.GetDmManager().Ctx, data)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (e Event) findAll() []K8sEvent {
	query := bson.M{}
	objects := []Event{}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Event)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sEvent{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (e Event) findByNamespace() []K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": e.Obj.Namespace},
		},
	}
	objects := []Event{}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Event)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sEvent{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (e Event) findBykubeAgentNameAndNamespace() []K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": e.Obj.Namespace},
			{"agent_name": e.AgentName},
		},
	}
	objects := []Event{}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Event)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sEvent{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (e Event) findBykubeAgentName() []K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": e.AgentName},
		},
	}
	objects := []Event{}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Event)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sEvent{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (e Event) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(EventCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete ingress [ERROR]", err)
	}
	return err
}

func (e Event) findByName() K8sEvent {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": e.Obj.Name},
			{"obj.metadata.namespace": e.Obj.Namespace},
			{"agent_name": e.AgentName},
		},
	}
	temp := new(Event)
	coll := db.GetDmManager().Db.Collection(EventCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func NewEvent() KubeObject {
	return &Event{}
}
