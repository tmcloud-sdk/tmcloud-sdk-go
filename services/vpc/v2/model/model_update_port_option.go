package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePortOption struct {
	Name *string `json:"name,omitempty"`

	SecurityGroups *[]string `json:"security_groups,omitempty"`

	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	ExtraDhcpOpts *[]ExtraDhcpOpt `json:"extra_dhcp_opts,omitempty"`
}

func (o UpdatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortOption struct{}"
	}

	return strings.Join([]string{"UpdatePortOption", string(data)}, " ")
}
