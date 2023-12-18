package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusHealthMonitor struct {
	Type *string `json:"type,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
}

func (o LoadBalancerStatusHealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusHealthMonitor struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusHealthMonitor", string(data)}, " ")
}
