package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUnprotectScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUnprotectScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnprotectScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchUnprotectScalingInstancesResponse", string(data)}, " ")
}
