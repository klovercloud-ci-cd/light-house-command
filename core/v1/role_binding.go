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
	RoleBindingCollection = "roleBindingCollection"
)

type K8sRoleBinding struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Subjects   []Subject `json:"subjects,omitempty" protobuf:"bytes,2,rep,name=subjects" bson:"subjects"`
	RoleRef    RoleRef   `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef" bson:"roleRef"`
}

type RoleBinding struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sRoleBinding `bson:"obj" json:"obj"`
	AgentName          string         `bson:"agent_name" json:"agent_name"`
}

func (obj RoleBinding) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete rb [ERROR]", err)
	}
	return err
}

func NewRoleBinding() KubeObject {
	return &RoleBinding{}
}

func (obj RoleBinding) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(RoleBinding{Obj: obj.findByNameAndNamespace(), AgentName: obj.AgentName}, obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj RoleBinding) findById() K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(RoleBinding)
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
func (obj RoleBinding) findByNameAndNamespace() K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(RoleBinding)
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj RoleBinding) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj RoleBinding) Update(oldObj interface{}, agent string) error {
	var oldObject RoleBinding
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
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj RoleBinding) saveAll(objs []RoleBinding) error {
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
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

func (object RoleBinding) findAll() []K8sRoleBinding {
	query := bson.M{}
	objects := []RoleBinding{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(RoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object RoleBinding) findByNamespace() []K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []RoleBinding{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(RoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object RoleBinding) findBykubeAgentNameAndNamespace() []K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []RoleBinding{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(RoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object RoleBinding) findBykubeAgentName() []K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []RoleBinding{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(RoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object RoleBinding) findByName() K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(RoleBinding)
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
