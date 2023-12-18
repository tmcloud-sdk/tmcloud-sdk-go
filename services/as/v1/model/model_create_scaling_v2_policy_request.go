package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingV2PolicyRequest struct {
	Body *CreateScalingPolicyV2Option `json:"body,omitempty"`
}

func (o CreateScalingV2PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingV2PolicyRequest", string(data)}, " ")
}
