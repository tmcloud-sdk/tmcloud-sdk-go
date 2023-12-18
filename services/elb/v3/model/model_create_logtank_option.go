package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankOption struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	LogGroupId string `json:"log_group_id"`

	LogTopicId string `json:"log_topic_id"`
}

func (o CreateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankOption struct{}"
	}

	return strings.Join([]string{"CreateLogtankOption", string(data)}, " ")
}
