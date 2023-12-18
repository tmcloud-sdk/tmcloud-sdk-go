package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type VolumeType struct {
	Id string `json:"id"`

	Name string `json:"name"`

	ExtraSpecs *VolumeTypeExtraSpecs `json:"extra_specs,omitempty"`

	Description *string `json:"description,omitempty"`

	QosSpecsId *string `json:"qos_specs_id,omitempty"`

	IsPublic *bool `json:"is_public,omitempty"`
}

func (o VolumeType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeType struct{}"
	}

	return strings.Join([]string{"VolumeType", string(data)}, " ")
}
