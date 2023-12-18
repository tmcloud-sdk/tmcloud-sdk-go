package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusListener struct {
	Name *string `json:"name,omitempty"`

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	Pools *[]LoadBalancerStatusPool `json:"pools,omitempty"`

	L7policies *[]LoadBalancerStatusPolicy `json:"l7policies,omitempty"`

	Id *string `json:"id,omitempty"`

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusListener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusListener struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusListener", string(data)}, " ")
}
