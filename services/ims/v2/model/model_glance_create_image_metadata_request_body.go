package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceCreateImageMetadataRequestBody struct {
	OsVersion *string `json:"__os_version,omitempty"`

	ContainerFormat *string `json:"container_format,omitempty"`

	DiskFormat *GlanceCreateImageMetadataRequestBodyDiskFormat `json:"disk_format,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	Name *string `json:"name,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Visibility *string `json:"visibility,omitempty"`
}

func (o GlanceCreateImageMetadataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequestBody", string(data)}, " ")
}

type GlanceCreateImageMetadataRequestBodyDiskFormat struct {
	value string
}

type GlanceCreateImageMetadataRequestBodyDiskFormatEnum struct {
	VHD   GlanceCreateImageMetadataRequestBodyDiskFormat
	ZVHD  GlanceCreateImageMetadataRequestBodyDiskFormat
	ZVHD2 GlanceCreateImageMetadataRequestBodyDiskFormat
	RAW   GlanceCreateImageMetadataRequestBodyDiskFormat
	QCOW2 GlanceCreateImageMetadataRequestBodyDiskFormat
}

func GetGlanceCreateImageMetadataRequestBodyDiskFormatEnum() GlanceCreateImageMetadataRequestBodyDiskFormatEnum {
	return GlanceCreateImageMetadataRequestBodyDiskFormatEnum{
		VHD: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "zvhd",
		},
		ZVHD2: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "zvhd2",
		},
		RAW: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceCreateImageMetadataRequestBodyDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceCreateImageMetadataRequestBodyDiskFormat) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataRequestBodyDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataRequestBodyDiskFormat) UnmarshalJSON(b []byte) error {
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
