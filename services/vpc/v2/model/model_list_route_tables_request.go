package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRouteTablesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *string `json:"id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o ListRouteTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRouteTablesRequest struct{}"
	}

	return strings.Join([]string{"ListRouteTablesRequest", string(data)}, " ")
}
