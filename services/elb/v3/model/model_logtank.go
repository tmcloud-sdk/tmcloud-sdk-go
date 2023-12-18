package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Logtank struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	LoadbalancerId string `json:"loadbalancer_id"`

	LogGroupId string `json:"log_group_id"`

	LogTopicId string `json:"log_topic_id"`
}

func (o Logtank) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Logtank struct{}"
	}

	return strings.Join([]string{"Logtank", string(data)}, " ")
}
