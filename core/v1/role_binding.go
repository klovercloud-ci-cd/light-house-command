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
	KubeClusterId      string         `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj RoleBinding) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete rb [ERROR]", err)
	}
	return err
}

func (obj RoleBinding) saveByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj RoleBinding) deleteByClusterId(clusterId string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": clusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func NewRoleBinding() KubeObject {
	return &RoleBinding{}
}

func (obj RoleBinding) save() error {
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj RoleBinding) findById() K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
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
			{"kubeClusterId": obj.KubeClusterId},
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

func (obj RoleBinding) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(RoleBindingCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj RoleBinding) update() error {

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

func (object RoleBinding) findBykubeClusterIdAndNamespace() []K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
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

func (object RoleBinding) findBykubeClusterId() []K8sRoleBinding {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
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
			{"kubeClusterId": object.KubeClusterId},
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
