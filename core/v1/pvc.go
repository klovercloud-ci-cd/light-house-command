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
	PVCCollection = "pvcCollection"
)

type K8sPersistentVolumeClaim struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       PersistentVolumeClaimSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     PersistentVolumeClaimStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type PersistentVolumeClaim struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sPersistentVolumeClaim `bson:"obj" json:"obj"`
	KubeClusterId      string                   `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj PersistentVolumeClaim) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete pvc [ERROR]", err)
	}
	return err
}

func NewPersistentVolumeClaim() KubeObject {
	return &PersistentVolumeClaim{}
}

func (obj PersistentVolumeClaim) save() error {
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(PVCCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj PersistentVolumeClaim) findByNameAndClusterId() K8sPersistentVolume {
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

func (obj PersistentVolumeClaim) findById() K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(PersistentVolumeClaim)
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj PersistentVolumeClaim) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj PersistentVolumeClaim) update() error {

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
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj PersistentVolumeClaim) saveAll(objs []PersistentVolumeClaim) error {
	coll := db.GetDmManager().Db.Collection(PVCCollection)
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

func (object PersistentVolumeClaim) findAll() []K8sPersistentVolumeClaim {
	query := bson.M{}
	objects := []PersistentVolumeClaim{}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolumeClaim)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolumeClaim{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object PersistentVolumeClaim) findByNamespace() []K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []PersistentVolumeClaim{}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolumeClaim)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolumeClaim{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object PersistentVolumeClaim) findByNameAndNamespace() K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"obj.metadata.name": object.Obj.Name},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	temp := new(PersistentVolumeClaim)
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (object PersistentVolumeClaim) findBykubeClusterIdAndNamespace() []K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	objects := []PersistentVolumeClaim{}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolumeClaim)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolumeClaim{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object PersistentVolumeClaim) findBykubeClusterId() []K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	objects := []PersistentVolumeClaim{}
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(PersistentVolumeClaim)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPersistentVolumeClaim{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object PersistentVolumeClaim) findByName() K8sPersistentVolumeClaim {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	temp := new(PersistentVolumeClaim)
	coll := db.GetDmManager().Db.Collection(PVCCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
