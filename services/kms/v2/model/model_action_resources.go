package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ActionResources struct {
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *KeyDetails `json:"resource_detail,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`

	Tags *[]TagItem `json:"tags,omitempty"`
}

func (o ActionResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionResources struct{}"
	}

	return strings.Join([]string{"ActionResources", string(data)}, " ")
}
