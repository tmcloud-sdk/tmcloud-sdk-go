package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSubnetsByTagsRequest struct {
	Body *ListSubnetsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListSubnetsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsRequest", string(data)}, " ")
}
