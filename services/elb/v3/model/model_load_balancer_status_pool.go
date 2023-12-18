package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusPool struct {
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	Name *string `json:"name,omitempty"`

	Healthmonitor *LoadBalancerStatusHealthMonitor `json:"healthmonitor,omitempty"`

	Members *[]LoadBalancerStatusMember `json:"members,omitempty"`

	Id *string `json:"id,omitempty"`

	OperatingStatus *string `json:"operating_status,omitempty"`
}

func (o LoadBalancerStatusPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPool struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPool", string(data)}, " ")
}
