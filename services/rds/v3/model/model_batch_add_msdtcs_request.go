package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddMsdtcsRequest struct {
	InstanceId string `json:"instance_id"`

	Body *AddMsdtcRequestBody `json:"body,omitempty"`
}

func (o BatchAddMsdtcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMsdtcsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddMsdtcsRequest", string(data)}, " ")
}
