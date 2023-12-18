package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateAutoExpireScanTaskResponse struct {
	Id *string `json:"id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *CreateAutoExpireScanTaskResponseStatus `json:"status,omitempty"`

	ScanType *CreateAutoExpireScanTaskResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt     *string `json:"finished_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAutoExpireScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoExpireScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAutoExpireScanTaskResponse", string(data)}, " ")
}

type CreateAutoExpireScanTaskResponseStatus struct {
	value string
}

type CreateAutoExpireScanTaskResponseStatusEnum struct {
	WAITING CreateAutoExpireScanTaskResponseStatus
	RUNNING CreateAutoExpireScanTaskResponseStatus
	SUCCESS CreateAutoExpireScanTaskResponseStatus
	FAILED  CreateAutoExpireScanTaskResponseStatus
}

func GetCreateAutoExpireScanTaskResponseStatusEnum() CreateAutoExpireScanTaskResponseStatusEnum {
	return CreateAutoExpireScanTaskResponseStatusEnum{
		WAITING: CreateAutoExpireScanTaskResponseStatus{
			value: "waiting",
		},
		RUNNING: CreateAutoExpireScanTaskResponseStatus{
			value: "running",
		},
		SUCCESS: CreateAutoExpireScanTaskResponseStatus{
			value: "success",
		},
		FAILED: CreateAutoExpireScanTaskResponseStatus{
			value: "failed",
		},
	}
}

func (c CreateAutoExpireScanTaskResponseStatus) Value() string {
	return c.value
}

func (c CreateAutoExpireScanTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAutoExpireScanTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAutoExpireScanTaskResponseScanType struct {
	value string
}

type CreateAutoExpireScanTaskResponseScanTypeEnum struct {
	MANUAL CreateAutoExpireScanTaskResponseScanType
	AUTO   CreateAutoExpireScanTaskResponseScanType
}

func GetCreateAutoExpireScanTaskResponseScanTypeEnum() CreateAutoExpireScanTaskResponseScanTypeEnum {
	return CreateAutoExpireScanTaskResponseScanTypeEnum{
		MANUAL: CreateAutoExpireScanTaskResponseScanType{
			value: "manual",
		},
		AUTO: CreateAutoExpireScanTaskResponseScanType{
			value: "auto",
		},
	}
}

func (c CreateAutoExpireScanTaskResponseScanType) Value() string {
	return c.value
}

func (c CreateAutoExpireScanTaskResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAutoExpireScanTaskResponseScanType) UnmarshalJSON(b []byte) error {
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
