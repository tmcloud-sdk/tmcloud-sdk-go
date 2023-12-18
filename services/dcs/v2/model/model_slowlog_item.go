package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowlogItem struct {
	Id *int32 `json:"id,omitempty"`

	Command *string `json:"command,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	Duration *string `json:"duration,omitempty"`

	ShardName *string `json:"shard_name,omitempty"`
}

func (o SlowlogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogItem struct{}"
	}

	return strings.Join([]string{"SlowlogItem", string(data)}, " ")
}
