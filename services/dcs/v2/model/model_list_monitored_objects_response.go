package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMonitoredObjectsResponse struct {
	Router *[]string `json:"router,omitempty"`

	Children *[]DimChild `json:"children,omitempty"`

	Instances *[]InstancesMonitoredObject `json:"instances,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMonitoredObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsResponse", string(data)}, " ")
}
