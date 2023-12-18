package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatus struct {
	Name string `json:"name"`

	ProvisioningStatus string `json:"provisioning_status"`

	Listeners []LoadBalancerStatusListener `json:"listeners"`

	Pools []LoadBalancerStatusPool `json:"pools"`

	Id string `json:"id"`

	OperatingStatus string `json:"operating_status"`
}

func (o LoadBalancerStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatus struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatus", string(data)}, " ")
}
