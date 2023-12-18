package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateFirewallPolicyRequestBody struct {
	FirewallPolicy *NeutronUpdateFirewallPolicyOption `json:"firewall_policy"`
}

func (o NeutronUpdateFirewallPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallPolicyRequestBody", string(data)}, " ")
}
