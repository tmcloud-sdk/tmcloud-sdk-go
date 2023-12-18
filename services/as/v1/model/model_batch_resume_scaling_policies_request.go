package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchResumeScalingPoliciesRequest struct {
	Body *BatchResumeScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchResumeScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResumeScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchResumeScalingPoliciesRequest", string(data)}, " ")
}
