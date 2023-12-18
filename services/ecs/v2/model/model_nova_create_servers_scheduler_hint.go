package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaCreateServersSchedulerHint struct {
	Group *string `json:"group,omitempty"`

	DifferentHost *[]string `json:"different_host,omitempty"`

	SameHost *[]string `json:"same_host,omitempty"`

	Cidr *string `json:"cidr,omitempty"`

	BuildNearHostIp *string `json:"build_near_host_ip,omitempty"`

	Tenancy *string `json:"tenancy,omitempty"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`
}

func (o NovaCreateServersSchedulerHint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersSchedulerHint struct{}"
	}

	return strings.Join([]string{"NovaCreateServersSchedulerHint", string(data)}, " ")
}
