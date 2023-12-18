package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesRequest struct {
	InstanceId *string `json:"instance_id,omitempty"`

	IncludeFailure *string `json:"include_failure,omitempty"`

	IncludeDelete *string `json:"include_delete,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Status *string `json:"status,omitempty"`

	NameEqual *string `json:"name_equal,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Ip *string `json:"ip,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
