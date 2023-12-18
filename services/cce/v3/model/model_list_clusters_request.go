package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListClustersRequest struct {
	ErrorStatus *string `json:"errorStatus,omitempty"`

	Detail *string `json:"detail,omitempty"`

	Status *ListClustersRequestStatus `json:"status,omitempty"`

	Type *ListClustersRequestType `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}

type ListClustersRequestStatus struct {
	value string
}

type ListClustersRequestStatusEnum struct {
	AVAILABLE       ListClustersRequestStatus
	UNAVAILABLE     ListClustersRequestStatus
	SCALING_UP      ListClustersRequestStatus
	SCALING_DOWN    ListClustersRequestStatus
	CREATING        ListClustersRequestStatus
	DELETING        ListClustersRequestStatus
	UPGRADING       ListClustersRequestStatus
	RESIZING        ListClustersRequestStatus
	ROLLING_BACK    ListClustersRequestStatus
	ROLLBACK_FAILED ListClustersRequestStatus
	HIBERNATING     ListClustersRequestStatus
	HIBERNATION     ListClustersRequestStatus
	AWAKING         ListClustersRequestStatus
	EMPTY           ListClustersRequestStatus
}

func GetListClustersRequestStatusEnum() ListClustersRequestStatusEnum {
	return ListClustersRequestStatusEnum{
		AVAILABLE: ListClustersRequestStatus{
			value: "Available",
		},
		UNAVAILABLE: ListClustersRequestStatus{
			value: "Unavailable",
		},
		SCALING_UP: ListClustersRequestStatus{
			value: "ScalingUp",
		},
		SCALING_DOWN: ListClustersRequestStatus{
			value: "ScalingDown",
		},
		CREATING: ListClustersRequestStatus{
			value: "Creating",
		},
		DELETING: ListClustersRequestStatus{
			value: "Deleting",
		},
		UPGRADING: ListClustersRequestStatus{
			value: "Upgrading",
		},
		RESIZING: ListClustersRequestStatus{
			value: "Resizing",
		},
		ROLLING_BACK: ListClustersRequestStatus{
			value: "RollingBack",
		},
		ROLLBACK_FAILED: ListClustersRequestStatus{
			value: "RollbackFailed",
		},
		HIBERNATING: ListClustersRequestStatus{
			value: "Hibernating",
		},
		HIBERNATION: ListClustersRequestStatus{
			value: "Hibernation",
		},
		AWAKING: ListClustersRequestStatus{
			value: "Awaking",
		},
		EMPTY: ListClustersRequestStatus{
			value: "Empty",
		},
	}
}

func (c ListClustersRequestStatus) Value() string {
	return c.value
}

func (c ListClustersRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListClustersRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListClustersRequestType struct {
	value string
}

type ListClustersRequestTypeEnum struct {
	VIRTUAL_MACHINE ListClustersRequestType
	ARM64           ListClustersRequestType
}

func GetListClustersRequestTypeEnum() ListClustersRequestTypeEnum {
	return ListClustersRequestTypeEnum{
		VIRTUAL_MACHINE: ListClustersRequestType{
			value: "VirtualMachine",
		},
		ARM64: ListClustersRequestType{
			value: "ARM64",
		},
	}
}

func (c ListClustersRequestType) Value() string {
	return c.value
}

func (c ListClustersRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListClustersRequestType) UnmarshalJSON(b []byte) error {
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
