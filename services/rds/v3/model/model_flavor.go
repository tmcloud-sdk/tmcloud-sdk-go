package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Flavor struct {
	Id string `json:"id"`

	Vcpus string `json:"vcpus"`

	Ram int32 `json:"ram"`

	SpecCode string `json:"spec_code"`

	InstanceMode string `json:"instance_mode"`

	AzStatus map[string]string `json:"az_status"`

	AzDesc map[string]string `json:"az_desc"`

	VersionName []string `json:"version_name"`

	GroupType string `json:"group_type"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
