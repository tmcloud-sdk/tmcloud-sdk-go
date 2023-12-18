package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7PolicyRequest struct {
	L7policyId string `json:"l7policy_id"`

	Body *UpdateL7PolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyRequest", string(data)}, " ")
}
