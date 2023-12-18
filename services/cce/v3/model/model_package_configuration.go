package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PackageConfiguration struct {
	Name *string `json:"name,omitempty"`

	Configurations *[]ConfigurationItem `json:"configurations,omitempty"`
}

func (o PackageConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageConfiguration struct{}"
	}

	return strings.Join([]string{"PackageConfiguration", string(data)}, " ")
}
