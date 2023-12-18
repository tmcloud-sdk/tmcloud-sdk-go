package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Resource struct {
	ResourceId string `json:"resource_id"`

	ResourceName *string `json:"resource_name,omitempty"`

	ResourceDetail *VolumeDetailForTag `json:"resource_detail"`

	Tags []map[string]string `json:"tags"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
