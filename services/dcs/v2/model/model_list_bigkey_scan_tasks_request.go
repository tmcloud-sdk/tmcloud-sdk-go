package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListBigkeyScanTasksRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Status *ListBigkeyScanTasksRequestStatus `json:"status,omitempty"`
}

func (o ListBigkeyScanTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigkeyScanTasksRequest struct{}"
	}

	return strings.Join([]string{"ListBigkeyScanTasksRequest", string(data)}, " ")
}

type ListBigkeyScanTasksRequestStatus struct {
	value string
}

type ListBigkeyScanTasksRequestStatusEnum struct {
	WAITING ListBigkeyScanTasksRequestStatus
	RUNNING ListBigkeyScanTasksRequestStatus
	SUCCESS ListBigkeyScanTasksRequestStatus
	FAILED  ListBigkeyScanTasksRequestStatus
}

func GetListBigkeyScanTasksRequestStatusEnum() ListBigkeyScanTasksRequestStatusEnum {
	return ListBigkeyScanTasksRequestStatusEnum{
		WAITING: ListBigkeyScanTasksRequestStatus{
			value: "waiting",
		},
		RUNNING: ListBigkeyScanTasksRequestStatus{
			value: "running",
		},
		SUCCESS: ListBigkeyScanTasksRequestStatus{
			value: "success",
		},
		FAILED: ListBigkeyScanTasksRequestStatus{
			value: "failed",
		},
	}
}

func (c ListBigkeyScanTasksRequestStatus) Value() string {
	return c.value
}

func (c ListBigkeyScanTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigkeyScanTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
