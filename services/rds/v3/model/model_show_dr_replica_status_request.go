package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDrReplicaStatusRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowDrReplicaStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrReplicaStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDrReplicaStatusRequest", string(data)}, " ")
}
