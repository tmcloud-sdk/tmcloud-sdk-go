package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowAutoEnlargePolicyResponse struct {
	SwitchOption *bool `json:"switch_option,omitempty"`

	LimitSize *int32 `json:"limit_size,omitempty"`

	TriggerThreshold *int32 `json:"trigger_threshold,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoEnlargePolicyResponse", string(data)}, " ")
}
