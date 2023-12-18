package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRequest struct {
	Body *CreateVpcRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcRequest", string(data)}, " ")
}
