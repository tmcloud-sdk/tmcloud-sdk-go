package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowScalingConfigResponse struct {
	ScalingConfiguration *ScalingConfiguration `json:"scaling_configuration,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowScalingConfigResponse", string(data)}, " ")
}
