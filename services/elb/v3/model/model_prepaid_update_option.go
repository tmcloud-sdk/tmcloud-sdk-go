package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrepaidUpdateOption struct {
	AutoPay *bool `json:"auto_pay,omitempty"`

	ChangeMode *PrepaidUpdateOptionChangeMode `json:"change_mode,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	PeriodType *PrepaidUpdateOptionPeriodType `json:"period_type,omitempty"`
}

func (o PrepaidUpdateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidUpdateOption struct{}"
	}

	return strings.Join([]string{"PrepaidUpdateOption", string(data)}, " ")
}

type PrepaidUpdateOptionChangeMode struct {
	value string
}

type PrepaidUpdateOptionChangeModeEnum struct {
	IMMEDIATE PrepaidUpdateOptionChangeMode
	DELAY     PrepaidUpdateOptionChangeMode
}

func GetPrepaidUpdateOptionChangeModeEnum() PrepaidUpdateOptionChangeModeEnum {
	return PrepaidUpdateOptionChangeModeEnum{
		IMMEDIATE: PrepaidUpdateOptionChangeMode{
			value: "immediate",
		},
		DELAY: PrepaidUpdateOptionChangeMode{
			value: "delay",
		},
	}
}

func (c PrepaidUpdateOptionChangeMode) Value() string {
	return c.value
}

func (c PrepaidUpdateOptionChangeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidUpdateOptionChangeMode) UnmarshalJSON(b []byte) error {
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

type PrepaidUpdateOptionPeriodType struct {
	value string
}

type PrepaidUpdateOptionPeriodTypeEnum struct {
	MONTH PrepaidUpdateOptionPeriodType
	YEAR  PrepaidUpdateOptionPeriodType
}

func GetPrepaidUpdateOptionPeriodTypeEnum() PrepaidUpdateOptionPeriodTypeEnum {
	return PrepaidUpdateOptionPeriodTypeEnum{
		MONTH: PrepaidUpdateOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidUpdateOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidUpdateOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidUpdateOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidUpdateOptionPeriodType) UnmarshalJSON(b []byte) error {
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
