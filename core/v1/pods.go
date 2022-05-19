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
	PodCollection = "podCollection"
)

type K8sPod struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type Pod struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sPod `bson:"obj" json:"obj"`
	KubeClusterId      string `json:"kubeClusterId" bson:"kubeClusterId"`
}

func (obj Pod) saveByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	if obj.findByNameAndNamespace().Name == "" {
		//log.Println("CPU usages:",obj.Obj.Spec.Containers[0].Resources.Requests["cpu"])
		coll := db.GetDmManager().Db.Collection(PodCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		log.Println("saving pod:", obj.Obj.Name+" exists!")
		err := obj.delete()
		if err != nil {
			return err
		}
		log.Println("again saving pod status:", obj.Obj.Status.Phase)
		coll := db.GetDmManager().Db.Collection(PodCollection)
		_, err = coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	//obj.calculateResource().add()
	return nil
}

func (obj Pod) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete pod [ERROR]", err)
	}
	return err
}

func (obj Pod) deleteByClusterId(clusterId string) error {
	obj.KubeClusterId = clusterId
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.kind": "Pod"},
			{"kubeClusterId": clusterId},
		},
	}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to delete pod [ERROR]", err)
	}
	//obj.calculateResource().remove()
	return err
}

func NewPod() KubeObject {
	return &Pod{}
}

func (obj Pod) save() error {
	log.Println("saving pod status:", obj.Obj.Status.Phase)
	if obj.findByNameAndNamespace().Name == "" {
		log.Println("saving pod:", obj.Obj.Name+" not exists!")
		//log.Println("CPU usages:",obj.Obj.Spec.Containers[0].Resources.Requests["cpu"])
		coll := db.GetDmManager().Db.Collection(PodCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		log.Println("saving pod:", obj.Obj.Name+" exists!")
		err := obj.delete()
		if err != nil {
			return err
		}
		log.Println("again saving pod status:", obj.Obj.Status.Phase)
		coll := db.GetDmManager().Db.Collection(PodCollection)
		_, err = coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}

	//obj.calculateResource().add()
	return nil
}

func (obj Pod) findById() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Pod) findByNameAndNamespace() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"obj.kind": "Pod"},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Pod) findByLabel() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.labels": obj.Obj.Labels},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (obj Pod) delete() error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"kubeClusterId": obj.KubeClusterId},
		},
	}
	log.Println("deleting pod:", obj.Obj.Name+"!")
	coll := db.GetDmManager().Db.Collection(PodCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("[ERROR]", err)
	}
	//	obj.calculateResource().remove()
	return err
}

func (obj Pod) update() error {

	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
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
	coll := db.GetDmManager().Db.Collection(PodCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}

	return nil
}

func (obj Pod) saveAll(objs []Pod) error {
	coll := db.GetDmManager().Db.Collection(PodCollection)
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

func (object Pod) findAll() []K8sPod {
	query := bson.M{}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findByNamespace() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findBykubeClusterIdAndNamespace() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findBykubeClusterId() []K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	objects := []Pod{}
	coll := db.GetDmManager().Db.Collection(PodCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Pod)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sPod{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Pod) findByName() K8sPod {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"kubeClusterId": object.KubeClusterId},
		},
	}
	temp := new(Pod)
	coll := db.GetDmManager().Db.Collection(PodCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}
