package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrepaidCreateOption struct {
	PeriodType PrepaidCreateOptionPeriodType `json:"period_type"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	AutoRenew *bool `json:"auto_renew,omitempty"`

	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o PrepaidCreateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidCreateOption struct{}"
	}

	return strings.Join([]string{"PrepaidCreateOption", string(data)}, " ")
}

type PrepaidCreateOptionPeriodType struct {
	value string
}

type PrepaidCreateOptionPeriodTypeEnum struct {
	MONTH PrepaidCreateOptionPeriodType
	YEAR  PrepaidCreateOptionPeriodType
}

func GetPrepaidCreateOptionPeriodTypeEnum() PrepaidCreateOptionPeriodTypeEnum {
	return PrepaidCreateOptionPeriodTypeEnum{
		MONTH: PrepaidCreateOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidCreateOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidCreateOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidCreateOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidCreateOptionPeriodType) UnmarshalJSON(b []byte) error {
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
