package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StopMigrationTaskResult struct {
	Result *StopMigrationTaskResultResult `json:"result,omitempty"`

	TaskId *string `json:"task_id,omitempty"`
}

func (o StopMigrationTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskResult struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskResult", string(data)}, " ")
}

type StopMigrationTaskResultResult struct {
	value string
}

type StopMigrationTaskResultResultEnum struct {
	SUCCESS StopMigrationTaskResultResult
	FAILED  StopMigrationTaskResultResult
}

func GetStopMigrationTaskResultResultEnum() StopMigrationTaskResultResultEnum {
	return StopMigrationTaskResultResultEnum{
		SUCCESS: StopMigrationTaskResultResult{
			value: "success",
		},
		FAILED: StopMigrationTaskResultResult{
			value: "failed",
		},
	}
}

func (c StopMigrationTaskResultResult) Value() string {
	return c.value
}

func (c StopMigrationTaskResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResultResult) UnmarshalJSON(b []byte) error {
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
