package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronListFirewallRulesResponse struct {
	FirewallRules *[]NeutronFirewallRule `json:"firewall_rules,omitempty"`

	FirewallRulesLinks *[]NeutronPageLink `json:"firewall_rules_links,omitempty"`
	HttpStatusCode     int                `json:"-"`
}

func (o NeutronListFirewallRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallRulesResponse", string(data)}, " ")
}
