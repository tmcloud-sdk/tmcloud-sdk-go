package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateDiagnosisTaskRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateDiagnosisTaskBody `json:"body,omitempty"`
}

func (o CreateDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskRequest", string(data)}, " ")
}
