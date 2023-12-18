package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AutoscalingRef struct {
	Enable bool `json:"enable"`

	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o AutoscalingRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingRef struct{}"
	}

	return strings.Join([]string{"AutoscalingRef", string(data)}, " ")
}
