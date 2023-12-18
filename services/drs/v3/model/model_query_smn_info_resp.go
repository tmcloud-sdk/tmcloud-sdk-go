package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QuerySmnInfoResp struct {
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty"`

	TopicName *string `json:"topic_name,omitempty"`

	DelayTime *int64 `json:"delay_time,omitempty"`

	RtoDelay *int64 `json:"rto_delay,omitempty"`

	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	AlarmToUser *bool `json:"alarm_to_user,omitempty"`
}

func (o QuerySmnInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySmnInfoResp struct{}"
	}

	return strings.Join([]string{"QuerySmnInfoResp", string(data)}, " ")
}
