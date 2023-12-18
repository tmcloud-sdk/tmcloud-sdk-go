package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallGroupOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	EgressFirewallPolicyId *string `json:"egress_firewall_policy_id,omitempty"`

	IngressFirewallPolicyId *string `json:"ingress_firewall_policy_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Ports *[]string `json:"ports,omitempty"`
}

func (o NeutronUpdateFirewallGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupOption", string(data)}, " ")
}
