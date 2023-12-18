package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainNameResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	DnsName *string `json:"dns_name,omitempty"`

	DnsType *string `json:"dns_type,omitempty"`

	Ipv4Address *string `json:"ipv4_address,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainNameResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainNameResponse", string(data)}, " ")
}
