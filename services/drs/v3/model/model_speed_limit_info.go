package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SpeedLimitInfo struct {
	Begin string `json:"begin"`

	End string `json:"end"`

	Speed string `json:"speed"`

	IsUtc *bool `json:"is_utc,omitempty"`
}

func (o SpeedLimitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimitInfo struct{}"
	}

	return strings.Join([]string{"SpeedLimitInfo", string(data)}, " ")
}
