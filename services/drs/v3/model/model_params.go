package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Params struct {
	CompareResult *ParamsCompareResult `json:"compare_result,omitempty"`

	DataType *string `json:"data_type,omitempty"`

	Group *ParamsGroup `json:"group,omitempty"`

	Key *string `json:"key,omitempty"`

	NeedRestart *ParamsNeedRestart `json:"need_restart,omitempty"`

	SourceValue *string `json:"source_value,omitempty"`

	TargetValue *string `json:"target_value,omitempty"`

	ValueRange *string `json:"value_range,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o Params) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Params struct{}"
	}

	return strings.Join([]string{"Params", string(data)}, " ")
}

type ParamsCompareResult struct {
	value string
}

type ParamsCompareResultEnum struct {
	TRUE  ParamsCompareResult
	FALSE ParamsCompareResult
}

func GetParamsCompareResultEnum() ParamsCompareResultEnum {
	return ParamsCompareResultEnum{
		TRUE: ParamsCompareResult{
			value: "true",
		},
		FALSE: ParamsCompareResult{
			value: "false",
		},
	}
}

func (c ParamsCompareResult) Value() string {
	return c.value
}

func (c ParamsCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsCompareResult) UnmarshalJSON(b []byte) error {
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

type ParamsGroup struct {
	value string
}

type ParamsGroupEnum struct {
	COMMON      ParamsGroup
	PERFORMANCE ParamsGroup
}

func GetParamsGroupEnum() ParamsGroupEnum {
	return ParamsGroupEnum{
		COMMON: ParamsGroup{
			value: "common",
		},
		PERFORMANCE: ParamsGroup{
			value: "performance",
		},
	}
}

func (c ParamsGroup) Value() string {
	return c.value
}

func (c ParamsGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsGroup) UnmarshalJSON(b []byte) error {
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

type ParamsNeedRestart struct {
	value string
}

type ParamsNeedRestartEnum struct {
	TRUE  ParamsNeedRestart
	FALSE ParamsNeedRestart
}

func GetParamsNeedRestartEnum() ParamsNeedRestartEnum {
	return ParamsNeedRestartEnum{
		TRUE: ParamsNeedRestart{
			value: "true",
		},
		FALSE: ParamsNeedRestart{
			value: "false",
		},
	}
}

func (c ParamsNeedRestart) Value() string {
	return c.value
}

func (c ParamsNeedRestart) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsNeedRestart) UnmarshalJSON(b []byte) error {
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
