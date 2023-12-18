package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceCreateImageMetadataResponse struct {
	Visibility *string `json:"visibility,omitempty"`

	Name *string `json:"name,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	ContainerFormat *string `json:"container_format,omitempty"`

	DiskFormat *GlanceCreateImageMetadataResponseDiskFormat `json:"disk_format,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	Status *GlanceCreateImageMetadataResponseStatus `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Self *string `json:"self,omitempty"`

	Id *string `json:"id,omitempty"`

	File *string `json:"file,omitempty"`

	Schema *string `json:"schema,omitempty"`

	ImageSourceType *string `json:"__image_source_type,omitempty"`

	ImageSize *string `json:"__image_size,omitempty"`

	Isregistered *string `json:"__isregistered,omitempty"`

	OsVersion *string `json:"__os_version,omitempty"`

	OsType *GlanceCreateImageMetadataResponseOsType `json:"__os_type,omitempty"`

	Platform *string `json:"__platform,omitempty"`

	OsBit *GlanceCreateImageMetadataResponseOsBit `json:"__os_bit,omitempty"`

	Imagetype *string `json:"__imagetype,omitempty"`

	VirtualEnvType *GlanceCreateImageMetadataResponseVirtualEnvType `json:"virtual_env_type,omitempty"`

	Owner *string `json:"owner,omitempty"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Properties *interface{} `json:"properties,omitempty"`

	RootOrigin *string `json:"__root_origin,omitempty"`

	Checksum *string `json:"checksum,omitempty"`

	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o GlanceCreateImageMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataResponse", string(data)}, " ")
}

type GlanceCreateImageMetadataResponseDiskFormat struct {
	value string
}

type GlanceCreateImageMetadataResponseDiskFormatEnum struct {
	VHD   GlanceCreateImageMetadataResponseDiskFormat
	ZVHD  GlanceCreateImageMetadataResponseDiskFormat
	RAW   GlanceCreateImageMetadataResponseDiskFormat
	QCOW2 GlanceCreateImageMetadataResponseDiskFormat
}

func GetGlanceCreateImageMetadataResponseDiskFormatEnum() GlanceCreateImageMetadataResponseDiskFormatEnum {
	return GlanceCreateImageMetadataResponseDiskFormatEnum{
		VHD: GlanceCreateImageMetadataResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceCreateImageMetadataResponseDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceCreateImageMetadataResponseDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceCreateImageMetadataResponseDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceCreateImageMetadataResponseDiskFormat) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseStatus struct {
	value string
}

type GlanceCreateImageMetadataResponseStatusEnum struct {
	QUEUED  GlanceCreateImageMetadataResponseStatus
	SAVING  GlanceCreateImageMetadataResponseStatus
	DELETED GlanceCreateImageMetadataResponseStatus
	KILLED  GlanceCreateImageMetadataResponseStatus
	ACTIVE  GlanceCreateImageMetadataResponseStatus
}

func GetGlanceCreateImageMetadataResponseStatusEnum() GlanceCreateImageMetadataResponseStatusEnum {
	return GlanceCreateImageMetadataResponseStatusEnum{
		QUEUED: GlanceCreateImageMetadataResponseStatus{
			value: "queued",
		},
		SAVING: GlanceCreateImageMetadataResponseStatus{
			value: "saving",
		},
		DELETED: GlanceCreateImageMetadataResponseStatus{
			value: "deleted",
		},
		KILLED: GlanceCreateImageMetadataResponseStatus{
			value: "killed",
		},
		ACTIVE: GlanceCreateImageMetadataResponseStatus{
			value: "active",
		},
	}
}

func (c GlanceCreateImageMetadataResponseStatus) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseStatus) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseOsType struct {
	value string
}

type GlanceCreateImageMetadataResponseOsTypeEnum struct {
	WINDOWS GlanceCreateImageMetadataResponseOsType
	LINUX   GlanceCreateImageMetadataResponseOsType
	OTHER   GlanceCreateImageMetadataResponseOsType
}

func GetGlanceCreateImageMetadataResponseOsTypeEnum() GlanceCreateImageMetadataResponseOsTypeEnum {
	return GlanceCreateImageMetadataResponseOsTypeEnum{
		WINDOWS: GlanceCreateImageMetadataResponseOsType{
			value: "Windows",
		},
		LINUX: GlanceCreateImageMetadataResponseOsType{
			value: "Linux",
		},
		OTHER: GlanceCreateImageMetadataResponseOsType{
			value: "other",
		},
	}
}

func (c GlanceCreateImageMetadataResponseOsType) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseOsType) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseOsBit struct {
	value string
}

type GlanceCreateImageMetadataResponseOsBitEnum struct {
	E_32 GlanceCreateImageMetadataResponseOsBit
	E_64 GlanceCreateImageMetadataResponseOsBit
}

func GetGlanceCreateImageMetadataResponseOsBitEnum() GlanceCreateImageMetadataResponseOsBitEnum {
	return GlanceCreateImageMetadataResponseOsBitEnum{
		E_32: GlanceCreateImageMetadataResponseOsBit{
			value: "32",
		},
		E_64: GlanceCreateImageMetadataResponseOsBit{
			value: "64",
		},
	}
}

func (c GlanceCreateImageMetadataResponseOsBit) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceCreateImageMetadataResponseVirtualEnvType struct {
	value string
}

type GlanceCreateImageMetadataResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceCreateImageMetadataResponseVirtualEnvType
	IRONIC         GlanceCreateImageMetadataResponseVirtualEnvType
	DATA_IMAGE     GlanceCreateImageMetadataResponseVirtualEnvType
}

func GetGlanceCreateImageMetadataResponseVirtualEnvTypeEnum() GlanceCreateImageMetadataResponseVirtualEnvTypeEnum {
	return GlanceCreateImageMetadataResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceCreateImageMetadataResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceCreateImageMetadataResponseVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceCreateImageMetadataResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceCreateImageMetadataResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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
