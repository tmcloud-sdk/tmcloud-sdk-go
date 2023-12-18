package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BandwidthResult struct {
	Size *int32 `json:"size,omitempty"`

	ShareType *BandwidthResultShareType `json:"share_type,omitempty"`

	ChargingMode *BandwidthResultChargingMode `json:"charging_mode,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o BandwidthResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthResult struct{}"
	}

	return strings.Join([]string{"BandwidthResult", string(data)}, " ")
}

type BandwidthResultShareType struct {
	value string
}

type BandwidthResultShareTypeEnum struct {
	PER   BandwidthResultShareType
	WHOLE BandwidthResultShareType
}

func GetBandwidthResultShareTypeEnum() BandwidthResultShareTypeEnum {
	return BandwidthResultShareTypeEnum{
		PER: BandwidthResultShareType{
			value: "PER",
		},
		WHOLE: BandwidthResultShareType{
			value: "WHOLE",
		},
	}
}

func (c BandwidthResultShareType) Value() string {
	return c.value
}

func (c BandwidthResultShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthResultShareType) UnmarshalJSON(b []byte) error {
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

type BandwidthResultChargingMode struct {
	value string
}

type BandwidthResultChargingModeEnum struct {
	BANDWIDTH BandwidthResultChargingMode
	TRAFFIC   BandwidthResultChargingMode
}

func GetBandwidthResultChargingModeEnum() BandwidthResultChargingModeEnum {
	return BandwidthResultChargingModeEnum{
		BANDWIDTH: BandwidthResultChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: BandwidthResultChargingMode{
			value: "traffic",
		},
	}
}

func (c BandwidthResultChargingMode) Value() string {
	return c.value
}

func (c BandwidthResultChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthResultChargingMode) UnmarshalJSON(b []byte) error {
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
