package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateBigkeyScanTaskResponse struct {
	Id *string `json:"id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *CreateBigkeyScanTaskResponseStatus `json:"status,omitempty"`

	ScanType *CreateBigkeyScanTaskResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt *string `json:"finished_at,omitempty"`

	Num *int32 `json:"num,omitempty"`

	Keys           *[]BigkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateBigkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBigkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateBigkeyScanTaskResponse", string(data)}, " ")
}

type CreateBigkeyScanTaskResponseStatus struct {
	value string
}

type CreateBigkeyScanTaskResponseStatusEnum struct {
	WAITING CreateBigkeyScanTaskResponseStatus
	RUNNING CreateBigkeyScanTaskResponseStatus
	SUCCESS CreateBigkeyScanTaskResponseStatus
	FAILED  CreateBigkeyScanTaskResponseStatus
}

func GetCreateBigkeyScanTaskResponseStatusEnum() CreateBigkeyScanTaskResponseStatusEnum {
	return CreateBigkeyScanTaskResponseStatusEnum{
		WAITING: CreateBigkeyScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: CreateBigkeyScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: CreateBigkeyScanTaskResponseStatus{
			value: "success",
		},
		FAILED: CreateBigkeyScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c CreateBigkeyScanTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateBigkeyScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBigkeyScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateBigkeyScanTaskResponseScanType struct {
	value string
}

type CreateBigkeyScanTaskResponseScanTypeEnum struct {
	MANUAL CreateBigkeyScanTaskResponseScanType
	AUTO   CreateBigkeyScanTaskResponseScanType
}

func GetCreateBigkeyScanTaskResponseScanTypeEnum() CreateBigkeyScanTaskResponseScanTypeEnum {
	return CreateBigkeyScanTaskResponseScanTypeEnum{
		MANUAL: CreateBigkeyScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: CreateBigkeyScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c CreateBigkeyScanTaskResponseScanType) Value() string {
	return c.value
}

func (c CreateBigkeyScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBigkeyScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
