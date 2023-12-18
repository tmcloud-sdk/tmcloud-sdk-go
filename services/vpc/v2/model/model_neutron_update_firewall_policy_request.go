package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallPolicyRequest struct {
	FirewallPolicyId string `json:"firewall_policy_id"`

	Body *NeutronUpdateFirewallPolicyRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallPolicyRequest", string(data)}, " ")
}
