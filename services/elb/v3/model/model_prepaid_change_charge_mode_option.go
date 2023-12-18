package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrepaidChangeChargeModeOption struct {
	IncludePublicip *bool `json:"include_publicip,omitempty"`

	PeriodType PrepaidChangeChargeModeOptionPeriodType `json:"period_type"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	AutoRenew *bool `json:"auto_renew,omitempty"`

	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o PrepaidChangeChargeModeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidChangeChargeModeOption struct{}"
	}

	return strings.Join([]string{"PrepaidChangeChargeModeOption", string(data)}, " ")
}

type PrepaidChangeChargeModeOptionPeriodType struct {
	value string
}

type PrepaidChangeChargeModeOptionPeriodTypeEnum struct {
	MONTH PrepaidChangeChargeModeOptionPeriodType
	YEAR  PrepaidChangeChargeModeOptionPeriodType
}

func GetPrepaidChangeChargeModeOptionPeriodTypeEnum() PrepaidChangeChargeModeOptionPeriodTypeEnum {
	return PrepaidChangeChargeModeOptionPeriodTypeEnum{
		MONTH: PrepaidChangeChargeModeOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidChangeChargeModeOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidChangeChargeModeOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidChangeChargeModeOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidChangeChargeModeOptionPeriodType) UnmarshalJSON(b []byte) error {
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
