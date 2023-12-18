package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerLimits struct {
	MaxImageMeta int32 `json:"maxImageMeta"`

	MaxPersonality int32 `json:"maxPersonality"`

	MaxPersonalitySize int32 `json:"maxPersonalitySize"`

	MaxSecurityGroupRules int32 `json:"maxSecurityGroupRules"`

	MaxSecurityGroups int32 `json:"maxSecurityGroups"`

	MaxServerGroupMembers int32 `json:"maxServerGroupMembers"`

	MaxServerGroups int32 `json:"maxServerGroups"`

	MaxServerMeta int32 `json:"maxServerMeta"`

	MaxTotalCores int32 `json:"maxTotalCores"`

	MaxTotalFloatingIps int32 `json:"maxTotalFloatingIps"`

	MaxTotalInstances int32 `json:"maxTotalInstances"`

	MaxTotalKeypairs int32 `json:"maxTotalKeypairs"`

	MaxTotalRAMSize int32 `json:"maxTotalRAMSize"`

	TotalCoresUsed int32 `json:"totalCoresUsed"`

	TotalFloatingIpsUsed int32 `json:"totalFloatingIpsUsed"`

	TotalInstancesUsed int32 `json:"totalInstancesUsed"`

	TotalRAMUsed int32 `json:"totalRAMUsed"`

	TotalSecurityGroupsUsed int32 `json:"totalSecurityGroupsUsed"`

	TotalServerGroupsUsed int32 `json:"totalServerGroupsUsed"`

	MaxTotalSpotInstances *int32 `json:"maxTotalSpotInstances,omitempty"`

	MaxTotalSpotCores *int32 `json:"maxTotalSpotCores,omitempty"`

	MaxTotalSpotRAMSize *int32 `json:"maxTotalSpotRAMSize,omitempty"`

	TotalSpotInstancesUsed *int32 `json:"totalSpotInstancesUsed,omitempty"`

	TotalSpotCoresUsed *int32 `json:"totalSpotCoresUsed,omitempty"`

	TotalSpotRAMUsed *int32 `json:"totalSpotRAMUsed,omitempty"`

	LimitByFlavor *[]ProjectFlavorLimit `json:"limit_by_flavor,omitempty"`
}

func (o ServerLimits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerLimits struct{}"
	}

	return strings.Join([]string{"ServerLimits", string(data)}, " ")
}
