package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ChargeInfo struct {
	ChargeMode ChargeInfoChargeMode `json:"charge_mode"`

	PeriodType *ChargeInfoPeriodType `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfo struct{}"
	}

	return strings.Join([]string{"ChargeInfo", string(data)}, " ")
}

type ChargeInfoChargeMode struct {
	value string
}

type ChargeInfoChargeModeEnum struct {
	PRE_PAID  ChargeInfoChargeMode
	POST_PAID ChargeInfoChargeMode
}

func GetChargeInfoChargeModeEnum() ChargeInfoChargeModeEnum {
	return ChargeInfoChargeModeEnum{
		PRE_PAID: ChargeInfoChargeMode{
			value: "prePaid",
		},
		POST_PAID: ChargeInfoChargeMode{
			value: "postPaid",
		},
	}
}

func (c ChargeInfoChargeMode) Value() string {
	return c.value
}

func (c ChargeInfoChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargeInfoChargeMode) UnmarshalJSON(b []byte) error {
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

type ChargeInfoPeriodType struct {
	value string
}

type ChargeInfoPeriodTypeEnum struct {
	MONTH ChargeInfoPeriodType
	YEAR  ChargeInfoPeriodType
}

func GetChargeInfoPeriodTypeEnum() ChargeInfoPeriodTypeEnum {
	return ChargeInfoPeriodTypeEnum{
		MONTH: ChargeInfoPeriodType{
			value: "month",
		},
		YEAR: ChargeInfoPeriodType{
			value: "year",
		},
	}
}

func (c ChargeInfoPeriodType) Value() string {
	return c.value
}

func (c ChargeInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargeInfoPeriodType) UnmarshalJSON(b []byte) error {
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
