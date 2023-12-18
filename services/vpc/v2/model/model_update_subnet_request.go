package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubnetRequest struct {
	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	Body *UpdateSubnetRequestBody `json:"body,omitempty"`
}

func (o UpdateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetRequest", string(data)}, " ")
}
