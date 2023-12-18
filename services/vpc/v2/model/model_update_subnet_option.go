package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubnetOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	PrimaryDns *string `json:"primary_dns,omitempty"`

	SecondaryDns *string `json:"secondary_dns,omitempty"`

	DnsList *[]string `json:"dnsList,omitempty"`

	ExtraDhcpOpts *[]ExtraDhcpOption `json:"extra_dhcp_opts,omitempty"`
}

func (o UpdateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetOption struct{}"
	}

	return strings.Join([]string{"UpdateSubnetOption", string(data)}, " ")
}
