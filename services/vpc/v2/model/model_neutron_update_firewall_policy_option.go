package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallPolicyOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	FirewallRules *[]string `json:"firewall_rules,omitempty"`

	Audited *bool `json:"audited,omitempty"`
}

func (o NeutronUpdateFirewallPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallPolicyOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallPolicyOption", string(data)}, " ")
}
