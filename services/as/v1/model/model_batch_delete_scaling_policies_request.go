package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteScalingPoliciesRequest struct {
	Body *BatchDeleteScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchDeleteScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingPoliciesRequest", string(data)}, " ")
}
