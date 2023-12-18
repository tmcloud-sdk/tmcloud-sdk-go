package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaListKeypairsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaListKeypairsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsRequest struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsRequest", string(data)}, " ")
}
