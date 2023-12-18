package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronListFirewallPoliciesResponse struct {
	FirewallPolicies *[]NeutronFirewallPolicy `json:"firewall_policies,omitempty"`

	FirewallPoliciesLinks *[]NeutronPageLink `json:"firewall_policies_links,omitempty"`
	HttpStatusCode        int                `json:"-"`
}

func (o NeutronListFirewallPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallPoliciesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallPoliciesResponse", string(data)}, " ")
}
