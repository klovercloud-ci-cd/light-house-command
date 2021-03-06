package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/light-house-command/api/common"
	v1 "github.com/klovercloud-ci-cd/light-house-command/core/v1"
	"github.com/klovercloud-ci-cd/light-house-command/enums"
	"github.com/labstack/echo/v4"
	"log"
)

func Router(g *echo.Group) {
	KubeEvents(g.Group("/kube_events"))
}

func KubeEvents(g *echo.Group) {
	g.POST("", StoreKubeEvents)
}

// Post... Post Api
// @Summary Post api
// @Description Api for storing all kube events
// @Tags KubeEvents
// @Produce json
// @Success 200 {object} common.ResponseDTO{data=v1.KubeEventMessage{}.Body{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/kube_events [POST]
func StoreKubeEvents(context echo.Context) error {
	var kubeEvents v1.KubeEventMessage
	if err := context.Bind(&kubeEvents); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	type TempBody struct {
		Obj interface{} `json:"obj"`
	}
	extra := make(map[string]string)
	if kubeEvents.Header.Command == enums.UPDATE {
		type KubeObjectForUpdate struct {
			OldK8sObj interface{} `json:"old_k8s_obj" bson:"old_k8s_obj"`
			NewK8sObj interface{} `json:"new_k8s_obj" bson:"new_k8s_obj"`
		}
		var body KubeObjectForUpdate
		data, err := json.MarshalIndent(kubeEvents.Body, "", "  ")
		if err != nil {
			log.Println("marshaling error: ", err.Error())
		}
		err = json.Unmarshal(data, &body)
		if err != nil {
			log.Println("Unmarshalling error: ", err.Error())
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		var oldKubeObject v1.KubeObject
		oldKubeObject = v1.GetObject(enums.RESOURCE_TYPE(kubeEvents.Header.Extras["object"]))
		var newKubeObject v1.KubeObject
		newKubeObject = v1.GetObject(enums.RESOURCE_TYPE(kubeEvents.Header.Extras["object"]))

		var tempOldBody TempBody
		tempOldBody.Obj = body.OldK8sObj
		old, err := json.MarshalIndent(tempOldBody, "", "  ")
		err = json.Unmarshal(old, &oldKubeObject)

		if err != nil {
			log.Println("marshaling error: ", err.Error())
			return common.GenerateErrorResponse(context, nil, err.Error())
		}

		var tempNewBody TempBody
		tempNewBody.Obj = body.NewK8sObj
		newObj, err := json.MarshalIndent(tempNewBody, "", "  ")
		err = json.Unmarshal(newObj, &newKubeObject)
		if err != nil {
			log.Println("marshaling error: ", err.Error())
			log.Println(err.Error())
		}
		err = newKubeObject.Update(oldKubeObject, kubeEvents.Header.Extras["agent"])
		if err != nil {
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		return common.GenerateSuccessResponse(context, newKubeObject, nil, "Successfully Updated!")
	} else if kubeEvents.Header.Command == enums.ADD {
		var kubeObject v1.KubeObject
		kubeObject = v1.GetObject(enums.RESOURCE_TYPE(kubeEvents.Header.Extras["object"]))
		var tempOldBody TempBody
		tempOldBody.Obj = kubeEvents.Body
		old, err := json.MarshalIndent(tempOldBody, "", "  ")
		if err != nil {
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		err = json.Unmarshal(old, &kubeObject)
		if err != nil {
			log.Println("marshaling error: ", err.Error())
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		if kubeEvents.Header.Extras != nil {
			if res, ok := kubeEvents.Header.Extras["agent"]; ok {
				extra["agent_name"] = res
			}
		}
		err = kubeObject.Save(extra)
		if err != nil {
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		return common.GenerateSuccessResponse(context, kubeEvents.Body, nil, "Successfully Added!")
	} else if kubeEvents.Header.Command == enums.DELETE {
		var kubeObject v1.KubeObject
		kubeObject = v1.GetObject(enums.RESOURCE_TYPE(kubeEvents.Header.Extras["object"]))
		var tempOldBody TempBody
		tempOldBody.Obj = kubeEvents.Body
		old, err := json.MarshalIndent(tempOldBody, "", "  ")
		if err != nil {
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		err = json.Unmarshal(old, &kubeObject)
		if err != nil {
			log.Println("marshaling error: ", err.Error())
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		err = kubeObject.Delete(kubeEvents.Header.Extras["agent"])
		if err != nil {
			return common.GenerateErrorResponse(context, nil, err.Error())
		}
		return common.GenerateSuccessResponse(context, kubeEvents.Body, nil, "Successfully Deleted!")
	}
	return nil
}
