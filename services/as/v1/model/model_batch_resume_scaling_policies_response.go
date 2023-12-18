package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchResumeScalingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchResumeScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResumeScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchResumeScalingPoliciesResponse", string(data)}, " ")
}
