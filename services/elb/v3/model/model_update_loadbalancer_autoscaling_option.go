package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateLoadbalancerAutoscalingOption struct {
	Enable bool `json:"enable"`

	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o UpdateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerAutoscalingOption", string(data)}, " ")
}
