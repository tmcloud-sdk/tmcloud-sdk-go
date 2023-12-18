package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateipsRequest struct {
	SubnetId string `json:"subnet_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListPrivateipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateipsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateipsRequest", string(data)}, " ")
}
