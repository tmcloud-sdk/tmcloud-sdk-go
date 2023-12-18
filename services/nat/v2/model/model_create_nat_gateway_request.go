package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayRequest struct {
	Body *CreateNatGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayRequest", string(data)}, " ")
}
