package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GetTaskDetailListRspJobs struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Status GetTaskDetailListRspJobsStatus `json:"status"`

	Created string `json:"created"`

	Ended *string `json:"ended,omitempty"`

	Process *string `json:"process,omitempty"`

	TaskDetail *string `json:"task_detail,omitempty"`

	Instance *GetTaskDetailListRspJobsInstance `json:"instance"`

	Entities *interface{} `json:"entities,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`
}

func (o GetTaskDetailListRspJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTaskDetailListRspJobs struct{}"
	}

	return strings.Join([]string{"GetTaskDetailListRspJobs", string(data)}, " ")
}

type GetTaskDetailListRspJobsStatus struct {
	value string
}

type GetTaskDetailListRspJobsStatusEnum struct {
	RUNNING   GetTaskDetailListRspJobsStatus
	COMPLETED GetTaskDetailListRspJobsStatus
	FAILED    GetTaskDetailListRspJobsStatus
}

func GetGetTaskDetailListRspJobsStatusEnum() GetTaskDetailListRspJobsStatusEnum {
	return GetTaskDetailListRspJobsStatusEnum{
		RUNNING: GetTaskDetailListRspJobsStatus{
			value: "Running",
		},
		COMPLETED: GetTaskDetailListRspJobsStatus{
			value: "Completed",
		},
		FAILED: GetTaskDetailListRspJobsStatus{
			value: "Failed",
		},
	}
}

func (c GetTaskDetailListRspJobsStatus) Value() string {
	return c.value
}

func (c GetTaskDetailListRspJobsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetTaskDetailListRspJobsStatus) UnmarshalJSON(b []byte) error {
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
