package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowScalingV2PolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o ShowScalingV2PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingV2PolicyRequest", string(data)}, " ")
}
