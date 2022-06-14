package v1

import (
	"github.com/go-bongo/bongo"
	"github.com/klovercloud-ci-cd/light-house-command/core/v1/db"
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
	coll := db.GetDmManager().Db.Collection(AgentIndexCollection)
	_, err := coll.InsertOne(db.GetDmManager().Ctx, a)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
}
