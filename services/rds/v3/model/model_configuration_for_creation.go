package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ConfigurationForCreation struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Datastore *ParaGroupDatastore `json:"datastore"`

	Values map[string]string `json:"values,omitempty"`
}

func (o ConfigurationForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationForCreation struct{}"
	}

	return strings.Join([]string{"ConfigurationForCreation", string(data)}, " ")
}
