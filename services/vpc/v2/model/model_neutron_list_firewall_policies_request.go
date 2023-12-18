package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronListFirewallPoliciesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Description *[]string `json:"description,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListFirewallPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallPoliciesRequest struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallPoliciesRequest", string(data)}, " ")
}
