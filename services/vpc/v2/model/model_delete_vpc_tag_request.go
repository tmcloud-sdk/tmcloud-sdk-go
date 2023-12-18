package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcTagRequest struct {
	VpcId string `json:"vpc_id"`

	Key string `json:"key"`
}

func (o DeleteVpcTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcTagRequest", string(data)}, " ")
}
