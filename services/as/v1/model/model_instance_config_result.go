package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceConfigResult struct {
	FlavorRef *string `json:"flavorRef,omitempty"`

	ImageRef *string `json:"imageRef,omitempty"`

	Disk *[]DiskResult `json:"disk,omitempty"`

	KeyName *string `json:"key_name,omitempty"`

	KeyFingerprint *string `json:"key_fingerprint,omitempty"`

	InstanceName *string `json:"instance_name,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	AdminPass *string `json:"adminPass,omitempty"`

	Personality *[]PersonalityResult `json:"personality,omitempty"`

	PublicIp *PublicipResult `json:"public_ip,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Metadata *VmMetaData `json:"metadata,omitempty"`

	SecurityGroups *[]SecurityGroups `json:"security_groups,omitempty"`

	ServerGroupId *string `json:"server_group_id,omitempty"`

	Tenancy *string `json:"tenancy,omitempty"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	MarketType *string `json:"market_type,omitempty"`

	MultiFlavorPriorityPolicy *string `json:"multi_flavor_priority_policy,omitempty"`
}

func (o InstanceConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfigResult struct{}"
	}

	return strings.Join([]string{"InstanceConfigResult", string(data)}, " ")
}
