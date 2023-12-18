package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateHotkeyScanTaskResponse struct {
	Id *string `json:"id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *CreateHotkeyScanTaskResponseStatus `json:"status,omitempty"`

	ScanType *CreateHotkeyScanTaskResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt *string `json:"finished_at,omitempty"`

	Num *int32 `json:"num,omitempty"`

	Keys           *[]HotkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateHotkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateHotkeyScanTaskResponse", string(data)}, " ")
}

type CreateHotkeyScanTaskResponseStatus struct {
	value string
}

type CreateHotkeyScanTaskResponseStatusEnum struct {
	WAITING CreateHotkeyScanTaskResponseStatus
	RUNNING CreateHotkeyScanTaskResponseStatus
	SUCCESS CreateHotkeyScanTaskResponseStatus
	FAILED  CreateHotkeyScanTaskResponseStatus
}

func GetCreateHotkeyScanTaskResponseStatusEnum() CreateHotkeyScanTaskResponseStatusEnum {
	return CreateHotkeyScanTaskResponseStatusEnum{
		WAITING: CreateHotkeyScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: CreateHotkeyScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: CreateHotkeyScanTaskResponseStatus{
			value: "success",
		},
		FAILED: CreateHotkeyScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c CreateHotkeyScanTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateHotkeyScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHotkeyScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateHotkeyScanTaskResponseScanType struct {
	value string
}

type CreateHotkeyScanTaskResponseScanTypeEnum struct {
	MANUAL CreateHotkeyScanTaskResponseScanType
	AUTO   CreateHotkeyScanTaskResponseScanType
}

func GetCreateHotkeyScanTaskResponseScanTypeEnum() CreateHotkeyScanTaskResponseScanTypeEnum {
	return CreateHotkeyScanTaskResponseScanTypeEnum{
		MANUAL: CreateHotkeyScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: CreateHotkeyScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c CreateHotkeyScanTaskResponseScanType) Value() string {
	return c.value
}

func (c CreateHotkeyScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHotkeyScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
