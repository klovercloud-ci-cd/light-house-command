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
	ServiceAccountCollection = "serviceAccountCollection"
)

type K8sServiceAccount struct {
	TypeMeta                     `json:",inline" bson:",inline"`
	ObjectMeta                   `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Secrets                      []ObjectReference      `json:"secrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=secrets" bson:"secrets"`
	ImagePullSecrets             []LocalObjectReference `json:"imagePullSecrets,omitempty" protobuf:"bytes,3,rep,name=imagePullSecrets" bson:"imagePullSecrets"`
	AutomountServiceAccountToken *bool                  `json:"automountServiceAccountToken,omitempty" protobuf:"varint,4,opt,name=automountServiceAccountToken" bson:"automountServiceAccountToken"`
}

type ServiceAccount struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sServiceAccount `bson:"obj" json:"obj"`
	AgentName          string            `bson:"agent_name" json:"agent_name"`
}

func (obj ServiceAccount) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete sa [ERROR]", err)
	}
	return err
}

func NewServiceAccount() KubeObject {
	return &ServiceAccount{}
}

func (obj ServiceAccount) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespaceAndCompanyId().Name == "" {
		coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
		go AgentIndex{}.Build(obj.Obj.ObjectMeta.Labels["company"], obj.AgentName).Save()
	} else {
		err := obj.Update(ServiceAccount{Obj: obj.findByNameAndNamespaceAndCompanyId(), AgentName: obj.AgentName}, obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj ServiceAccount) findById() K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(ServiceAccount)
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj ServiceAccount) findByNameAndNamespaceAndCompanyId() K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.metadata.labels.company": obj.Obj.ObjectMeta.Labels["company"]},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(ServiceAccount)
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj ServiceAccount) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ServiceAccount) Update(oldObj interface{}, agent string) error {
	var oldObject ServiceAccount
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
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
	go AgentIndex{}.Build(obj.Obj.ObjectMeta.Labels["company"], obj.AgentName).Save()
	return nil
}

func (obj ServiceAccount) saveAll(objs []ServiceAccount) error {
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
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

func (object ServiceAccount) findAll() []K8sServiceAccount {
	query := bson.M{}
	objects := []ServiceAccount{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ServiceAccount)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sServiceAccount{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ServiceAccount) findByNamespace() []K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []ServiceAccount{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ServiceAccount)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sServiceAccount{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ServiceAccount) findBykubeAgentNameAndNamespace() []K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []ServiceAccount{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ServiceAccount)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sServiceAccount{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ServiceAccount) findBykubeAgentName() []K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []ServiceAccount{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ServiceAccount)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sServiceAccount{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ServiceAccount) findByName() K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(ServiceAccount)
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
