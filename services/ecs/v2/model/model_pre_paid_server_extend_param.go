package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrePaidServerExtendParam struct {
	ChargingMode *PrePaidServerExtendParamChargingMode `json:"chargingMode,omitempty"`

	RegionID *string `json:"regionID,omitempty"`

	PeriodType *PrePaidServerExtendParamPeriodType `json:"periodType,omitempty"`

	PeriodNum *int32 `json:"periodNum,omitempty"`

	IsAutoRenew *PrePaidServerExtendParamIsAutoRenew `json:"isAutoRenew,omitempty"`

	IsAutoPay *PrePaidServerExtendParamIsAutoPay `json:"isAutoPay,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SupportAutoRecovery *bool `json:"support_auto_recovery,omitempty"`

	MarketType *string `json:"marketType,omitempty"`

	SpotPrice *string `json:"spotPrice,omitempty"`

	DiskPrior *string `json:"diskPrior,omitempty"`

	SpotDurationHours *int32 `json:"spot_duration_hours,omitempty"`

	InterruptionPolicy *PrePaidServerExtendParamInterruptionPolicy `json:"interruption_policy,omitempty"`

	SpotDurationCount *int32 `json:"spot_duration_count,omitempty"`
}

func (o PrePaidServerExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerExtendParam struct{}"
	}

	return strings.Join([]string{"PrePaidServerExtendParam", string(data)}, " ")
}

type PrePaidServerExtendParamChargingMode struct {
	value string
}

type PrePaidServerExtendParamChargingModeEnum struct {
	PRE_PAID  PrePaidServerExtendParamChargingMode
	POST_PAID PrePaidServerExtendParamChargingMode
}

func GetPrePaidServerExtendParamChargingModeEnum() PrePaidServerExtendParamChargingModeEnum {
	return PrePaidServerExtendParamChargingModeEnum{
		PRE_PAID: PrePaidServerExtendParamChargingMode{
			value: "prePaid",
		},
		POST_PAID: PrePaidServerExtendParamChargingMode{
			value: "postPaid",
		},
	}
}

func (c PrePaidServerExtendParamChargingMode) Value() string {
	return c.value
}

func (c PrePaidServerExtendParamChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerExtendParamChargingMode) UnmarshalJSON(b []byte) error {
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

type PrePaidServerExtendParamPeriodType struct {
	value string
}

type PrePaidServerExtendParamPeriodTypeEnum struct {
	MONTH PrePaidServerExtendParamPeriodType
	YEAR  PrePaidServerExtendParamPeriodType
}

func GetPrePaidServerExtendParamPeriodTypeEnum() PrePaidServerExtendParamPeriodTypeEnum {
	return PrePaidServerExtendParamPeriodTypeEnum{
		MONTH: PrePaidServerExtendParamPeriodType{
			value: "month",
		},
		YEAR: PrePaidServerExtendParamPeriodType{
			value: "year",
		},
	}
}

func (c PrePaidServerExtendParamPeriodType) Value() string {
	return c.value
}

func (c PrePaidServerExtendParamPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerExtendParamPeriodType) UnmarshalJSON(b []byte) error {
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

type PrePaidServerExtendParamIsAutoRenew struct {
	value string
}

type PrePaidServerExtendParamIsAutoRenewEnum struct {
	TRUE  PrePaidServerExtendParamIsAutoRenew
	FALSE PrePaidServerExtendParamIsAutoRenew
}

func GetPrePaidServerExtendParamIsAutoRenewEnum() PrePaidServerExtendParamIsAutoRenewEnum {
	return PrePaidServerExtendParamIsAutoRenewEnum{
		TRUE: PrePaidServerExtendParamIsAutoRenew{
			value: "true",
		},
		FALSE: PrePaidServerExtendParamIsAutoRenew{
			value: "false",
		},
	}
}

func (c PrePaidServerExtendParamIsAutoRenew) Value() string {
	return c.value
}

func (c PrePaidServerExtendParamIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerExtendParamIsAutoRenew) UnmarshalJSON(b []byte) error {
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

type PrePaidServerExtendParamIsAutoPay struct {
	value string
}

type PrePaidServerExtendParamIsAutoPayEnum struct {
	TRUE  PrePaidServerExtendParamIsAutoPay
	FALSE PrePaidServerExtendParamIsAutoPay
}

func GetPrePaidServerExtendParamIsAutoPayEnum() PrePaidServerExtendParamIsAutoPayEnum {
	return PrePaidServerExtendParamIsAutoPayEnum{
		TRUE: PrePaidServerExtendParamIsAutoPay{
			value: "true",
		},
		FALSE: PrePaidServerExtendParamIsAutoPay{
			value: "false",
		},
	}
}

func (c PrePaidServerExtendParamIsAutoPay) Value() string {
	return c.value
}

func (c PrePaidServerExtendParamIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerExtendParamIsAutoPay) UnmarshalJSON(b []byte) error {
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

type PrePaidServerExtendParamInterruptionPolicy struct {
	value string
}

type PrePaidServerExtendParamInterruptionPolicyEnum struct {
	IMMEDIATE PrePaidServerExtendParamInterruptionPolicy
}

func GetPrePaidServerExtendParamInterruptionPolicyEnum() PrePaidServerExtendParamInterruptionPolicyEnum {
	return PrePaidServerExtendParamInterruptionPolicyEnum{
		IMMEDIATE: PrePaidServerExtendParamInterruptionPolicy{
			value: "immediate",
		},
	}
}

func (c PrePaidServerExtendParamInterruptionPolicy) Value() string {
	return c.value
}

func (c PrePaidServerExtendParamInterruptionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerExtendParamInterruptionPolicy) UnmarshalJSON(b []byte) error {
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
