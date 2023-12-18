package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListBigkeyScanTasksResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Records        *[]RecordsResponse `json:"records,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBigkeyScanTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigkeyScanTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBigkeyScanTasksResponse", string(data)}, " ")
}
