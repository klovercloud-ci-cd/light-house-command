package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/klovercloud-ci-cd/light-house-command/core/v1/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const AgentIndexCollection = "agentIndexCollection"

type AgentIndex struct {
	bongo.DocumentBase `bson:",inline"`
	CompanyId          string `json:"company" bson:"company"`
	AgentName          string `json:"agent_name" bson:"agent_name"`
}

func (a AgentIndex) Build(companyId, agentName string) AgentIndex {
	if companyId == "" || agentName == "" {
		return AgentIndex{}
	}
	return AgentIndex{
		CompanyId: companyId,
		AgentName: agentName,
	}
}

func (a AgentIndex) Save() {
	if a.CompanyId == "" {
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"agent_name": a.AgentName},
			{"company": a.CompanyId},
		},
	}
	update := bson.M{
		"$set": a,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := db.GetDmManager().Db.Collection(AgentIndexCollection)
	err := coll.FindOneAndUpdate(db.GetDmManager().Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
	}
}
