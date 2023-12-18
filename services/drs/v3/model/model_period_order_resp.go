package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PeriodOrderResp struct {
	Status *string `json:"status,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	PeriodType *int32 `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	EffTime *string `json:"eff_time,omitempty"`

	ExpTime *string `json:"exp_time,omitempty"`
}

func (o PeriodOrderResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderResp struct{}"
	}

	return strings.Join([]string{"PeriodOrderResp", string(data)}, " ")
}
