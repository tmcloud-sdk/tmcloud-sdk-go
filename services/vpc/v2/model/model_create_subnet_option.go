package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Cidr string `json:"cidr"`

	VpcId string `json:"vpc_id"`

	GatewayIp string `json:"gateway_ip"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	DhcpEnable *bool `json:"dhcp_enable,omitempty"`

	PrimaryDns *string `json:"primary_dns,omitempty"`

	SecondaryDns *string `json:"secondary_dns,omitempty"`

	DnsList *[]string `json:"dnsList,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	ExtraDhcpOpts *[]ExtraDhcpOption `json:"extra_dhcp_opts,omitempty"`
}

func (o CreateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetOption struct{}"
	}

	return strings.Join([]string{"CreateSubnetOption", string(data)}, " ")
}
