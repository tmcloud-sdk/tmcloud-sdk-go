package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RevokeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *RevokeRequestBody `json:"body,omitempty"`
}

func (o RevokeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRequest struct{}"
	}

	return strings.Join([]string{"RevokeRequest", string(data)}, " ")
}
