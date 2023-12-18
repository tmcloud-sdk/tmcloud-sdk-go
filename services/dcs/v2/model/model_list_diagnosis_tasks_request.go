package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListDiagnosisTasksRequest struct {
	InstanceId string `json:"instance_id"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDiagnosisTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisTasksRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnosisTasksRequest", string(data)}, " ")
}
