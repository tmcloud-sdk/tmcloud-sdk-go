package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayResponse", string(data)}, " ")
}
