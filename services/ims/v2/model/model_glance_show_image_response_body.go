package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceShowImageResponseBody struct {
	BackupId *string `json:"__backup_id,omitempty"`

	DataOrigin *string `json:"__data_origin,omitempty"`

	Description *string `json:"__description,omitempty"`

	ImageSize string `json:"__image_size"`

	ImageSourceType GlanceShowImageResponseBodyImageSourceType `json:"__image_source_type"`

	Imagetype GlanceShowImageResponseBodyImagetype `json:"__imagetype"`

	Isregistered GlanceShowImageResponseBodyIsregistered `json:"__isregistered"`

	Originalimagename *string `json:"__originalimagename,omitempty"`

	OsBit *GlanceShowImageResponseBodyOsBit `json:"__os_bit,omitempty"`

	OsType GlanceShowImageResponseBodyOsType `json:"__os_type"`

	OsVersion *string `json:"__os_version,omitempty"`

	Platform *GlanceShowImageResponseBodyPlatform `json:"__platform,omitempty"`

	Productcode *string `json:"__productcode,omitempty"`

	SupportDiskintensive *string `json:"__support_diskintensive,omitempty"`

	SupportHighperformance *string `json:"__support_highperformance,omitempty"`

	SupportKvm *string `json:"__support_kvm,omitempty"`

	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty"`

	SupportKvmInfiniband *string `json:"__support_kvm_infiniband,omitempty"`

	SupportLargememory *string `json:"__support_largememory,omitempty"`

	SupportXen *string `json:"__support_xen,omitempty"`

	SupportXenGpuType *string `json:"__support_xen_gpu_type,omitempty"`

	SupportXenHana *string `json:"__support_xen_hana,omitempty"`

	Checksum *string `json:"checksum,omitempty"`

	ContainerFormat string `json:"container_format"`

	CreatedAt string `json:"created_at"`

	DiskFormat GlanceShowImageResponseBodyDiskFormat `json:"disk_format"`

	File string `json:"file"`

	Id string `json:"id"`

	MinDisk int32 `json:"min_disk"`

	MinRam int32 `json:"min_ram"`

	Name string `json:"name"`

	Owner string `json:"owner"`

	Protected bool `json:"protected"`

	Schema string `json:"schema"`

	Self string `json:"self"`

	Size *int64 `json:"size,omitempty"`

	Status GlanceShowImageResponseBodyStatus `json:"status"`

	Tags []string `json:"tags"`

	UpdatedAt string `json:"updated_at"`

	VirtualEnvType GlanceShowImageResponseBodyVirtualEnvType `json:"virtual_env_type"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Visibility GlanceShowImageResponseBodyVisibility `json:"visibility"`

	SupportFcInject *GlanceShowImageResponseBodySupportFcInject `json:"__support_fc_inject,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HwFirmwareType *GlanceShowImageResponseBodyHwFirmwareType `json:"hw_firmware_type,omitempty"`

	SupportArm *GlanceShowImageResponseBodySupportArm `json:"__support_arm,omitempty"`

	IsOffshelved *GlanceShowImageResponseBodyIsOffshelved `json:"__is_offshelved,omitempty"`

	Lazyloading *string `json:"__lazyloading,omitempty"`

	OsFeatureList *string `json:"__os_feature_list,omitempty"`

	RootOrigin *string `json:"__root_origin,omitempty"`

	SequenceNum *string `json:"__sequence_num,omitempty"`

	SupportAgentList *string `json:"__support_agent_list,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	ActiveAt *string `json:"active_at,omitempty"`

	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty"`

	MaxRam *string `json:"max_ram,omitempty"`

	ImageLocation *string `json:"__image_location,omitempty"`

	IsConfigInit *string `json:"__is_config_init,omitempty"`

	AccountCode *string `json:"__account_code,omitempty"`

	SupportAmd *string `json:"__support_amd,omitempty"`
}

func (o GlanceShowImageResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageResponseBody struct{}"
	}

	return strings.Join([]string{"GlanceShowImageResponseBody", string(data)}, " ")
}

type GlanceShowImageResponseBodyImageSourceType struct {
	value string
}

type GlanceShowImageResponseBodyImageSourceTypeEnum struct {
	UDS   GlanceShowImageResponseBodyImageSourceType
	SWIFT GlanceShowImageResponseBodyImageSourceType
}

func GetGlanceShowImageResponseBodyImageSourceTypeEnum() GlanceShowImageResponseBodyImageSourceTypeEnum {
	return GlanceShowImageResponseBodyImageSourceTypeEnum{
		UDS: GlanceShowImageResponseBodyImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceShowImageResponseBodyImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceShowImageResponseBodyImageSourceType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyImageSourceType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyImagetype struct {
	value string
}

type GlanceShowImageResponseBodyImagetypeEnum struct {
	GOLD    GlanceShowImageResponseBodyImagetype
	PRIVATE GlanceShowImageResponseBodyImagetype
	SHARED  GlanceShowImageResponseBodyImagetype
	MARKET  GlanceShowImageResponseBodyImagetype
}

func GetGlanceShowImageResponseBodyImagetypeEnum() GlanceShowImageResponseBodyImagetypeEnum {
	return GlanceShowImageResponseBodyImagetypeEnum{
		GOLD: GlanceShowImageResponseBodyImagetype{
			value: "gold",
		},
		PRIVATE: GlanceShowImageResponseBodyImagetype{
			value: "private",
		},
		SHARED: GlanceShowImageResponseBodyImagetype{
			value: "shared",
		},
		MARKET: GlanceShowImageResponseBodyImagetype{
			value: "market",
		},
	}
}

func (c GlanceShowImageResponseBodyImagetype) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyIsregistered struct {
	value string
}

type GlanceShowImageResponseBodyIsregisteredEnum struct {
	TRUE  GlanceShowImageResponseBodyIsregistered
	FALSE GlanceShowImageResponseBodyIsregistered
}

func GetGlanceShowImageResponseBodyIsregisteredEnum() GlanceShowImageResponseBodyIsregisteredEnum {
	return GlanceShowImageResponseBodyIsregisteredEnum{
		TRUE: GlanceShowImageResponseBodyIsregistered{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodyIsregistered{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodyIsregistered) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyIsregistered) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyOsBit struct {
	value string
}

type GlanceShowImageResponseBodyOsBitEnum struct {
	E_32 GlanceShowImageResponseBodyOsBit
	E_64 GlanceShowImageResponseBodyOsBit
}

func GetGlanceShowImageResponseBodyOsBitEnum() GlanceShowImageResponseBodyOsBitEnum {
	return GlanceShowImageResponseBodyOsBitEnum{
		E_32: GlanceShowImageResponseBodyOsBit{
			value: "32",
		},
		E_64: GlanceShowImageResponseBodyOsBit{
			value: "64",
		},
	}
}

func (c GlanceShowImageResponseBodyOsBit) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyOsType struct {
	value string
}

type GlanceShowImageResponseBodyOsTypeEnum struct {
	LINUX   GlanceShowImageResponseBodyOsType
	WINDOWS GlanceShowImageResponseBodyOsType
	OTHER   GlanceShowImageResponseBodyOsType
}

func GetGlanceShowImageResponseBodyOsTypeEnum() GlanceShowImageResponseBodyOsTypeEnum {
	return GlanceShowImageResponseBodyOsTypeEnum{
		LINUX: GlanceShowImageResponseBodyOsType{
			value: "Linux",
		},
		WINDOWS: GlanceShowImageResponseBodyOsType{
			value: "Windows",
		},
		OTHER: GlanceShowImageResponseBodyOsType{
			value: "Other",
		},
	}
}

func (c GlanceShowImageResponseBodyOsType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyOsType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyPlatform struct {
	value string
}

type GlanceShowImageResponseBodyPlatformEnum struct {
	WINDOWS      GlanceShowImageResponseBodyPlatform
	UBUNTU       GlanceShowImageResponseBodyPlatform
	RED_HAT      GlanceShowImageResponseBodyPlatform
	SUSE         GlanceShowImageResponseBodyPlatform
	CENT_OS      GlanceShowImageResponseBodyPlatform
	DEBIAN       GlanceShowImageResponseBodyPlatform
	OPEN_SUSE    GlanceShowImageResponseBodyPlatform
	ORACLE_LINUX GlanceShowImageResponseBodyPlatform
	FEDORA       GlanceShowImageResponseBodyPlatform
	OTHER        GlanceShowImageResponseBodyPlatform
	CORE_OS      GlanceShowImageResponseBodyPlatform
	EULER_OS     GlanceShowImageResponseBodyPlatform
}

func GetGlanceShowImageResponseBodyPlatformEnum() GlanceShowImageResponseBodyPlatformEnum {
	return GlanceShowImageResponseBodyPlatformEnum{
		WINDOWS: GlanceShowImageResponseBodyPlatform{
			value: "Windows",
		},
		UBUNTU: GlanceShowImageResponseBodyPlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceShowImageResponseBodyPlatform{
			value: "RedHat",
		},
		SUSE: GlanceShowImageResponseBodyPlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceShowImageResponseBodyPlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceShowImageResponseBodyPlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceShowImageResponseBodyPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceShowImageResponseBodyPlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceShowImageResponseBodyPlatform{
			value: "Fedora",
		},
		OTHER: GlanceShowImageResponseBodyPlatform{
			value: "Other",
		},
		CORE_OS: GlanceShowImageResponseBodyPlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceShowImageResponseBodyPlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceShowImageResponseBodyPlatform) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyPlatform) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyDiskFormat struct {
	value string
}

type GlanceShowImageResponseBodyDiskFormatEnum struct {
	VHD   GlanceShowImageResponseBodyDiskFormat
	ZVHD  GlanceShowImageResponseBodyDiskFormat
	RAW   GlanceShowImageResponseBodyDiskFormat
	QCOW2 GlanceShowImageResponseBodyDiskFormat
	ZVHD2 GlanceShowImageResponseBodyDiskFormat
}

func GetGlanceShowImageResponseBodyDiskFormatEnum() GlanceShowImageResponseBodyDiskFormatEnum {
	return GlanceShowImageResponseBodyDiskFormatEnum{
		VHD: GlanceShowImageResponseBodyDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceShowImageResponseBodyDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceShowImageResponseBodyDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceShowImageResponseBodyDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceShowImageResponseBodyDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceShowImageResponseBodyDiskFormat) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyStatus struct {
	value string
}

type GlanceShowImageResponseBodyStatusEnum struct {
	QUEUED  GlanceShowImageResponseBodyStatus
	SAVING  GlanceShowImageResponseBodyStatus
	DELETED GlanceShowImageResponseBodyStatus
	KILLED  GlanceShowImageResponseBodyStatus
	ACTIVE  GlanceShowImageResponseBodyStatus
}

func GetGlanceShowImageResponseBodyStatusEnum() GlanceShowImageResponseBodyStatusEnum {
	return GlanceShowImageResponseBodyStatusEnum{
		QUEUED: GlanceShowImageResponseBodyStatus{
			value: "queued",
		},
		SAVING: GlanceShowImageResponseBodyStatus{
			value: "saving",
		},
		DELETED: GlanceShowImageResponseBodyStatus{
			value: "deleted",
		},
		KILLED: GlanceShowImageResponseBodyStatus{
			value: "killed",
		},
		ACTIVE: GlanceShowImageResponseBodyStatus{
			value: "active",
		},
	}
}

func (c GlanceShowImageResponseBodyStatus) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyVirtualEnvType struct {
	value string
}

type GlanceShowImageResponseBodyVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceShowImageResponseBodyVirtualEnvType
	IRONIC         GlanceShowImageResponseBodyVirtualEnvType
	DATA_IMAGE     GlanceShowImageResponseBodyVirtualEnvType
}

func GetGlanceShowImageResponseBodyVirtualEnvTypeEnum() GlanceShowImageResponseBodyVirtualEnvTypeEnum {
	return GlanceShowImageResponseBodyVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceShowImageResponseBodyVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceShowImageResponseBodyVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceShowImageResponseBodyVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceShowImageResponseBodyVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyVisibility struct {
	value string
}

type GlanceShowImageResponseBodyVisibilityEnum struct {
	PRIVATE GlanceShowImageResponseBodyVisibility
	PUBLIC  GlanceShowImageResponseBodyVisibility
	SHARED  GlanceShowImageResponseBodyVisibility
}

func GetGlanceShowImageResponseBodyVisibilityEnum() GlanceShowImageResponseBodyVisibilityEnum {
	return GlanceShowImageResponseBodyVisibilityEnum{
		PRIVATE: GlanceShowImageResponseBodyVisibility{
			value: "private",
		},
		PUBLIC: GlanceShowImageResponseBodyVisibility{
			value: "public",
		},
		SHARED: GlanceShowImageResponseBodyVisibility{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseBodyVisibility) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyVisibility) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodySupportFcInject struct {
	value string
}

type GlanceShowImageResponseBodySupportFcInjectEnum struct {
	TRUE  GlanceShowImageResponseBodySupportFcInject
	FALSE GlanceShowImageResponseBodySupportFcInject
}

func GetGlanceShowImageResponseBodySupportFcInjectEnum() GlanceShowImageResponseBodySupportFcInjectEnum {
	return GlanceShowImageResponseBodySupportFcInjectEnum{
		TRUE: GlanceShowImageResponseBodySupportFcInject{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodySupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodySupportFcInject) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodySupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodySupportFcInject) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyHwFirmwareType struct {
	value string
}

type GlanceShowImageResponseBodyHwFirmwareTypeEnum struct {
	BIOS GlanceShowImageResponseBodyHwFirmwareType
	UEFI GlanceShowImageResponseBodyHwFirmwareType
}

func GetGlanceShowImageResponseBodyHwFirmwareTypeEnum() GlanceShowImageResponseBodyHwFirmwareTypeEnum {
	return GlanceShowImageResponseBodyHwFirmwareTypeEnum{
		BIOS: GlanceShowImageResponseBodyHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceShowImageResponseBodyHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceShowImageResponseBodyHwFirmwareType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodySupportArm struct {
	value string
}

type GlanceShowImageResponseBodySupportArmEnum struct {
	TRUE  GlanceShowImageResponseBodySupportArm
	FALSE GlanceShowImageResponseBodySupportArm
}

func GetGlanceShowImageResponseBodySupportArmEnum() GlanceShowImageResponseBodySupportArmEnum {
	return GlanceShowImageResponseBodySupportArmEnum{
		TRUE: GlanceShowImageResponseBodySupportArm{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodySupportArm{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodySupportArm) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodySupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodySupportArm) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseBodyIsOffshelved struct {
	value string
}

type GlanceShowImageResponseBodyIsOffshelvedEnum struct {
	TRUE  GlanceShowImageResponseBodyIsOffshelved
	FALSE GlanceShowImageResponseBodyIsOffshelved
}

func GetGlanceShowImageResponseBodyIsOffshelvedEnum() GlanceShowImageResponseBodyIsOffshelvedEnum {
	return GlanceShowImageResponseBodyIsOffshelvedEnum{
		TRUE: GlanceShowImageResponseBodyIsOffshelved{
			value: "true",
		},
		FALSE: GlanceShowImageResponseBodyIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseBodyIsOffshelved) Value() string {
	return c.value
}

func (c GlanceShowImageResponseBodyIsOffshelved) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseBodyIsOffshelved) UnmarshalJSON(b []byte) error {
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
