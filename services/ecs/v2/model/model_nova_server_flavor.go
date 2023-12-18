package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaServerFlavor struct {
	Id *string `json:"id,omitempty"`

	Links *[]NovaLink `json:"links,omitempty"`

	Vcpus *int32 `json:"vcpus,omitempty"`

	Ram *int32 `json:"ram,omitempty"`

	Disk *int32 `json:"disk,omitempty"`

	Ephemeral *int32 `json:"ephemeral,omitempty"`

	Swap *int32 `json:"swap,omitempty"`

	OriginalName *string `json:"original_name,omitempty"`

	ExtraSpecs map[string]string `json:"extra_specs,omitempty"`
}

func (o NovaServerFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerFlavor struct{}"
	}

	return strings.Join([]string{"NovaServerFlavor", string(data)}, " ")
}
