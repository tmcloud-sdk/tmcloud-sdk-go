package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationResponse struct {
	Configuration  *UpdateConfigurationRspConfiguration `json:"configuration,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o UpdateConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationResponse", string(data)}, " ")
}
