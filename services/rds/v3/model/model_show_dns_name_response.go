package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDnsNameResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	DnsName *string `json:"dns_name,omitempty"`

	DnsType *string `json:"dns_type,omitempty"`

	Ipv6Address *string `json:"ipv6_address,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDnsNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDnsNameResponse struct{}"
	}

	return strings.Join([]string{"ShowDnsNameResponse", string(data)}, " ")
}
