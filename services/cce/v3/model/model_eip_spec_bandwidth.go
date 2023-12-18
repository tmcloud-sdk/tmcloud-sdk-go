package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type EipSpecBandwidth struct {
	Size *int32 `json:"size,omitempty"`

	Sharetype *EipSpecBandwidthSharetype `json:"sharetype,omitempty"`
}

func (o EipSpecBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipSpecBandwidth struct{}"
	}

	return strings.Join([]string{"EipSpecBandwidth", string(data)}, " ")
}

type EipSpecBandwidthSharetype struct {
	value string
}

type EipSpecBandwidthSharetypeEnum struct {
	PER   EipSpecBandwidthSharetype
	WHOLE EipSpecBandwidthSharetype
}

func GetEipSpecBandwidthSharetypeEnum() EipSpecBandwidthSharetypeEnum {
	return EipSpecBandwidthSharetypeEnum{
		PER: EipSpecBandwidthSharetype{
			value: "PER",
		},
		WHOLE: EipSpecBandwidthSharetype{
			value: "WHOLE",
		},
	}
}

func (c EipSpecBandwidthSharetype) Value() string {
	return c.value
}

func (c EipSpecBandwidthSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipSpecBandwidthSharetype) UnmarshalJSON(b []byte) error {
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
