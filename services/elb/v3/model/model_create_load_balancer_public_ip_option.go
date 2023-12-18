package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadBalancerPublicIpOption struct {
	IpVersion *int32 `json:"ip_version,omitempty"`

	NetworkType string `json:"network_type"`

	BillingInfo *string `json:"billing_info,omitempty"`

	Description *string `json:"description,omitempty"`

	Bandwidth *CreateLoadBalancerBandwidthOption `json:"bandwidth"`
}

func (o CreateLoadBalancerPublicIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerPublicIpOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerPublicIpOption", string(data)}, " ")
}
