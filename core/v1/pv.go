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
	PVCollection = "pvCollection"
)

type K8sPersistentVolume struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       PersistentVolumeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     PersistentVolumeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type PersistentVolume struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sPersistentVolume `bson:"obj" json:"obj"`
	KubeClusterId      string              `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj PersistentVolume) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(PVCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete pv [ERROR]", err)
	}
	return err
}
func (obj PersistentVolume) saveByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	if obj.findByNameAndClusterId().Name == "" {
		coll := db.GetDmManager().Db.Collection(PVCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj PersistentVolume) deleteByClusterId(clusterId string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": clusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(PVCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func NewPersistentVolume() KubeObject {
	return &PersistentVolume{}
}

func (obj PersistentVolume) save() error {
	if obj.findByNameAndClusterId().Name == "" {
		coll := db.GetDmManager().Db.Collection(PVCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj PersistentVolume) findById() K8sPersistentVolume {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(PersistentVolume)
	coll := db.GetDmManager().Db.Collection(PVCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj PersistentVolume) findByNameAndClusterId() K8sPersistentVolume {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(PersistentVolume)
	coll := db.GetDmManager().Db.Collection(PVCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj PersistentVolume) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(PVCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj PersistentVolume) update() error {

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
	coll := db.GetDmManager().Db.Collection(PVCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj PersistentVolume) saveAll(objs []PersistentVolume) error {
	coll := db.GetDmManager().Db.Collection(PVCollection)
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

func (object PersistentVolume) findAll() []K8sPersistentVolume {
	query := bson.M{}
	objects := []PersistentVolume{}
	coll := db.GetDmManager().Db.Collection(PVCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolume)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolume{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object PersistentVolume) findBykubeClusterId() []K8sPersistentVolume {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	objects := []PersistentVolume{}
	coll := db.GetDmManager().Db.Collection(PVCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolume)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolume{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj PersistentVolume) findByName() K8sPersistentVolume {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(PersistentVolume)
	coll := db.GetDmManager().Db.Collection(PVCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
