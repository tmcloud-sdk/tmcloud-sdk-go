package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ExtensionRequest struct {
	DatabaseName string `json:"database_name"`

	ExtensionName string `json:"extension_name"`
}

func (o ExtensionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionRequest struct{}"
	}

	return strings.Join([]string{"ExtensionRequest", string(data)}, " ")
}
