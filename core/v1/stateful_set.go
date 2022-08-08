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
	StatefulSetCollection = "statefulSetCollection"
)

type K8sStatefulSet struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       StatefulSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     StatefulSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type StatefulSet struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sStatefulSet `bson:"obj" json:"obj"`
	AgentName          string         `bson:"agent_name" json:"agent_name"`
}

func (obj StatefulSet) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete statefulSet [ERROR]", err)
	}
	return err
}

func NewStatefulSet() KubeObject {
	return &StatefulSet{}
}

func (obj StatefulSet) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespaceAndCompanyId().Name == "" {
		coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
		go AgentIndex{}.Build(obj.Obj.ObjectMeta.Labels["company"], obj.AgentName).Save()
	} else {
		err := obj.Update(StatefulSet{Obj: obj.findByNameAndNamespaceAndCompanyId(), AgentName: obj.AgentName}, obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj StatefulSet) findById() K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(StatefulSet)
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj StatefulSet) findByNameAndNamespaceAndCompanyId() K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.metadata.labels.company": obj.Obj.ObjectMeta.Labels["company"]},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(StatefulSet)
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj StatefulSet) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj StatefulSet) Update(oldObj interface{}, agent string) error {
	var oldObject StatefulSet
	body, _ := json.Marshal(oldObj)
	errorOfUnmarshal := json.Unmarshal(body, &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
	}
	if obj.AgentName == "" {
		obj.AgentName = agent
	}
	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": obj.Obj.Namespace},
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
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
	go AgentIndex{}.Build(obj.Obj.ObjectMeta.Labels["company"], obj.AgentName).Save()
	return nil
}

func (obj StatefulSet) saveAll(objs []StatefulSet) error {
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
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

func (object StatefulSet) findAll() []K8sStatefulSet {
	query := bson.M{}
	objects := []StatefulSet{}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(StatefulSet)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sStatefulSet{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object StatefulSet) findByNamespace() []K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []StatefulSet{}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(StatefulSet)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sStatefulSet{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object StatefulSet) findBykubeAgentNameAndNamespace() []K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []StatefulSet{}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(StatefulSet)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sStatefulSet{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object StatefulSet) findBykubeClusterId() []K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []StatefulSet{}
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(StatefulSet)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sStatefulSet{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object StatefulSet) findByName() K8sStatefulSet {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(StatefulSet)
	coll := db.GetDmManager().Db.Collection(StatefulSetCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
