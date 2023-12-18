package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallRuleOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	Action *string `json:"action,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`

	DestinationIpAddress *string `json:"destination_ip_address,omitempty"`

	DestinationPort *string `json:"destination_port,omitempty"`

	SourceIpAddress *string `json:"source_ip_address,omitempty"`

	SourcePort *string `json:"source_port,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

func (o NeutronCreateFirewallRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleOption", string(data)}, " ")
}
