package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingTagInfosByTenantIdResponse struct {
	Tags           *[]TagsMultiValue `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListScalingTagInfosByTenantIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByTenantIdResponse struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByTenantIdResponse", string(data)}, " ")
}
