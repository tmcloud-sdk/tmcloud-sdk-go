package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListSlowLogsNewRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	Offset *int64 `json:"offset,omitempty"`

	Limit *int64 `json:"limit,omitempty"`

	Type *ListSlowLogsNewRequestType `json:"type,omitempty"`
}

func (o ListSlowLogsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsNewRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogsNewRequest", string(data)}, " ")
}

type ListSlowLogsNewRequestType struct {
	value string
}

type ListSlowLogsNewRequestTypeEnum struct {
	INSERT ListSlowLogsNewRequestType
	UPDATE ListSlowLogsNewRequestType
	SELECT ListSlowLogsNewRequestType
	DELETE ListSlowLogsNewRequestType
	CREATE ListSlowLogsNewRequestType
}

func GetListSlowLogsNewRequestTypeEnum() ListSlowLogsNewRequestTypeEnum {
	return ListSlowLogsNewRequestTypeEnum{
		INSERT: ListSlowLogsNewRequestType{
			value: "INSERT",
		},
		UPDATE: ListSlowLogsNewRequestType{
			value: "UPDATE",
		},
		SELECT: ListSlowLogsNewRequestType{
			value: "SELECT",
		},
		DELETE: ListSlowLogsNewRequestType{
			value: "DELETE",
		},
		CREATE: ListSlowLogsNewRequestType{
			value: "CREATE",
		},
	}
}

func (c ListSlowLogsNewRequestType) Value() string {
	return c.value
}

func (c ListSlowLogsNewRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowLogsNewRequestType) UnmarshalJSON(b []byte) error {
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
