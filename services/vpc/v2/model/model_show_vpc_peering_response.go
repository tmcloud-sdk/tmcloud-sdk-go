package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcPeeringResponse struct {
	Peering        *VpcPeering `json:"peering,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVpcPeeringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcPeeringResponse", string(data)}, " ")
}
