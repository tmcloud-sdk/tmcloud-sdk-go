package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingConfigResponse struct {
	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o CreateScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigResponse", string(data)}, " ")
}
