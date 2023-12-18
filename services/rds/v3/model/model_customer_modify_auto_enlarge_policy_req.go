package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CustomerModifyAutoEnlargePolicyReq struct {
	SwitchOption bool `json:"switch_option"`

	LimitSize *int32 `json:"limit_size,omitempty"`

	TriggerThreshold *CustomerModifyAutoEnlargePolicyReqTriggerThreshold `json:"trigger_threshold,omitempty"`
}

func (o CustomerModifyAutoEnlargePolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerModifyAutoEnlargePolicyReq struct{}"
	}

	return strings.Join([]string{"CustomerModifyAutoEnlargePolicyReq", string(data)}, " ")
}

type CustomerModifyAutoEnlargePolicyReqTriggerThreshold struct {
	value int32
}

type CustomerModifyAutoEnlargePolicyReqTriggerThresholdEnum struct {
	E_10 CustomerModifyAutoEnlargePolicyReqTriggerThreshold
	E_15 CustomerModifyAutoEnlargePolicyReqTriggerThreshold
	E_20 CustomerModifyAutoEnlargePolicyReqTriggerThreshold
}

func GetCustomerModifyAutoEnlargePolicyReqTriggerThresholdEnum() CustomerModifyAutoEnlargePolicyReqTriggerThresholdEnum {
	return CustomerModifyAutoEnlargePolicyReqTriggerThresholdEnum{
		E_10: CustomerModifyAutoEnlargePolicyReqTriggerThreshold{
			value: 10,
		}, E_15: CustomerModifyAutoEnlargePolicyReqTriggerThreshold{
			value: 15,
		}, E_20: CustomerModifyAutoEnlargePolicyReqTriggerThreshold{
			value: 20,
		},
	}
}

func (c CustomerModifyAutoEnlargePolicyReqTriggerThreshold) Value() int32 {
	return c.value
}

func (c CustomerModifyAutoEnlargePolicyReqTriggerThreshold) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomerModifyAutoEnlargePolicyReqTriggerThreshold) UnmarshalJSON(b []byte) error {
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
