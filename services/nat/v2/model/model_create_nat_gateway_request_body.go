package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayRequestBody struct {
	NatGateway *CreateNatGatewayOption `json:"nat_gateway"`
}

func (o CreateNatGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayRequestBody", string(data)}, " ")
}
