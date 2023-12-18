package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ErrorlogForLtsRequest struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	Level *ErrorlogForLtsRequestLevel `json:"level,omitempty"`

	LineNum *string `json:"line_num,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SearchType *string `json:"search_type,omitempty"`
}

func (o ErrorlogForLtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorlogForLtsRequest struct{}"
	}

	return strings.Join([]string{"ErrorlogForLtsRequest", string(data)}, " ")
}

type ErrorlogForLtsRequestLevel struct {
	value string
}

type ErrorlogForLtsRequestLevelEnum struct {
	ALL     ErrorlogForLtsRequestLevel
	INFO    ErrorlogForLtsRequestLevel
	LOG     ErrorlogForLtsRequestLevel
	WARNING ErrorlogForLtsRequestLevel
	ERROR   ErrorlogForLtsRequestLevel
	FATAL   ErrorlogForLtsRequestLevel
	PANIC   ErrorlogForLtsRequestLevel
	NOTE    ErrorlogForLtsRequestLevel
}

func GetErrorlogForLtsRequestLevelEnum() ErrorlogForLtsRequestLevelEnum {
	return ErrorlogForLtsRequestLevelEnum{
		ALL: ErrorlogForLtsRequestLevel{
			value: "ALL",
		},
		INFO: ErrorlogForLtsRequestLevel{
			value: "INFO",
		},
		LOG: ErrorlogForLtsRequestLevel{
			value: "LOG",
		},
		WARNING: ErrorlogForLtsRequestLevel{
			value: "WARNING",
		},
		ERROR: ErrorlogForLtsRequestLevel{
			value: "ERROR",
		},
		FATAL: ErrorlogForLtsRequestLevel{
			value: "FATAL",
		},
		PANIC: ErrorlogForLtsRequestLevel{
			value: "PANIC",
		},
		NOTE: ErrorlogForLtsRequestLevel{
			value: "NOTE",
		},
	}
}

func (c ErrorlogForLtsRequestLevel) Value() string {
	return c.value
}

func (c ErrorlogForLtsRequestLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ErrorlogForLtsRequestLevel) UnmarshalJSON(b []byte) error {
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
