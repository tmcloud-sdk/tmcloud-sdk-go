package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AcceptVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o AcceptVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"AcceptVpcPeeringRequest", string(data)}, " ")
}
