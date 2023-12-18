package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type RecordsResponse struct {
	Id *string `json:"id,omitempty"`

	Status *RecordsResponseStatus `json:"status,omitempty"`

	ScanType *RecordsResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt *string `json:"finished_at,omitempty"`
}

func (o RecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordsResponse struct{}"
	}

	return strings.Join([]string{"RecordsResponse", string(data)}, " ")
}

type RecordsResponseStatus struct {
	value string
}

type RecordsResponseStatusEnum struct {
	WAITING RecordsResponseStatus
	RUNNING RecordsResponseStatus
	SUCCESS RecordsResponseStatus
	FAILED  RecordsResponseStatus
}

func GetRecordsResponseStatusEnum() RecordsResponseStatusEnum {
	return RecordsResponseStatusEnum{
		WAITING: RecordsResponseStatus{
			value: "waiting",
		},
		RUNNING: RecordsResponseStatus{
			value: "running",
		},
		SUCCESS: RecordsResponseStatus{
			value: "success",
		},
		FAILED: RecordsResponseStatus{
			value: "failed",
		},
	}
}

func (c RecordsResponseStatus) Value() string {
	return c.value
}

func (c RecordsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordsResponseStatus) UnmarshalJSON(b []byte) error {
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

type RecordsResponseScanType struct {
	value string
}

type RecordsResponseScanTypeEnum struct {
	MANUAL RecordsResponseScanType
	AUTO   RecordsResponseScanType
}

func GetRecordsResponseScanTypeEnum() RecordsResponseScanTypeEnum {
	return RecordsResponseScanTypeEnum{
		MANUAL: RecordsResponseScanType{
			value: "manual",
		},
		AUTO: RecordsResponseScanType{
			value: "auto",
		},
	}
}

func (c RecordsResponseScanType) Value() string {
	return c.value
}

func (c RecordsResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordsResponseScanType) UnmarshalJSON(b []byte) error {
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
