package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationForUpdate struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Values map[string]string `json:"values,omitempty"`
}

func (o ConfigurationForUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationForUpdate struct{}"
	}

	return strings.Join([]string{"ConfigurationForUpdate", string(data)}, " ")
}
