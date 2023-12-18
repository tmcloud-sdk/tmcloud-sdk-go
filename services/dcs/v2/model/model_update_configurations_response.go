package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateConfigurationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationsResponse", string(data)}, " ")
}
