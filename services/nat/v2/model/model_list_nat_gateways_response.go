package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListNatGatewaysResponse struct {
	NatGateways    *[]NatGatewayResponseBody `json:"nat_gateways,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListNatGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaysResponse", string(data)}, " ")
}
