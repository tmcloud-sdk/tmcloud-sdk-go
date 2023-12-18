package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeLoadbalancerChargeModeResponse struct {
	EipIdList *[]string `json:"eip_id_list,omitempty"`

	LoadbalancerIdList *[]string `json:"loadbalancer_id_list,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeLoadbalancerChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeResponse", string(data)}, " ")
}
