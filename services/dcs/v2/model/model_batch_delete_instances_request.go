package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstancesRequest struct {
	AllFailure *bool `json:"all_failure,omitempty"`

	Body *BatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstancesRequest", string(data)}, " ")
}
