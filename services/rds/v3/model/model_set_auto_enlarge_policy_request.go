package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetAutoEnlargePolicyRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`

	Body *CustomerModifyAutoEnlargePolicyReq `json:"body,omitempty"`
}

func (o SetAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePolicyRequest", string(data)}, " ")
}
