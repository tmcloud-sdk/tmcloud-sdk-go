package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRedislogRequest struct {
	InstanceId string `json:"instance_id"`

	QueryTime *int32 `json:"query_time,omitempty"`

	LogType string `json:"log_type"`

	ReplicationId *string `json:"replication_id,omitempty"`
}

func (o CreateRedislogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogRequest struct{}"
	}

	return strings.Join([]string{"CreateRedislogRequest", string(data)}, " ")
}
