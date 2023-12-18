package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowNatGatewayRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`
}

func (o ShowNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayRequest", string(data)}, " ")
}
