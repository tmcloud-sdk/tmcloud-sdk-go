package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListLoadBalancersResponse struct {
	Loadbalancers *[]LoadBalancer `json:"loadbalancers,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLoadBalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersResponse struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersResponse", string(data)}, " ")
}
