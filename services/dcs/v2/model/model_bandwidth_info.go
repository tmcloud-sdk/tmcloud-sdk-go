package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BandwidthInfo struct {
	BeginTime *int64 `json:"begin_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	CurrentTime *int64 `json:"current_time,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	NextExpandTime *int64 `json:"next_expand_time,omitempty"`

	ExpandCount *int32 `json:"expand_count,omitempty"`

	ExpandEffectTime *int64 `json:"expand_effect_time,omitempty"`

	ExpandIntervalTime *int64 `json:"expand_interval_time,omitempty"`

	MaxExpandCount *int32 `json:"max_expand_count,omitempty"`

	TaskRunning *bool `json:"task_running,omitempty"`
}

func (o BandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfo struct{}"
	}

	return strings.Join([]string{"BandwidthInfo", string(data)}, " ")
}
