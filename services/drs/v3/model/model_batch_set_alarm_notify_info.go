package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetAlarmNotifyInfo struct {
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty"`

	TopicUrn *string `json:"topic_urn,omitempty"`

	DelayTime *int64 `json:"delay_time,omitempty"`

	RtoDelay *int64 `json:"rto_delay,omitempty"`

	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	AlarmToUser *bool `json:"alarm_to_user,omitempty"`
}

func (o BatchSetAlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"BatchSetAlarmNotifyInfo", string(data)}, " ")
}
