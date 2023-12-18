package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListLogtanksRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	LogGroupId *[]string `json:"log_group_id,omitempty"`

	LogTopicId *[]string `json:"log_topic_id,omitempty"`
}

func (o ListLogtanksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtanksRequest struct{}"
	}

	return strings.Join([]string{"ListLogtanksRequest", string(data)}, " ")
}
