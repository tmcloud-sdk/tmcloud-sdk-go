package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EipInfo struct {
	EipId *string `json:"eip_id,omitempty"`

	EipAddress *string `json:"eip_address,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
