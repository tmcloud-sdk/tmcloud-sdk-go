package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreatePortOption struct {
	Name *string `json:"name,omitempty"`

	NetworkId string `json:"network_id"`

	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`

	DeviceOwner *string `json:"device_owner,omitempty"`

	SecurityGroups *[]string `json:"security_groups,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	ExtraDhcpOpts *[]ExtraDhcpOpt `json:"extra_dhcp_opts,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o CreatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortOption struct{}"
	}

	return strings.Join([]string{"CreatePortOption", string(data)}, " ")
}
