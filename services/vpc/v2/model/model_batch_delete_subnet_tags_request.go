package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteSubnetTagsRequest struct {
	SubnetId string `json:"subnet_id"`

	Body *BatchDeleteSubnetTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsRequest", string(data)}, " ")
}
