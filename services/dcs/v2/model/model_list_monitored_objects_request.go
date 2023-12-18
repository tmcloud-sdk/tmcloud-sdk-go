package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMonitoredObjectsRequest struct {
	DimName string `json:"dim_name"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMonitoredObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsRequest", string(data)}, " ")
}
