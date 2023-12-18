package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronInsertFirewallRuleRequestBody struct {
	FirewallRuleId string `json:"firewall_rule_id"`

	InsertAfter *string `json:"insert_after,omitempty"`

	InsertBefore *string `json:"insert_before,omitempty"`
}

func (o NeutronInsertFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronInsertFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronInsertFirewallRuleRequestBody", string(data)}, " ")
}
