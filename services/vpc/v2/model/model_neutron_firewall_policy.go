package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronFirewallPolicy struct {
	Audited bool `json:"audited"`

	Description string `json:"description"`

	FirewallRules []string `json:"firewall_rules"`

	Id string `json:"id"`

	Name string `json:"name"`

	Public bool `json:"public"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`
}

func (o NeutronFirewallPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronFirewallPolicy struct{}"
	}

	return strings.Join([]string{"NeutronFirewallPolicy", string(data)}, " ")
}
