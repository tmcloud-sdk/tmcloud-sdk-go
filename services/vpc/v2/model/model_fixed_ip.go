package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FixedIp struct {
	IpAddress *string `json:"ip_address,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o FixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixedIp struct{}"
	}

	return strings.Join([]string{"FixedIp", string(data)}, " ")
}
