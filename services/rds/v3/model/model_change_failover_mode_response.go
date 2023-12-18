package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeFailoverModeResponse struct {
	InstanceId *string `json:"instanceId,omitempty"`

	ReplicationMode *string `json:"replicationMode,omitempty"`

	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeFailoverModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFailoverModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeFailoverModeResponse", string(data)}, " ")
}
