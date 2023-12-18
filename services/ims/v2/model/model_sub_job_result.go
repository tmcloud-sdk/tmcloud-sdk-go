package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SubJobResult struct {
	Status *SubJobResultStatus `json:"status,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`

	Entities *SubJobEntities `json:"entities,omitempty"`
}

func (o SubJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobResult struct{}"
	}

	return strings.Join([]string{"SubJobResult", string(data)}, " ")
}

type SubJobResultStatus struct {
	value string
}

type SubJobResultStatusEnum struct {
	SUCCESS SubJobResultStatus
	FAIL    SubJobResultStatus
	RUNNING SubJobResultStatus
	INIT    SubJobResultStatus
}

func GetSubJobResultStatusEnum() SubJobResultStatusEnum {
	return SubJobResultStatusEnum{
		SUCCESS: SubJobResultStatus{
			value: "SUCCESS",
		},
		FAIL: SubJobResultStatus{
			value: "FAIL",
		},
		RUNNING: SubJobResultStatus{
			value: "RUNNING",
		},
		INIT: SubJobResultStatus{
			value: "INIT",
		},
	}
}

func (c SubJobResultStatus) Value() string {
	return c.value
}

func (c SubJobResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobResultStatus) UnmarshalJSON(b []byte) error {
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
