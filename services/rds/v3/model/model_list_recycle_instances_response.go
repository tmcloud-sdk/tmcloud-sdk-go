package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRecycleInstancesResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Instances      *[]RecycleInstsanceV3 `json:"instances,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRecycleInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesResponse", string(data)}, " ")
}
