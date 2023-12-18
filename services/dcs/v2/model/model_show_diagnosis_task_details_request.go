package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDiagnosisTaskDetailsRequest struct {
	ReportId string `json:"report_id"`
}

func (o ShowDiagnosisTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskDetailsRequest", string(data)}, " ")
}
