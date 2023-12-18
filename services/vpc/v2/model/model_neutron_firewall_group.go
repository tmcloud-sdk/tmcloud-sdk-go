package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronFirewallGroup struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	AdminStateUp bool `json:"admin_state_up"`

	EgressFirewallPolicyId string `json:"egress_firewall_policy_id"`

	IngressFirewallPolicyId string `json:"ingress_firewall_policy_id"`

	Ports []string `json:"ports"`

	Public bool `json:"public"`

	Status string `json:"status"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronFirewallGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronFirewallGroup struct{}"
	}

	return strings.Join([]string{"NeutronFirewallGroup", string(data)}, " ")
}
