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
	PodCollection = "podCollection"
)

type K8sPod struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type Pod struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sPod `bson:"obj" json:"obj"`
	AgentName          string `bson:"agent_name" json:"agent_name"`
}

func (obj Pod) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete pod [ERROR]", err)
	}
	return err
}

func NewPod() KubeObject {
	return &Pod{}
}

func (obj Pod) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(PodCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(Pod{Obj: obj.findByNameAndNamespace(), AgentName: obj.AgentName}, obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj Pod) findById() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Pod) findByNameAndNamespace() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Pod) findByLabel() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.labels": obj.Obj.Labels},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Pod) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": agent},
		},
	}
	log.Println("deleting pod:", obj.Obj.Name+"!")
	coll := db.GetDmManager().Db.Collection(PodCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj Pod) Update(oldObj interface{}, agent string) error {
	log.Println("Pod:", obj.Obj.Name, ", Status: ", obj.Obj.Status.Phase)
	var oldObject Pod
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
			{"obj.metadata.name": oldObject.Obj.Name},
			{"obj.metadata.namespace": oldObject.Obj.Namespace},
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
	coll := db.GetDmManager().Db.Collection(PodCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj Pod) saveAll(objs []Pod) error {
	coll := db.GetDmManager().Db.Collection(PodCollection)
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

func (object Pod) findAll() []K8sPod {
	query := bson.M{}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findByNamespace() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findBykubeAgentNameAndNamespace() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findBykubeAgentName() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findByName() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
