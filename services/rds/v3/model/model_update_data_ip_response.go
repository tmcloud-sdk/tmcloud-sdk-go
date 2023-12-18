package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateDataIpResponse struct {
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataIpResponse", string(data)}, " ")
}
