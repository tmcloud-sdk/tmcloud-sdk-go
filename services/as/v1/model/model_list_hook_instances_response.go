package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListHookInstancesResponse struct {
	InstanceHangingInfo *[]InstanceHangingInfos `json:"instance_hanging_info,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o ListHookInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHookInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListHookInstancesResponse", string(data)}, " ")
}
