package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Resources struct {
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *string `json:"resource_detail,omitempty"`

	Tags *[]ResourceTags `json:"tags,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
