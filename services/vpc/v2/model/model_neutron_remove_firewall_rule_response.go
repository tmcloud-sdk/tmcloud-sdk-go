package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronRemoveFirewallRuleResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	FirewallRules *[]string `json:"firewall_rules,omitempty"`

	Audited *bool `json:"audited,omitempty"`

	Public *bool `json:"public,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o NeutronRemoveFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleResponse", string(data)}, " ")
}
