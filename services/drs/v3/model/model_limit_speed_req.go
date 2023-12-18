package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LimitSpeedReq struct {
	JobId string `json:"job_id"`

	SpeedLimit []SpeedLimitInfo `json:"speed_limit"`
}

func (o LimitSpeedReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitSpeedReq struct{}"
	}

	return strings.Join([]string{"LimitSpeedReq", string(data)}, " ")
}
