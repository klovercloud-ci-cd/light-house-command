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
	ClusterRoleCollection = "clusterRoleCollection"
)

type K8sClusterRole struct {
	TypeMeta        `json:",inline" bson:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Rules           []PolicyRule     `json:"rules" protobuf:"bytes,2,rep,name=rules" bson:"rules"`
	AggregationRule *AggregationRule `json:"aggregationRule,omitempty" protobuf:"bytes,3,opt,name=aggregationRule" bson:"aggregationRule"`
}
type ClusterRole struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sClusterRole `bson:"obj" json:"obj"`
	AgentName          string         `bson:"agent_name" json:"agent_name"`
}

func (obj ClusterRole) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete CR [ERROR]", err)
	}
	return err
}

func NewClusterRole() KubeObject {
	return &ClusterRole{}
}

func (obj ClusterRole) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByName().Name == "" {
		coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(obj.findByName())
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj ClusterRole) findById() K8sClusterRole {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(ClusterRole)
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj ClusterRole) findByName() K8sClusterRole {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(ClusterRole)
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj ClusterRole) Delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ClusterRole) deleteAllBykubeAgentName() error {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": obj.AgentName},
		},
	}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ClusterRole) Update(oldObj interface{}) error {
	var oldObject ClusterRole
	errorOfUnmarshal := json.Unmarshal([]byte(oldObj.(string)), &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
	}

	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": oldObject.Obj.UID},
			{"agent_name": obj.AgentName},
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
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj ClusterRole) saveAll(objs []ClusterRole) error {
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
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

func (object ClusterRole) findAll() []K8sClusterRole {
	query := bson.M{}
	objects := []ClusterRole{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRole)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sClusterRole{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object ClusterRole) findBykubeAgentName() []K8sClusterRole {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []ClusterRole{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(ClusterRole)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sClusterRole{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}
