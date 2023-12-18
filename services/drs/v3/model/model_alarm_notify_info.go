package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AlarmNotifyInfo struct {
	DelayTime *int64 `json:"delay_time,omitempty"`

	RtoDelay *int64 `json:"rto_delay,omitempty"`

	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	AlarmToUser *bool `json:"alarm_to_user,omitempty"`

	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty"`
}

func (o AlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"AlarmNotifyInfo", string(data)}, " ")
}
