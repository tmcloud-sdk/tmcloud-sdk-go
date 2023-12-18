package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PauseScalingGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupResponse", string(data)}, " ")
}
