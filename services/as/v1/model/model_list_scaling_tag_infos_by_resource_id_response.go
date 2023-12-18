package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingTagInfosByResourceIdResponse struct {
	Tags *[]TagsSingleValue `json:"tags,omitempty"`

	SysTags        *[]TagsSingleValue `json:"sys_tags,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListScalingTagInfosByResourceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByResourceIdResponse struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByResourceIdResponse", string(data)}, " ")
}
