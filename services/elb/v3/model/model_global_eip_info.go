package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlobalEipInfo struct {
	GlobalEipId *string `json:"global_eip_id,omitempty"`

	GlobalEipAddress *string `json:"global_eip_address,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o GlobalEipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalEipInfo struct{}"
	}

	return strings.Join([]string{"GlobalEipInfo", string(data)}, " ")
}
