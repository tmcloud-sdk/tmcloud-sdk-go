package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LvmConfig struct {
	LvType string `json:"lvType"`

	Path *string `json:"path,omitempty"`
}

func (o LvmConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LvmConfig struct{}"
	}

	return strings.Join([]string{"LvmConfig", string(data)}, " ")
}
