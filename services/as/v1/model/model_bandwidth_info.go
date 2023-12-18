package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BandwidthInfo struct {
	Size *int32 `json:"size,omitempty"`

	ShareType BandwidthInfoShareType `json:"share_type"`

	ChargingMode *BandwidthInfoChargingMode `json:"charging_mode,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o BandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfo struct{}"
	}

	return strings.Join([]string{"BandwidthInfo", string(data)}, " ")
}

type BandwidthInfoShareType struct {
	value string
}

type BandwidthInfoShareTypeEnum struct {
	PER   BandwidthInfoShareType
	WHOLE BandwidthInfoShareType
}

func GetBandwidthInfoShareTypeEnum() BandwidthInfoShareTypeEnum {
	return BandwidthInfoShareTypeEnum{
		PER: BandwidthInfoShareType{
			value: "PER",
		},
		WHOLE: BandwidthInfoShareType{
			value: "WHOLE",
		},
	}
}

func (c BandwidthInfoShareType) Value() string {
	return c.value
}

func (c BandwidthInfoShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthInfoShareType) UnmarshalJSON(b []byte) error {
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

type BandwidthInfoChargingMode struct {
	value string
}

type BandwidthInfoChargingModeEnum struct {
	BANDWIDTH BandwidthInfoChargingMode
	TRAFFIC   BandwidthInfoChargingMode
}

func GetBandwidthInfoChargingModeEnum() BandwidthInfoChargingModeEnum {
	return BandwidthInfoChargingModeEnum{
		BANDWIDTH: BandwidthInfoChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: BandwidthInfoChargingMode{
			value: "traffic",
		},
	}
}

func (c BandwidthInfoChargingMode) Value() string {
	return c.value
}

func (c BandwidthInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthInfoChargingMode) UnmarshalJSON(b []byte) error {
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
