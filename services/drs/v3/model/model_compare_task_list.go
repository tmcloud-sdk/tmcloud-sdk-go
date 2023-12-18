package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CompareTaskList struct {
	CompareTaskId string `json:"compare_task_id"`

	CompareType string `json:"compare_type"`

	CompareTaskStatus CompareTaskListCompareTaskStatus `json:"compare_task_status"`

	CreateTime string `json:"create_time"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o CompareTaskList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskList struct{}"
	}

	return strings.Join([]string{"CompareTaskList", string(data)}, " ")
}

type CompareTaskListCompareTaskStatus struct {
	value string
}

type CompareTaskListCompareTaskStatusEnum struct {
	RUNNING             CompareTaskListCompareTaskStatus
	WAITING_FOR_RUNNING CompareTaskListCompareTaskStatus
	SUCCESSFUL          CompareTaskListCompareTaskStatus
	FAILED              CompareTaskListCompareTaskStatus
	CANCELLED           CompareTaskListCompareTaskStatus
	TIMEOUT_INTERRUPT   CompareTaskListCompareTaskStatus
	FULL_DOING          CompareTaskListCompareTaskStatus
	INCRE_DOING         CompareTaskListCompareTaskStatus
}

func GetCompareTaskListCompareTaskStatusEnum() CompareTaskListCompareTaskStatusEnum {
	return CompareTaskListCompareTaskStatusEnum{
		RUNNING: CompareTaskListCompareTaskStatus{
			value: "RUNNING",
		},
		WAITING_FOR_RUNNING: CompareTaskListCompareTaskStatus{
			value: "WAITING_FOR_RUNNING",
		},
		SUCCESSFUL: CompareTaskListCompareTaskStatus{
			value: "SUCCESSFUL",
		},
		FAILED: CompareTaskListCompareTaskStatus{
			value: "FAILED",
		},
		CANCELLED: CompareTaskListCompareTaskStatus{
			value: "CANCELLED",
		},
		TIMEOUT_INTERRUPT: CompareTaskListCompareTaskStatus{
			value: "TIMEOUT_INTERRUPT",
		},
		FULL_DOING: CompareTaskListCompareTaskStatus{
			value: "FULL_DOING",
		},
		INCRE_DOING: CompareTaskListCompareTaskStatus{
			value: "INCRE_DOING",
		},
	}
}

func (c CompareTaskListCompareTaskStatus) Value() string {
	return c.value
}

func (c CompareTaskListCompareTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompareTaskListCompareTaskStatus) UnmarshalJSON(b []byte) error {
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
