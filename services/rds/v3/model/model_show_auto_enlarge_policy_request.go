package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowAutoEnlargePolicyRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyRequest", string(data)}, " ")
}
