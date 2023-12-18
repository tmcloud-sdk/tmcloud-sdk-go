package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SlowlogForLtsRequest struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Type *SlowlogForLtsRequestType `json:"type,omitempty"`

	LineNum *string `json:"line_num,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SearchType *SlowlogForLtsRequestSearchType `json:"search_type,omitempty"`

	Database *string `json:"database,omitempty"`
}

func (o SlowlogForLtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogForLtsRequest struct{}"
	}

	return strings.Join([]string{"SlowlogForLtsRequest", string(data)}, " ")
}

type SlowlogForLtsRequestType struct {
	value string
}

type SlowlogForLtsRequestTypeEnum struct {
	INSERT SlowlogForLtsRequestType
	UPDATE SlowlogForLtsRequestType
	SELECT SlowlogForLtsRequestType
	DELETE SlowlogForLtsRequestType
	CREATE SlowlogForLtsRequestType
}

func GetSlowlogForLtsRequestTypeEnum() SlowlogForLtsRequestTypeEnum {
	return SlowlogForLtsRequestTypeEnum{
		INSERT: SlowlogForLtsRequestType{
			value: "INSERT",
		},
		UPDATE: SlowlogForLtsRequestType{
			value: "UPDATE",
		},
		SELECT: SlowlogForLtsRequestType{
			value: "SELECT",
		},
		DELETE: SlowlogForLtsRequestType{
			value: "DELETE",
		},
		CREATE: SlowlogForLtsRequestType{
			value: "CREATE",
		},
	}
}

func (c SlowlogForLtsRequestType) Value() string {
	return c.value
}

func (c SlowlogForLtsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlowlogForLtsRequestType) UnmarshalJSON(b []byte) error {
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

type SlowlogForLtsRequestSearchType struct {
	value string
}

type SlowlogForLtsRequestSearchTypeEnum struct {
	FORWARDS  SlowlogForLtsRequestSearchType
	BACKWARDS SlowlogForLtsRequestSearchType
}

func GetSlowlogForLtsRequestSearchTypeEnum() SlowlogForLtsRequestSearchTypeEnum {
	return SlowlogForLtsRequestSearchTypeEnum{
		FORWARDS: SlowlogForLtsRequestSearchType{
			value: "forwards",
		},
		BACKWARDS: SlowlogForLtsRequestSearchType{
			value: "backwards",
		},
	}
}

func (c SlowlogForLtsRequestSearchType) Value() string {
	return c.value
}

func (c SlowlogForLtsRequestSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlowlogForLtsRequestSearchType) UnmarshalJSON(b []byte) error {
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
