package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type AvailableZones struct {
	Code *string `json:"code,omitempty"`

	Port *string `json:"port,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	ResourceAvailability *AvailableZonesResourceAvailability `json:"resource_availability,omitempty"`
}

func (o AvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZones struct{}"
	}

	return strings.Join([]string{"AvailableZones", string(data)}, " ")
}

type AvailableZonesResourceAvailability struct {
	value string
}

type AvailableZonesResourceAvailabilityEnum struct {
	TRUE  AvailableZonesResourceAvailability
	FALSE AvailableZonesResourceAvailability
}

func GetAvailableZonesResourceAvailabilityEnum() AvailableZonesResourceAvailabilityEnum {
	return AvailableZonesResourceAvailabilityEnum{
		TRUE: AvailableZonesResourceAvailability{
			value: "true",
		},
		FALSE: AvailableZonesResourceAvailability{
			value: "false",
		},
	}
}

func (c AvailableZonesResourceAvailability) Value() string {
	return c.value
}

func (c AvailableZonesResourceAvailability) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AvailableZonesResourceAvailability) UnmarshalJSON(b []byte) error {
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
