package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DiagnosisNodeReport struct {
	NodeIp string `json:"node_ip"`

	AzCode string `json:"az_code"`

	GroupName string `json:"group_name"`

	AbnormalSum int32 `json:"abnormal_sum"`

	FailedSum int32 `json:"failed_sum"`

	Role DiagnosisNodeReportRole `json:"role"`

	DiagnosisDimensionList []DiagnosisDimension `json:"diagnosis_dimension_list"`

	CommandTimeTakenList *CommandTimeTakenList `json:"command_time_taken_list"`
}

func (o DiagnosisNodeReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisNodeReport struct{}"
	}

	return strings.Join([]string{"DiagnosisNodeReport", string(data)}, " ")
}

type DiagnosisNodeReportRole struct {
	value string
}

type DiagnosisNodeReportRoleEnum struct {
	MASTER DiagnosisNodeReportRole
	SLAVE  DiagnosisNodeReportRole
}

func GetDiagnosisNodeReportRoleEnum() DiagnosisNodeReportRoleEnum {
	return DiagnosisNodeReportRoleEnum{
		MASTER: DiagnosisNodeReportRole{
			value: "master",
		},
		SLAVE: DiagnosisNodeReportRole{
			value: "slave",
		},
	}
}

func (c DiagnosisNodeReportRole) Value() string {
	return c.value
}

func (c DiagnosisNodeReportRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisNodeReportRole) UnmarshalJSON(b []byte) error {
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
