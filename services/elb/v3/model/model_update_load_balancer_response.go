package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateLoadBalancerResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerResponse", string(data)}, " ")
}
