package v1

import (
	"context"
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
	KubeClusterId      string         `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj ClusterRole) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete CR [ERROR]", err)
	}
	return err
}

func NewClusterRole() KubeObject {
	return &ClusterRole{}
}

func (obj ClusterRole) save() error {
	if obj.findByName().Name == "" {
		coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj ClusterRole) findById() K8sClusterRole {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
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
			{"kubeClusterId": obj.KubeClusterId},
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

func (obj ClusterRole) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ClusterRole) deleteAllBykubeClusterId() error {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(ClusterRoleCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ClusterRole) update() error {

	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
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

func (object ClusterRole) findBykubeClusterId() []K8sClusterRole {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
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
