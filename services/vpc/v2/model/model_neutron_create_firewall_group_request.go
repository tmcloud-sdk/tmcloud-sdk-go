package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallGroupRequest struct {
	Body *NeutronCreateFirewallGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFirewallGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallGroupRequest", string(data)}, " ")
}
