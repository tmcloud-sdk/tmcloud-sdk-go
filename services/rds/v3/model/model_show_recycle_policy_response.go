package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowRecyclePolicyResponse struct {
	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`
	HttpStatusCode        int    `json:"-"`
}

func (o ShowRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyResponse", string(data)}, " ")
}
