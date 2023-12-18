package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRedislogRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	LogType string `json:"log_type"`
}

func (o ListRedislogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedislogRequest struct{}"
	}

	return strings.Join([]string{"ListRedislogRequest", string(data)}, " ")
}
