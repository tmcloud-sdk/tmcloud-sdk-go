package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SetOnlineMigrationTaskResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status         *SetOnlineMigrationTaskResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o SetOnlineMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskResponse", string(data)}, " ")
}

type SetOnlineMigrationTaskResponseStatus struct {
	value string
}

type SetOnlineMigrationTaskResponseStatusEnum struct {
	SUCCESS    SetOnlineMigrationTaskResponseStatus
	FAILED     SetOnlineMigrationTaskResponseStatus
	MIGRATING  SetOnlineMigrationTaskResponseStatus
	TERMINATED SetOnlineMigrationTaskResponseStatus
}

func GetSetOnlineMigrationTaskResponseStatusEnum() SetOnlineMigrationTaskResponseStatusEnum {
	return SetOnlineMigrationTaskResponseStatusEnum{
		SUCCESS: SetOnlineMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: SetOnlineMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: SetOnlineMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: SetOnlineMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c SetOnlineMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c SetOnlineMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
