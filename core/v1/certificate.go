package v1

import (
	"context"
	"encoding/json"
	"github.com/go-bongo/bongo"
	"github.com/klovercloud-ci-cd/light-house-command/core/v1/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	CertificateCollection = "certificateCollection"
)

type K8sCertificate struct {
	TypeMeta   `json:",inline" bson:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata" bson:"metadata"`
	Spec       CertificateSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec" bson:"spec"`
	Status     CertificateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status" bson:"status"`
}
type Certificate struct {
	bongo.DocumentBase `bson:",inline"`
	Obj                K8sCertificate `bson:"obj" json:"obj"`
	AgentName          string         `bson:"agent_name" json:"agent_name"`
}

func (obj Certificate) findByNameAndNamespace() K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Certificate)
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func NewCertificate() KubeObject {
	return &Certificate{}
}

func (obj Certificate) Save(extra map[string]string) error {
	obj.AgentName = extra["agent_name"]
	if obj.findByNameAndNamespace().Name == "" {
		coll := db.GetDmManager().Db.Collection(CertificateCollection)
		_, err := coll.InsertOne(db.GetDmManager().Ctx, obj)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	} else {
		err := obj.Update(Certificate{Obj: obj.findByNameAndNamespace(), AgentName: obj.AgentName}, obj.AgentName)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
			return err
		}
	}
	return nil
}

func (obj Certificate) findById() K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.uid": obj.Obj.UID},
			{"agent_name": obj.AgentName},
		},
	}
	temp := new(Certificate)
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Certificate) Delete(agent string) error {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
			{"agent_name": agent},
		},
	}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	_, err := coll.DeleteOne(db.GetDmManager().Ctx, query)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func (obj Certificate) Update(oldObj interface{}, agent string) error {
	var oldObject Certificate
	body, _ := json.Marshal(oldObj)
	errorOfUnmarshal := json.Unmarshal(body, &oldObject)
	if errorOfUnmarshal != nil {
		return errorOfUnmarshal
	}
	if obj.AgentName == "" {
		obj.AgentName = agent
	}
	filter := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": obj.Obj.Name},
			{"obj.metadata.namespace": obj.Obj.Namespace},
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
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
	return nil
}

func (obj Certificate) saveAll(objs []Certificate) error {
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
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

func (object Certificate) findAll() []K8sCertificate {
	query := bson.M{}
	objects := []Certificate{}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Certificate)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sCertificate{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Certificate) findByNamespace() []K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
		},
	}
	objects := []Certificate{}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Certificate)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sCertificate{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Certificate) findBykubeAgentNameAndNamespace() []K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	objects := []Certificate{}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Certificate)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sCertificate{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Certificate) findBykubeAgentName() []K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"agent_name": object.AgentName},
		},
	}
	objects := []Certificate{}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	curser, _ := coll.Find(db.GetDmManager().Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(Certificate)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		objects = append(objects, *elemValue)
	}
	k8sObjects := []K8sCertificate{}
	for _, each := range objects {
		k8sObjects = append(k8sObjects, each.Obj)
	}
	return k8sObjects
}

func (object Certificate) findByName() K8sCertificate {
	query := bson.M{
		"$and": []bson.M{
			{"obj.metadata.name": object.Obj.Name},
			{"obj.metadata.namespace": object.Obj.Namespace},
			{"agent_name": object.AgentName},
		},
	}
	temp := new(Certificate)
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	result := coll.FindOne(db.GetDmManager().Ctx, query)

	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return temp.Obj
}

func (obj Certificate) deleteAll() error {
	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(CertificateCollection)
	_, err := coll.DeleteMany(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Failed to Delete sa [ERROR]", err)
	}
	return err
}
