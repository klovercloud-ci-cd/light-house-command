package v1

import (
	"context"
	"encoding/json"
	"github.com/go-bongo/bongo"
	"github.com/klovercloud/lighthouse-command/core/v1/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type FinalizerName string
type NamespacePhase string

const (
	NamespaceCollection                 = "namespaceCollection"
	FinalizerKubernetes  FinalizerName  = "kubernetes"
	NamespaceActive      NamespacePhase = "Active"
	NamespaceTerminating NamespacePhase = "Terminating"
)

type K8sNamespace struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       NamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     NamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type Namespace struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNamespace `bson:"obj" json:"obj"`
	AgentName          string       `bson:"agent_name" json:"agent_name"`
}

func (obj Namespace) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete namespace [ERROR]", err)
	}
	return err
}

func NewNamespace() KubeObject {
	return &Namespace{}
}

func (obj Namespace) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNamespaceAndAgentName().Name == "" {
		coll := db.GetDmManager().Db.Collection(NamespaceCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(Namespace{Obj:obj.findByName(), AgentName: obj.AgentName},obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (object Namespace) findByName() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (object Namespace) findByNamespaceAndAgentName() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
func (obj Namespace) findById() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Namespace) findBykubeAgentName() []K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": obj.AgentName},
		},
	}
	namespaces := []Namespace{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Namespace)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		namespaces = append(namespaces, *elemValue)
	}
	k8sObjects := []K8sNamespace{}
	for _, each := range namespaces {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Namespace) findAll() []K8sNamespace {
	query := bson.M{}
	namespaces := []Namespace{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Namespace)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		namespaces = append(namespaces, *elemValue)
	}
	k8sObjects := []K8sNamespace{}
	for _, each := range namespaces {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Namespace) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj Namespace) Update(oldObj interface{},agent string) error {
	var oldObject Namespace
	body, _ := json.Marshal(oldObj)
	errorOfUnmarshal := json.Unmarshal(body, &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
	}
	if obj.AgentName == ""{
		obj.AgentName=agent
	}
	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	update := bson.M{
		"$set": obj,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj Namespace) saveAll(objs []Namespace) error {
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
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
