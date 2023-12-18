package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CountPreoccupyIpNumRequest struct {
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumRequest struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}
