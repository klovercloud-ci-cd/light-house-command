package v1

import (
	"context"
	"github.com/go-bongo/bongo"
	"github.com/klovercloud/lighthouse-command/core/v1/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type FinalizerName string
type NamespacePhase string

const (
	NamespaceCollection                 = "namespaceCollection"
	FinalizerKubernetes  FinalizerName  = "kubernetes"
	NamespaceActive      NamespacePhase = "Active"
	NamespaceTerminating NamespacePhase = "Terminating"
)

type K8sNamespace struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       NamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     NamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}

type Namespace struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sNamespace `bson:"obj" json:"obj"`
	KubeClusterId      string       `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj Namespace) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete namespace [ERROR]", err)
	}
	return err
}

func (obj Namespace) saveByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	if obj.findByNamespaceAndClusterId().Name == "" {
		//obj.Obj.Kind="Namespace"
		//obj.Obj.APIVersion="v1"
		coll := db.GetDmManager().Db.Collection(NamespaceCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj Namespace) deleteByClusterId(clusterId string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": clusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func NewNamespace() KubeObject {
	return &Namespace{}
}

func (obj Namespace) save() error {
	if obj.findByNamespaceAndClusterId().Name == "" {
		//obj.Obj.Kind="Namespace"
		//obj.Obj.APIVersion="v1"
		coll := db.GetDmManager().Db.Collection(NamespaceCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (object Namespace) findByNamespace() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (object Namespace) findByNamespaceAndClusterId() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
func (obj Namespace) findById() K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(Namespace)
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Namespace) findBykubeClusterId() []K8sNamespace {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	namespaces := []Namespace{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Namespace)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		namespaces = append(namespaces, *elemValue)
	}
	k8sObjects := []K8sNamespace{}
	for _, each := range namespaces {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Namespace) findAll() []K8sNamespace {
	query := bson.M{}
	namespaces := []Namespace{}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Namespace)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		namespaces = append(namespaces, *elemValue)
	}
	k8sObjects := []K8sNamespace{}
	for _, each := range namespaces {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Namespace) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj Namespace) update() error {
	//obj.Obj.Kind="Namespace"
	//obj.Obj.APIVersion="v1"
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
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj Namespace) saveAll(objs []Namespace) error {
	coll := db.GetDmManager().Db.Collection(NamespaceCollection)
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
