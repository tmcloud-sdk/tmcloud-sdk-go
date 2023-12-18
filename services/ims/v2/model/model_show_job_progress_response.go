package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowJobProgressResponse struct {
	Status *ShowJobProgressResponseStatus `json:"status,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	JobType *string `json:"job_type,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`

	Entities       *JobProgressEntities `json:"entities,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowJobProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowJobProgressResponse", string(data)}, " ")
}

type ShowJobProgressResponseStatus struct {
	value string
}

type ShowJobProgressResponseStatusEnum struct {
	SUCCESS ShowJobProgressResponseStatus
	FAIL    ShowJobProgressResponseStatus
	RUNNING ShowJobProgressResponseStatus
	INIT    ShowJobProgressResponseStatus
}

func GetShowJobProgressResponseStatusEnum() ShowJobProgressResponseStatusEnum {
	return ShowJobProgressResponseStatusEnum{
		SUCCESS: ShowJobProgressResponseStatus{
			value: "SUCCESS",
		},
		FAIL: ShowJobProgressResponseStatus{
			value: "FAIL",
		},
		RUNNING: ShowJobProgressResponseStatus{
			value: "RUNNING",
		},
		INIT: ShowJobProgressResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobProgressResponseStatus) Value() string {
	return c.value
}

func (c ShowJobProgressResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobProgressResponseStatus) UnmarshalJSON(b []byte) error {
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
