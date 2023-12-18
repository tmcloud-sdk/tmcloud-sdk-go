package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PeriodOrderInfo struct {
	PeriodType *PeriodOrderInfoPeriodType `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *PeriodOrderInfoIsAutoRenew `json:"is_auto_renew,omitempty"`
}

func (o PeriodOrderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderInfo struct{}"
	}

	return strings.Join([]string{"PeriodOrderInfo", string(data)}, " ")
}

type PeriodOrderInfoPeriodType struct {
	value int32
}

type PeriodOrderInfoPeriodTypeEnum struct {
	E_2 PeriodOrderInfoPeriodType
	E_3 PeriodOrderInfoPeriodType
}

func GetPeriodOrderInfoPeriodTypeEnum() PeriodOrderInfoPeriodTypeEnum {
	return PeriodOrderInfoPeriodTypeEnum{
		E_2: PeriodOrderInfoPeriodType{
			value: 2,
		}, E_3: PeriodOrderInfoPeriodType{
			value: 3,
		},
	}
}

func (c PeriodOrderInfoPeriodType) Value() int32 {
	return c.value
}

func (c PeriodOrderInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodOrderInfoPeriodType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type PeriodOrderInfoIsAutoRenew struct {
	value int32
}

type PeriodOrderInfoIsAutoRenewEnum struct {
	E_0 PeriodOrderInfoIsAutoRenew
	E_1 PeriodOrderInfoIsAutoRenew
}

func GetPeriodOrderInfoIsAutoRenewEnum() PeriodOrderInfoIsAutoRenewEnum {
	return PeriodOrderInfoIsAutoRenewEnum{
		E_0: PeriodOrderInfoIsAutoRenew{
			value: 0,
		}, E_1: PeriodOrderInfoIsAutoRenew{
			value: 1,
		},
	}
}

func (c PeriodOrderInfoIsAutoRenew) Value() int32 {
	return c.value
}

func (c PeriodOrderInfoIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodOrderInfoIsAutoRenew) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
