package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusResult struct {
	Loadbalancer *LoadBalancerStatus `json:"loadbalancer"`
}

func (o LoadBalancerStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusResult struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusResult", string(data)}, " ")
}
