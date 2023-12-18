package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowReplicationStatusRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowReplicationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationStatusRequest", string(data)}, " ")
}
