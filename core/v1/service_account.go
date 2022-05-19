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
	KubeClusterId      string            `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj ServiceAccount) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete sa [ERROR]", err)
	}
	return err
}
func (obj ServiceAccount) saveByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj ServiceAccount) deleteByClusterId(clusterId string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": clusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func NewServiceAccount() KubeObject {
	return &ServiceAccount{}
}

func (obj ServiceAccount) save() error {
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj ServiceAccount) findById() K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
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

func (obj ServiceAccount) findByNameAndNamespace() K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"kubeClusterId": obj.KubeClusterId},
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

func (obj ServiceAccount) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj ServiceAccount) update() error {
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
	coll := db.GetDmManager().Db.Collection(ServiceAccountCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
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

func (object ServiceAccount) findBykubeClusterIdAndNamespace() []K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
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

func (object ServiceAccount) findBykubeClusterId() []K8sServiceAccount {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
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
			{"kubeClusterId": object.KubeClusterId},
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
