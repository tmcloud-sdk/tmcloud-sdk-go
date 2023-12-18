package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcTagsRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o ShowVpcTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcTagsRequest", string(data)}, " ")
}
