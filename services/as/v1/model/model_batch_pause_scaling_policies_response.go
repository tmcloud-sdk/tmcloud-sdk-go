package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchPauseScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchPauseScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesResponse", string(data)}, " ")
}
