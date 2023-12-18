package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BssParamEntity struct {
	IsAutoPay *BssParamEntityIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o BssParamEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamEntity struct{}"
	}

	return strings.Join([]string{"BssParamEntity", string(data)}, " ")
}

type BssParamEntityIsAutoPay struct {
	value string
}

type BssParamEntityIsAutoPayEnum struct {
	TRUE  BssParamEntityIsAutoPay
	FALSE BssParamEntityIsAutoPay
}

func GetBssParamEntityIsAutoPayEnum() BssParamEntityIsAutoPayEnum {
	return BssParamEntityIsAutoPayEnum{
		TRUE: BssParamEntityIsAutoPay{
			value: "true",
		},
		FALSE: BssParamEntityIsAutoPay{
			value: "false",
		},
	}
}

func (c BssParamEntityIsAutoPay) Value() string {
	return c.value
}

func (c BssParamEntityIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamEntityIsAutoPay) UnmarshalJSON(b []byte) error {
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
