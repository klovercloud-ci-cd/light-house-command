package v1

import (
	"context"
	"encoding/json"
	"github.com/go-bongo/bongo"
	"github.com/klovercloud-ci-cd/light-house-command/core/v1/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	NetworkPolicyCollection = "networkPolicyCollection"
)

type K8sNetworkPolicy struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"Obj"`
	Spec       NetworkPolicySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
}

type NetworkPolicy struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNetworkPolicy `bson:"obj" json:"obj"`
	AgentName          string           `json:"agent_name" bson:"agent_name"`
}

func (obj NetworkPolicy) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete networkPolicy [ERROR]", err)
	}
	return err
}

func NewNetworkPolicy() KubeObject {
	return &NetworkPolicy{}
}

func (obj NetworkPolicy) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(NetworkPolicy{Obj:obj.findByNameAndNamespace(), AgentName: obj.AgentName},obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj NetworkPolicy) findById() K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(NetworkPolicy)
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj NetworkPolicy) findByNameAndNamespace() K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(NetworkPolicy)
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj NetworkPolicy) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj NetworkPolicy) Update(oldObj interface{},agent string) error {
	var oldObject NetworkPolicy
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
			{"obj.metadata.namespace": obj.Obj.Namespace},
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
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj NetworkPolicy) saveAll(objs []NetworkPolicy) error {
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
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

func (object NetworkPolicy) findAll() []K8sNetworkPolicy {
	query := bson.M{}
	objects := []NetworkPolicy{}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(NetworkPolicy)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNetworkPolicy{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object NetworkPolicy) findByNamespace() []K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []NetworkPolicy{}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(NetworkPolicy)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNetworkPolicy{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object NetworkPolicy) findBykubeAgentNameAndNamespace() []K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []NetworkPolicy{}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(NetworkPolicy)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNetworkPolicy{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object NetworkPolicy) findBykubeAgentName() []K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []NetworkPolicy{}
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(NetworkPolicy)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNetworkPolicy{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object NetworkPolicy) findByName() K8sNetworkPolicy {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(NetworkPolicy)
	coll := db.GetDmManager().Db.Collection(NetworkPolicyCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
