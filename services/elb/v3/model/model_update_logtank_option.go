package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateLogtankOption struct {
	LogGroupId *string `json:"log_group_id,omitempty"`

	LogTopicId *string `json:"log_topic_id,omitempty"`
}

func (o UpdateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankOption struct{}"
	}

	return strings.Join([]string{"UpdateLogtankOption", string(data)}, " ")
}
