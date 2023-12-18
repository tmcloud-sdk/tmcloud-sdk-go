package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DiagnosisReportInfo struct {
	ReportId string `json:"report_id"`

	Status DiagnosisReportInfoStatus `json:"status"`

	BeginTime string `json:"begin_time"`

	EndTime string `json:"end_time"`

	CreatedAt string `json:"created_at"`

	NodeNum int32 `json:"node_num"`

	AbnormalItemSum int32 `json:"abnormal_item_sum"`

	FailedItemSum int32 `json:"failed_item_sum"`
}

func (o DiagnosisReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisReportInfo struct{}"
	}

	return strings.Join([]string{"DiagnosisReportInfo", string(data)}, " ")
}

type DiagnosisReportInfoStatus struct {
	value string
}

type DiagnosisReportInfoStatusEnum struct {
	DIAGNOSING DiagnosisReportInfoStatus
	FINISHED   DiagnosisReportInfoStatus
}

func GetDiagnosisReportInfoStatusEnum() DiagnosisReportInfoStatusEnum {
	return DiagnosisReportInfoStatusEnum{
		DIAGNOSING: DiagnosisReportInfoStatus{
			value: "diagnosing",
		},
		FINISHED: DiagnosisReportInfoStatus{
			value: "finished",
		},
	}
}

func (c DiagnosisReportInfoStatus) Value() string {
	return c.value
}

func (c DiagnosisReportInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisReportInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
