package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type JobRecords struct {
	JobName *string `json:"job_name,omitempty"`

	RecordType *JobRecordsRecordType `json:"record_type,omitempty"`

	RecordTime *string `json:"record_time,omitempty"`

	Request *string `json:"request,omitempty"`

	Response *string `json:"response,omitempty"`

	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`

	JobStatus *JobRecordsJobStatus `json:"job_status,omitempty"`
}

func (o JobRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRecords struct{}"
	}

	return strings.Join([]string{"JobRecords", string(data)}, " ")
}

type JobRecordsRecordType struct {
	value string
}

type JobRecordsRecordTypeEnum struct {
	API JobRecordsRecordType
	MEG JobRecordsRecordType
}

func GetJobRecordsRecordTypeEnum() JobRecordsRecordTypeEnum {
	return JobRecordsRecordTypeEnum{
		API: JobRecordsRecordType{
			value: "API",
		},
		MEG: JobRecordsRecordType{
			value: "MEG",
		},
	}
}

func (c JobRecordsRecordType) Value() string {
	return c.value
}

func (c JobRecordsRecordType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobRecordsRecordType) UnmarshalJSON(b []byte) error {
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

type JobRecordsJobStatus struct {
	value string
}

type JobRecordsJobStatusEnum struct {
	SUCCESS JobRecordsJobStatus
	FAIL    JobRecordsJobStatus
}

func GetJobRecordsJobStatusEnum() JobRecordsJobStatusEnum {
	return JobRecordsJobStatusEnum{
		SUCCESS: JobRecordsJobStatus{
			value: "SUCCESS",
		},
		FAIL: JobRecordsJobStatus{
			value: "FAIL",
		},
	}
}

func (c JobRecordsJobStatus) Value() string {
	return c.value
}

func (c JobRecordsJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobRecordsJobStatus) UnmarshalJSON(b []byte) error {
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
