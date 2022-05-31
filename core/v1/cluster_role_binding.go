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

const (
	ClusterRoleBindingCollection = "clusterRoleBindingCollection"
)

type k8sClusterRoleBinding struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Subjects   []Subject `json:"subjects,omitempty" protobuf:"bytes,2,rep,name=subjects" bson:"subjects"`
	RoleRef    RoleRef   `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef" bson:"roleRef"`
}

type ClusterRoleBinding struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                k8sClusterRoleBinding `bson:"obj" json:"obj"`
	AgentName          string                `bson:"agent_name" json:"agent_name"`
}

func (obj ClusterRoleBinding) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete crb [ERROR]", err)
	}
	return err
}

func NewClusterRoleBinding() KubeObject {
	return &ClusterRoleBinding{}
}

func (obj ClusterRoleBinding) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByName().Name == "" {
		coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(ClusterRoleBinding{Obj:obj.findByName(), AgentName: obj.AgentName},obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj ClusterRoleBinding) findById() k8sClusterRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(ClusterRoleBinding)
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (object ClusterRoleBinding) findByName() k8sClusterRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(ClusterRoleBinding)
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (object ClusterRoleBinding) findBykubeAgentName() []k8sClusterRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []ClusterRoleBinding{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []k8sClusterRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj ClusterRoleBinding) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ClusterRoleBinding) Update(oldObj interface{},agent string) error {
	var oldObject ClusterRoleBinding
	body, _ := json.Marshal(oldObj)
	errorOfUnmarshal := json.Unmarshal(body, &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
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
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj ClusterRoleBinding) saveAll(objs []ClusterRoleBinding) error {
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
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

func (object ClusterRoleBinding) findAll() []k8sClusterRoleBinding {
	query := bson.M{}
	objects := []ClusterRoleBinding{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []k8sClusterRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ClusterRoleBinding) findByNamespace() []k8sClusterRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []ClusterRoleBinding{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []k8sClusterRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ClusterRoleBinding) findBykubeAgentNameAndNamespace() []k8sClusterRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []ClusterRoleBinding{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleBindingCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRoleBinding)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []k8sClusterRoleBinding{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}
