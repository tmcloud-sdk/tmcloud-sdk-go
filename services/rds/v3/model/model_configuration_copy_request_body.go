package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationCopyRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`
}

func (o ConfigurationCopyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationCopyRequestBody struct{}"
	}

	return strings.Join([]string{"ConfigurationCopyRequestBody", string(data)}, " ")
}
