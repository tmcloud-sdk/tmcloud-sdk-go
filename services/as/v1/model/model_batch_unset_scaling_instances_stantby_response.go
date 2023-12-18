package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUnsetScalingInstancesStantbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUnsetScalingInstancesStantbyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnsetScalingInstancesStantbyResponse struct{}"
	}

	return strings.Join([]string{"BatchUnsetScalingInstancesStantbyResponse", string(data)}, " ")
}
