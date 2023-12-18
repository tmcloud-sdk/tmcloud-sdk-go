package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayResponse struct {
	NatGateway     *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayResponse", string(data)}, " ")
}
