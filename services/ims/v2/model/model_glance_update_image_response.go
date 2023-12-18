package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceUpdateImageResponse struct {
	BackupId *string `json:"__backup_id,omitempty"`

	DataOrigin *string `json:"__data_origin,omitempty"`

	Description *string `json:"__description,omitempty"`

	ImageSize *string `json:"__image_size,omitempty"`

	ImageSourceType *GlanceUpdateImageResponseImageSourceType `json:"__image_source_type,omitempty"`

	Imagetype *GlanceUpdateImageResponseImagetype `json:"__imagetype,omitempty"`

	Isregistered *GlanceUpdateImageResponseIsregistered `json:"__isregistered,omitempty"`

	Originalimagename *string `json:"__originalimagename,omitempty"`

	OsBit *GlanceUpdateImageResponseOsBit `json:"__os_bit,omitempty"`

	OsType *GlanceUpdateImageResponseOsType `json:"__os_type,omitempty"`

	OsVersion *string `json:"__os_version,omitempty"`

	Platform *GlanceUpdateImageResponsePlatform `json:"__platform,omitempty"`

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

	ContainerFormat *string `json:"container_format,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	DiskFormat *GlanceUpdateImageResponseDiskFormat `json:"disk_format,omitempty"`

	File *string `json:"file,omitempty"`

	Id *string `json:"id,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	Schema *string `json:"schema,omitempty"`

	Self *string `json:"self,omitempty"`

	Size *int64 `json:"size,omitempty"`

	Status *GlanceUpdateImageResponseStatus `json:"status,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	VirtualEnvType *GlanceUpdateImageResponseVirtualEnvType `json:"virtual_env_type,omitempty"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Visibility *GlanceUpdateImageResponseVisibility `json:"visibility,omitempty"`

	SupportFcInject *GlanceUpdateImageResponseSupportFcInject `json:"__support_fc_inject,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HwFirmwareType *GlanceUpdateImageResponseHwFirmwareType `json:"hw_firmware_type,omitempty"`

	SupportArm *GlanceUpdateImageResponseSupportArm `json:"__support_arm,omitempty"`

	IsOffshelved *GlanceUpdateImageResponseIsOffshelved `json:"__is_offshelved,omitempty"`

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

	SupportAmd     *string `json:"__support_amd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceUpdateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageResponse struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageResponse", string(data)}, " ")
}

type GlanceUpdateImageResponseImageSourceType struct {
	value string
}

type GlanceUpdateImageResponseImageSourceTypeEnum struct {
	UDS   GlanceUpdateImageResponseImageSourceType
	SWIFT GlanceUpdateImageResponseImageSourceType
}

func GetGlanceUpdateImageResponseImageSourceTypeEnum() GlanceUpdateImageResponseImageSourceTypeEnum {
	return GlanceUpdateImageResponseImageSourceTypeEnum{
		UDS: GlanceUpdateImageResponseImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceUpdateImageResponseImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceUpdateImageResponseImageSourceType) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseImageSourceType) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseImagetype struct {
	value string
}

type GlanceUpdateImageResponseImagetypeEnum struct {
	GOLD    GlanceUpdateImageResponseImagetype
	PRIVATE GlanceUpdateImageResponseImagetype
	SHARED  GlanceUpdateImageResponseImagetype
	MARKET  GlanceUpdateImageResponseImagetype
}

func GetGlanceUpdateImageResponseImagetypeEnum() GlanceUpdateImageResponseImagetypeEnum {
	return GlanceUpdateImageResponseImagetypeEnum{
		GOLD: GlanceUpdateImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: GlanceUpdateImageResponseImagetype{
			value: "private",
		},
		SHARED: GlanceUpdateImageResponseImagetype{
			value: "shared",
		},
		MARKET: GlanceUpdateImageResponseImagetype{
			value: "market",
		},
	}
}

func (c GlanceUpdateImageResponseImagetype) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseIsregistered struct {
	value string
}

type GlanceUpdateImageResponseIsregisteredEnum struct {
	TRUE  GlanceUpdateImageResponseIsregistered
	FALSE GlanceUpdateImageResponseIsregistered
}

func GetGlanceUpdateImageResponseIsregisteredEnum() GlanceUpdateImageResponseIsregisteredEnum {
	return GlanceUpdateImageResponseIsregisteredEnum{
		TRUE: GlanceUpdateImageResponseIsregistered{
			value: "true",
		},
		FALSE: GlanceUpdateImageResponseIsregistered{
			value: "false",
		},
	}
}

func (c GlanceUpdateImageResponseIsregistered) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseIsregistered) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseOsBit struct {
	value string
}

type GlanceUpdateImageResponseOsBitEnum struct {
	E_32 GlanceUpdateImageResponseOsBit
	E_64 GlanceUpdateImageResponseOsBit
}

func GetGlanceUpdateImageResponseOsBitEnum() GlanceUpdateImageResponseOsBitEnum {
	return GlanceUpdateImageResponseOsBitEnum{
		E_32: GlanceUpdateImageResponseOsBit{
			value: "32",
		},
		E_64: GlanceUpdateImageResponseOsBit{
			value: "64",
		},
	}
}

func (c GlanceUpdateImageResponseOsBit) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseOsType struct {
	value string
}

type GlanceUpdateImageResponseOsTypeEnum struct {
	LINUX   GlanceUpdateImageResponseOsType
	WINDOWS GlanceUpdateImageResponseOsType
	OTHER   GlanceUpdateImageResponseOsType
}

func GetGlanceUpdateImageResponseOsTypeEnum() GlanceUpdateImageResponseOsTypeEnum {
	return GlanceUpdateImageResponseOsTypeEnum{
		LINUX: GlanceUpdateImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: GlanceUpdateImageResponseOsType{
			value: "Windows",
		},
		OTHER: GlanceUpdateImageResponseOsType{
			value: "Other",
		},
	}
}

func (c GlanceUpdateImageResponseOsType) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseOsType) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponsePlatform struct {
	value string
}

type GlanceUpdateImageResponsePlatformEnum struct {
	WINDOWS      GlanceUpdateImageResponsePlatform
	UBUNTU       GlanceUpdateImageResponsePlatform
	RED_HAT      GlanceUpdateImageResponsePlatform
	SUSE         GlanceUpdateImageResponsePlatform
	CENT_OS      GlanceUpdateImageResponsePlatform
	DEBIAN       GlanceUpdateImageResponsePlatform
	OPEN_SUSE    GlanceUpdateImageResponsePlatform
	ORACLE_LINUX GlanceUpdateImageResponsePlatform
	FEDORA       GlanceUpdateImageResponsePlatform
	OTHER        GlanceUpdateImageResponsePlatform
	CORE_OS      GlanceUpdateImageResponsePlatform
	EULER_OS     GlanceUpdateImageResponsePlatform
}

func GetGlanceUpdateImageResponsePlatformEnum() GlanceUpdateImageResponsePlatformEnum {
	return GlanceUpdateImageResponsePlatformEnum{
		WINDOWS: GlanceUpdateImageResponsePlatform{
			value: "Windows",
		},
		UBUNTU: GlanceUpdateImageResponsePlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceUpdateImageResponsePlatform{
			value: "RedHat",
		},
		SUSE: GlanceUpdateImageResponsePlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceUpdateImageResponsePlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceUpdateImageResponsePlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceUpdateImageResponsePlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceUpdateImageResponsePlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceUpdateImageResponsePlatform{
			value: "Fedora",
		},
		OTHER: GlanceUpdateImageResponsePlatform{
			value: "Other",
		},
		CORE_OS: GlanceUpdateImageResponsePlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceUpdateImageResponsePlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceUpdateImageResponsePlatform) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponsePlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponsePlatform) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseDiskFormat struct {
	value string
}

type GlanceUpdateImageResponseDiskFormatEnum struct {
	VHD   GlanceUpdateImageResponseDiskFormat
	ZVHD  GlanceUpdateImageResponseDiskFormat
	RAW   GlanceUpdateImageResponseDiskFormat
	QCOW2 GlanceUpdateImageResponseDiskFormat
	ZVHD2 GlanceUpdateImageResponseDiskFormat
}

func GetGlanceUpdateImageResponseDiskFormatEnum() GlanceUpdateImageResponseDiskFormatEnum {
	return GlanceUpdateImageResponseDiskFormatEnum{
		VHD: GlanceUpdateImageResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceUpdateImageResponseDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceUpdateImageResponseDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceUpdateImageResponseDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceUpdateImageResponseDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceUpdateImageResponseDiskFormat) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseStatus struct {
	value string
}

type GlanceUpdateImageResponseStatusEnum struct {
	QUEUED  GlanceUpdateImageResponseStatus
	SAVING  GlanceUpdateImageResponseStatus
	DELETED GlanceUpdateImageResponseStatus
	KILLED  GlanceUpdateImageResponseStatus
	ACTIVE  GlanceUpdateImageResponseStatus
}

func GetGlanceUpdateImageResponseStatusEnum() GlanceUpdateImageResponseStatusEnum {
	return GlanceUpdateImageResponseStatusEnum{
		QUEUED: GlanceUpdateImageResponseStatus{
			value: "queued",
		},
		SAVING: GlanceUpdateImageResponseStatus{
			value: "saving",
		},
		DELETED: GlanceUpdateImageResponseStatus{
			value: "deleted",
		},
		KILLED: GlanceUpdateImageResponseStatus{
			value: "killed",
		},
		ACTIVE: GlanceUpdateImageResponseStatus{
			value: "active",
		},
	}
}

func (c GlanceUpdateImageResponseStatus) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseVirtualEnvType struct {
	value string
}

type GlanceUpdateImageResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceUpdateImageResponseVirtualEnvType
	IRONIC         GlanceUpdateImageResponseVirtualEnvType
	DATA_IMAGE     GlanceUpdateImageResponseVirtualEnvType
}

func GetGlanceUpdateImageResponseVirtualEnvTypeEnum() GlanceUpdateImageResponseVirtualEnvTypeEnum {
	return GlanceUpdateImageResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceUpdateImageResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceUpdateImageResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceUpdateImageResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceUpdateImageResponseVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseVisibility struct {
	value string
}

type GlanceUpdateImageResponseVisibilityEnum struct {
	PRIVATE GlanceUpdateImageResponseVisibility
	PUBLIC  GlanceUpdateImageResponseVisibility
	SHARED  GlanceUpdateImageResponseVisibility
}

func GetGlanceUpdateImageResponseVisibilityEnum() GlanceUpdateImageResponseVisibilityEnum {
	return GlanceUpdateImageResponseVisibilityEnum{
		PRIVATE: GlanceUpdateImageResponseVisibility{
			value: "private",
		},
		PUBLIC: GlanceUpdateImageResponseVisibility{
			value: "public",
		},
		SHARED: GlanceUpdateImageResponseVisibility{
			value: "shared",
		},
	}
}

func (c GlanceUpdateImageResponseVisibility) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseVisibility) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseSupportFcInject struct {
	value string
}

type GlanceUpdateImageResponseSupportFcInjectEnum struct {
	TRUE  GlanceUpdateImageResponseSupportFcInject
	FALSE GlanceUpdateImageResponseSupportFcInject
}

func GetGlanceUpdateImageResponseSupportFcInjectEnum() GlanceUpdateImageResponseSupportFcInjectEnum {
	return GlanceUpdateImageResponseSupportFcInjectEnum{
		TRUE: GlanceUpdateImageResponseSupportFcInject{
			value: "true",
		},
		FALSE: GlanceUpdateImageResponseSupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceUpdateImageResponseSupportFcInject) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseSupportFcInject) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseHwFirmwareType struct {
	value string
}

type GlanceUpdateImageResponseHwFirmwareTypeEnum struct {
	BIOS GlanceUpdateImageResponseHwFirmwareType
	UEFI GlanceUpdateImageResponseHwFirmwareType
}

func GetGlanceUpdateImageResponseHwFirmwareTypeEnum() GlanceUpdateImageResponseHwFirmwareTypeEnum {
	return GlanceUpdateImageResponseHwFirmwareTypeEnum{
		BIOS: GlanceUpdateImageResponseHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceUpdateImageResponseHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceUpdateImageResponseHwFirmwareType) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseSupportArm struct {
	value string
}

type GlanceUpdateImageResponseSupportArmEnum struct {
	TRUE  GlanceUpdateImageResponseSupportArm
	FALSE GlanceUpdateImageResponseSupportArm
}

func GetGlanceUpdateImageResponseSupportArmEnum() GlanceUpdateImageResponseSupportArmEnum {
	return GlanceUpdateImageResponseSupportArmEnum{
		TRUE: GlanceUpdateImageResponseSupportArm{
			value: "true",
		},
		FALSE: GlanceUpdateImageResponseSupportArm{
			value: "false",
		},
	}
}

func (c GlanceUpdateImageResponseSupportArm) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseSupportArm) UnmarshalJSON(b []byte) error {
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

type GlanceUpdateImageResponseIsOffshelved struct {
	value string
}

type GlanceUpdateImageResponseIsOffshelvedEnum struct {
	TRUE  GlanceUpdateImageResponseIsOffshelved
	FALSE GlanceUpdateImageResponseIsOffshelved
}

func GetGlanceUpdateImageResponseIsOffshelvedEnum() GlanceUpdateImageResponseIsOffshelvedEnum {
	return GlanceUpdateImageResponseIsOffshelvedEnum{
		TRUE: GlanceUpdateImageResponseIsOffshelved{
			value: "true",
		},
		FALSE: GlanceUpdateImageResponseIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceUpdateImageResponseIsOffshelved) Value() string {
	return c.value
}

func (c GlanceUpdateImageResponseIsOffshelved) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageResponseIsOffshelved) UnmarshalJSON(b []byte) error {
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
