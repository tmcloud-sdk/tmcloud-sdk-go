package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListHotKeyScanTasksRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Status *ListHotKeyScanTasksRequestStatus `json:"status,omitempty"`
}

func (o ListHotKeyScanTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotKeyScanTasksRequest struct{}"
	}

	return strings.Join([]string{"ListHotKeyScanTasksRequest", string(data)}, " ")
}

type ListHotKeyScanTasksRequestStatus struct {
	value string
}

type ListHotKeyScanTasksRequestStatusEnum struct {
	WAITING ListHotKeyScanTasksRequestStatus
	RUNNING ListHotKeyScanTasksRequestStatus
	SUCCESS ListHotKeyScanTasksRequestStatus
	FAILED  ListHotKeyScanTasksRequestStatus
}

func GetListHotKeyScanTasksRequestStatusEnum() ListHotKeyScanTasksRequestStatusEnum {
	return ListHotKeyScanTasksRequestStatusEnum{
		WAITING: ListHotKeyScanTasksRequestStatus{
			value: "waiting",
		},
		RUNNING: ListHotKeyScanTasksRequestStatus{
			value: "running",
		},
		SUCCESS: ListHotKeyScanTasksRequestStatus{
			value: "success",
		},
		FAILED: ListHotKeyScanTasksRequestStatus{
			value: "failed",
		},
	}
}

func (c ListHotKeyScanTasksRequestStatus) Value() string {
	return c.value
}

func (c ListHotKeyScanTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeyScanTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
