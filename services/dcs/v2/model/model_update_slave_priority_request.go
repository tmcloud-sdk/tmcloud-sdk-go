package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateSlavePriorityRequest struct {
	InstanceId string `json:"instance_id"`

	GroupId string `json:"group_id"`

	NodeId string `json:"node_id"`

	Body *PriorityBody `json:"body,omitempty"`
}

func (o UpdateSlavePriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityRequest struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityRequest", string(data)}, " ")
}
