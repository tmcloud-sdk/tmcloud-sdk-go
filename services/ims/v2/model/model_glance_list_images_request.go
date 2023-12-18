package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceListImagesRequest struct {
	Imagetype *GlanceListImagesRequestImagetype `json:"__imagetype,omitempty"`

	Isregistered *bool `json:"__isregistered,omitempty"`

	OsBit *GlanceListImagesRequestOsBit `json:"__os_bit,omitempty"`

	OsType *GlanceListImagesRequestOsType `json:"__os_type,omitempty"`

	Platform *GlanceListImagesRequestPlatform `json:"__platform,omitempty"`

	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`

	SupportHighperformance *string `json:"__support_highperformance,omitempty"`

	SupportKvm *string `json:"__support_kvm,omitempty"`

	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`

	SupportLargememory *string `json:"__support_largememory,omitempty"`

	SupportXen *string `json:"__support_xen,omitempty"`

	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	ContainerFormat *string `json:"container_format,omitempty"`

	DiskFormat *GlanceListImagesRequestDiskFormat `json:"disk_format,omitempty"`

	Id *string `json:"id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	MemberStatus *string `json:"member_status,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	Status *GlanceListImagesRequestStatus `json:"status,omitempty"`

	Tag *string `json:"tag,omitempty"`

	Visibility *GlanceListImagesRequestVisibility `json:"visibility,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o GlanceListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImagesRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImagesRequest", string(data)}, " ")
}

type GlanceListImagesRequestImagetype struct {
	value string
}

type GlanceListImagesRequestImagetypeEnum struct {
	GOLD    GlanceListImagesRequestImagetype
	PRIVATE GlanceListImagesRequestImagetype
	SHARED  GlanceListImagesRequestImagetype
	MARKET  GlanceListImagesRequestImagetype
}

func GetGlanceListImagesRequestImagetypeEnum() GlanceListImagesRequestImagetypeEnum {
	return GlanceListImagesRequestImagetypeEnum{
		GOLD: GlanceListImagesRequestImagetype{
			value: "gold",
		},
		PRIVATE: GlanceListImagesRequestImagetype{
			value: "private",
		},
		SHARED: GlanceListImagesRequestImagetype{
			value: "shared",
		},
		MARKET: GlanceListImagesRequestImagetype{
			value: "market",
		},
	}
}

func (c GlanceListImagesRequestImagetype) Value() string {
	return c.value
}

func (c GlanceListImagesRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestOsBit struct {
	value string
}

type GlanceListImagesRequestOsBitEnum struct {
	E_32 GlanceListImagesRequestOsBit
	E_64 GlanceListImagesRequestOsBit
}

func GetGlanceListImagesRequestOsBitEnum() GlanceListImagesRequestOsBitEnum {
	return GlanceListImagesRequestOsBitEnum{
		E_32: GlanceListImagesRequestOsBit{
			value: "32",
		},
		E_64: GlanceListImagesRequestOsBit{
			value: "64",
		},
	}
}

func (c GlanceListImagesRequestOsBit) Value() string {
	return c.value
}

func (c GlanceListImagesRequestOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestOsType struct {
	value string
}

type GlanceListImagesRequestOsTypeEnum struct {
	LINUX   GlanceListImagesRequestOsType
	WINDOWS GlanceListImagesRequestOsType
	OTHER   GlanceListImagesRequestOsType
}

func GetGlanceListImagesRequestOsTypeEnum() GlanceListImagesRequestOsTypeEnum {
	return GlanceListImagesRequestOsTypeEnum{
		LINUX: GlanceListImagesRequestOsType{
			value: "Linux",
		},
		WINDOWS: GlanceListImagesRequestOsType{
			value: "Windows",
		},
		OTHER: GlanceListImagesRequestOsType{
			value: "Other",
		},
	}
}

func (c GlanceListImagesRequestOsType) Value() string {
	return c.value
}

func (c GlanceListImagesRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestOsType) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestPlatform struct {
	value string
}

type GlanceListImagesRequestPlatformEnum struct {
	WINDOWS      GlanceListImagesRequestPlatform
	UBUNTU       GlanceListImagesRequestPlatform
	RED_HAT      GlanceListImagesRequestPlatform
	SUSE         GlanceListImagesRequestPlatform
	CENT_OS      GlanceListImagesRequestPlatform
	DEBIAN       GlanceListImagesRequestPlatform
	OPEN_SUSE    GlanceListImagesRequestPlatform
	ORACLE_LINUX GlanceListImagesRequestPlatform
	FEDORA       GlanceListImagesRequestPlatform
	OTHER        GlanceListImagesRequestPlatform
	CORE_OS      GlanceListImagesRequestPlatform
	EULER_OS     GlanceListImagesRequestPlatform
}

func GetGlanceListImagesRequestPlatformEnum() GlanceListImagesRequestPlatformEnum {
	return GlanceListImagesRequestPlatformEnum{
		WINDOWS: GlanceListImagesRequestPlatform{
			value: "Windows",
		},
		UBUNTU: GlanceListImagesRequestPlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceListImagesRequestPlatform{
			value: "RedHat",
		},
		SUSE: GlanceListImagesRequestPlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceListImagesRequestPlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceListImagesRequestPlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceListImagesRequestPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceListImagesRequestPlatform{
			value: "Oracle Linux",
		},
		FEDORA: GlanceListImagesRequestPlatform{
			value: "Fedora",
		},
		OTHER: GlanceListImagesRequestPlatform{
			value: "Other",
		},
		CORE_OS: GlanceListImagesRequestPlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceListImagesRequestPlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceListImagesRequestPlatform) Value() string {
	return c.value
}

func (c GlanceListImagesRequestPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestPlatform) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestDiskFormat struct {
	value string
}

type GlanceListImagesRequestDiskFormatEnum struct {
	VHD   GlanceListImagesRequestDiskFormat
	ZVHD  GlanceListImagesRequestDiskFormat
	RAW   GlanceListImagesRequestDiskFormat
	QCOW2 GlanceListImagesRequestDiskFormat
}

func GetGlanceListImagesRequestDiskFormatEnum() GlanceListImagesRequestDiskFormatEnum {
	return GlanceListImagesRequestDiskFormatEnum{
		VHD: GlanceListImagesRequestDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceListImagesRequestDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceListImagesRequestDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceListImagesRequestDiskFormat{
			value: "qcow2",
		},
	}
}

func (c GlanceListImagesRequestDiskFormat) Value() string {
	return c.value
}

func (c GlanceListImagesRequestDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestStatus struct {
	value string
}

type GlanceListImagesRequestStatusEnum struct {
	QUEUED  GlanceListImagesRequestStatus
	SAVING  GlanceListImagesRequestStatus
	DELETED GlanceListImagesRequestStatus
	KILLED  GlanceListImagesRequestStatus
	ACTIVE  GlanceListImagesRequestStatus
}

func GetGlanceListImagesRequestStatusEnum() GlanceListImagesRequestStatusEnum {
	return GlanceListImagesRequestStatusEnum{
		QUEUED: GlanceListImagesRequestStatus{
			value: "queued",
		},
		SAVING: GlanceListImagesRequestStatus{
			value: "saving",
		},
		DELETED: GlanceListImagesRequestStatus{
			value: "deleted",
		},
		KILLED: GlanceListImagesRequestStatus{
			value: "killed",
		},
		ACTIVE: GlanceListImagesRequestStatus{
			value: "active",
		},
	}
}

func (c GlanceListImagesRequestStatus) Value() string {
	return c.value
}

func (c GlanceListImagesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestStatus) UnmarshalJSON(b []byte) error {
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

type GlanceListImagesRequestVisibility struct {
	value string
}

type GlanceListImagesRequestVisibilityEnum struct {
	PUBLIC  GlanceListImagesRequestVisibility
	PRIVATE GlanceListImagesRequestVisibility
	SHARED  GlanceListImagesRequestVisibility
}

func GetGlanceListImagesRequestVisibilityEnum() GlanceListImagesRequestVisibilityEnum {
	return GlanceListImagesRequestVisibilityEnum{
		PUBLIC: GlanceListImagesRequestVisibility{
			value: "public",
		},
		PRIVATE: GlanceListImagesRequestVisibility{
			value: "private",
		},
		SHARED: GlanceListImagesRequestVisibility{
			value: "shared",
		},
	}
}

func (c GlanceListImagesRequestVisibility) Value() string {
	return c.value
}

func (c GlanceListImagesRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceListImagesRequestVisibility) UnmarshalJSON(b []byte) error {
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
