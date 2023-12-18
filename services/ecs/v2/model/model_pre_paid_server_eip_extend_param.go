package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrePaidServerEipExtendParam struct {
	ChargingMode *PrePaidServerEipExtendParamChargingMode `json:"chargingMode,omitempty"`
}

func (o PrePaidServerEipExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerEipExtendParam struct{}"
	}

	return strings.Join([]string{"PrePaidServerEipExtendParam", string(data)}, " ")
}

type PrePaidServerEipExtendParamChargingMode struct {
	value string
}

type PrePaidServerEipExtendParamChargingModeEnum struct {
	PRE_PAID  PrePaidServerEipExtendParamChargingMode
	POST_PAID PrePaidServerEipExtendParamChargingMode
}

func GetPrePaidServerEipExtendParamChargingModeEnum() PrePaidServerEipExtendParamChargingModeEnum {
	return PrePaidServerEipExtendParamChargingModeEnum{
		PRE_PAID: PrePaidServerEipExtendParamChargingMode{
			value: "prePaid",
		},
		POST_PAID: PrePaidServerEipExtendParamChargingMode{
			value: "postPaid",
		},
	}
}

func (c PrePaidServerEipExtendParamChargingMode) Value() string {
	return c.value
}

func (c PrePaidServerEipExtendParamChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerEipExtendParamChargingMode) UnmarshalJSON(b []byte) error {
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
