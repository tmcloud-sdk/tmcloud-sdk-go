package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListDiagnosisTasksResponse struct {
	DiagnosisReportList *[]DiagnosisReportInfo `json:"diagnosis_report_list,omitempty"`

	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDiagnosisTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisTasksResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnosisTasksResponse", string(data)}, " ")
}
