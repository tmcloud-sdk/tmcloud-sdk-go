package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronFirewallRule struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Action string `json:"action"`

	Protocol string `json:"protocol"`

	IpVersion int32 `json:"ip_version"`

	Enabled bool `json:"enabled"`

	Public bool `json:"public"`

	DestinationIpAddress string `json:"destination_ip_address"`

	DestinationPort string `json:"destination_port"`

	SourceIpAddress string `json:"source_ip_address"`

	SourcePort string `json:"source_port"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`
}

func (o NeutronFirewallRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronFirewallRule struct{}"
	}

	return strings.Join([]string{"NeutronFirewallRule", string(data)}, " ")
}
