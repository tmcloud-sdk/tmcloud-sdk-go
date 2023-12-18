package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListHotKeyScanTasksResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Records        *[]RecordsResponse `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListHotKeyScanTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotKeyScanTasksResponse struct{}"
	}

	return strings.Join([]string{"ListHotKeyScanTasksResponse", string(data)}, " ")
}
