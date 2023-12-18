package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RpoAndRtoInfo struct {
	CheckPoint *string `json:"check_point,omitempty"`

	Delay *string `json:"delay,omitempty"`

	GtidSet *string `json:"gtid_set,omitempty"`

	Time *string `json:"time,omitempty"`
}

func (o RpoAndRtoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RpoAndRtoInfo struct{}"
	}

	return strings.Join([]string{"RpoAndRtoInfo", string(data)}, " ")
}
