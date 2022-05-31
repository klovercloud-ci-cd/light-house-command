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
	NodeCollection = "NodeCollection"
)

type K8sNode struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type Node struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNode `bson:"obj" json:"obj"`
	AgentName          string  `bson:"agent_name" json:"agent_name"`
}

func (obj Node) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete node [ERROR]", err)
	}
	return err
}

func NewNode() KubeObject {
	return &Node{}
}

func (obj Node) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndAgentName().Name == "" {
		coll := db.GetDmManager().Db.Collection(NodeCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(Node{Obj:obj.findByNameAndAgentName(), AgentName: obj.AgentName},obj.AgentName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (object Node) findByNameAndAgentName() K8sNode {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Node)
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Node) findById() K8sNode {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Node)
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Node) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj Node) Update(oldObj interface{},agent string) error {
	var oldObject Node
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
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj Node) saveAll(objs []Node) error {
	coll := db.GetDmManager().Db.Collection(NodeCollection)
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

func (object Node) findAll() []K8sNode {
	query := bson.M{}
	objects := []Node{}
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Node)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNode{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Node) findBykubeAgentName() []K8sNode {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []Node{}
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Node)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sNode{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Node) findByName() K8sNode {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Node)
	coll := db.GetDmManager().Db.Collection(NodeCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
