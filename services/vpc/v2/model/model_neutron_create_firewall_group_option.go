package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallGroupOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	IngressFirewallPolicyId *string `json:"ingress_firewall_policy_id,omitempty"`

	EgressFirewallPolicyId *string `json:"egress_firewall_policy_id,omitempty"`

	Ports *[]string `json:"ports,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

func (o NeutronCreateFirewallGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallGroupOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallGroupOption", string(data)}, " ")
}
