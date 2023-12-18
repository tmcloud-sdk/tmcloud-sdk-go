package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaServerBlockDeviceMapping struct {
	SourceType NovaServerBlockDeviceMappingSourceType `json:"source_type"`

	DestinationType *NovaServerBlockDeviceMappingDestinationType `json:"destination_type,omitempty"`

	GuestFormat *string `json:"guest_format,omitempty"`

	DeviceName *string `json:"device_name,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	BootIndex *string `json:"boot_index,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	VolumeSize *int32 `json:"volume_size,omitempty"`

	VolumeType *string `json:"volume_type,omitempty"`
}

func (o NovaServerBlockDeviceMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerBlockDeviceMapping struct{}"
	}

	return strings.Join([]string{"NovaServerBlockDeviceMapping", string(data)}, " ")
}

type NovaServerBlockDeviceMappingSourceType struct {
	value string
}

type NovaServerBlockDeviceMappingSourceTypeEnum struct {
	BLANK    NovaServerBlockDeviceMappingSourceType
	SNAPSHOT NovaServerBlockDeviceMappingSourceType
	VOLUME   NovaServerBlockDeviceMappingSourceType
	IMAGE    NovaServerBlockDeviceMappingSourceType
}

func GetNovaServerBlockDeviceMappingSourceTypeEnum() NovaServerBlockDeviceMappingSourceTypeEnum {
	return NovaServerBlockDeviceMappingSourceTypeEnum{
		BLANK: NovaServerBlockDeviceMappingSourceType{
			value: "blank",
		},
		SNAPSHOT: NovaServerBlockDeviceMappingSourceType{
			value: "snapshot",
		},
		VOLUME: NovaServerBlockDeviceMappingSourceType{
			value: "volume",
		},
		IMAGE: NovaServerBlockDeviceMappingSourceType{
			value: "image",
		},
	}
}

func (c NovaServerBlockDeviceMappingSourceType) Value() string {
	return c.value
}

func (c NovaServerBlockDeviceMappingSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerBlockDeviceMappingSourceType) UnmarshalJSON(b []byte) error {
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

type NovaServerBlockDeviceMappingDestinationType struct {
	value string
}

type NovaServerBlockDeviceMappingDestinationTypeEnum struct {
	VOLUME NovaServerBlockDeviceMappingDestinationType
}

func GetNovaServerBlockDeviceMappingDestinationTypeEnum() NovaServerBlockDeviceMappingDestinationTypeEnum {
	return NovaServerBlockDeviceMappingDestinationTypeEnum{
		VOLUME: NovaServerBlockDeviceMappingDestinationType{
			value: "volume",
		},
	}
}

func (c NovaServerBlockDeviceMappingDestinationType) Value() string {
	return c.value
}

func (c NovaServerBlockDeviceMappingDestinationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerBlockDeviceMappingDestinationType) UnmarshalJSON(b []byte) error {
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
