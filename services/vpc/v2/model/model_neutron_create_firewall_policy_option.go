package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallPolicyOption struct {
	Audited *bool `json:"audited,omitempty"`

	Description *string `json:"description,omitempty"`

	FirewallRules *[]string `json:"firewall_rules,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o NeutronCreateFirewallPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyOption", string(data)}, " ")
}
