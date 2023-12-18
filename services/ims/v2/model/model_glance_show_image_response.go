package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceShowImageResponse struct {
	BackupId *string `json:"__backup_id,omitempty"`

	DataOrigin *string `json:"__data_origin,omitempty"`

	Description *string `json:"__description,omitempty"`

	ImageSize *string `json:"__image_size,omitempty"`

	ImageSourceType *GlanceShowImageResponseImageSourceType `json:"__image_source_type,omitempty"`

	Imagetype *GlanceShowImageResponseImagetype `json:"__imagetype,omitempty"`

	Isregistered *GlanceShowImageResponseIsregistered `json:"__isregistered,omitempty"`

	Originalimagename *string `json:"__originalimagename,omitempty"`

	OsBit *GlanceShowImageResponseOsBit `json:"__os_bit,omitempty"`

	OsType *GlanceShowImageResponseOsType `json:"__os_type,omitempty"`

	OsVersion *string `json:"__os_version,omitempty"`

	Platform *GlanceShowImageResponsePlatform `json:"__platform,omitempty"`

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

	DiskFormat *GlanceShowImageResponseDiskFormat `json:"disk_format,omitempty"`

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

	Status *GlanceShowImageResponseStatus `json:"status,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	VirtualEnvType *GlanceShowImageResponseVirtualEnvType `json:"virtual_env_type,omitempty"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Visibility *GlanceShowImageResponseVisibility `json:"visibility,omitempty"`

	SupportFcInject *GlanceShowImageResponseSupportFcInject `json:"__support_fc_inject,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HwFirmwareType *GlanceShowImageResponseHwFirmwareType `json:"hw_firmware_type,omitempty"`

	SupportArm *GlanceShowImageResponseSupportArm `json:"__support_arm,omitempty"`

	IsOffshelved *GlanceShowImageResponseIsOffshelved `json:"__is_offshelved,omitempty"`

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

func (o GlanceShowImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageResponse", string(data)}, " ")
}

type GlanceShowImageResponseImageSourceType struct {
	value string
}

type GlanceShowImageResponseImageSourceTypeEnum struct {
	UDS   GlanceShowImageResponseImageSourceType
	SWIFT GlanceShowImageResponseImageSourceType
}

func GetGlanceShowImageResponseImageSourceTypeEnum() GlanceShowImageResponseImageSourceTypeEnum {
	return GlanceShowImageResponseImageSourceTypeEnum{
		UDS: GlanceShowImageResponseImageSourceType{
			value: "uds",
		},
		SWIFT: GlanceShowImageResponseImageSourceType{
			value: "swift",
		},
	}
}

func (c GlanceShowImageResponseImageSourceType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseImageSourceType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseImagetype struct {
	value string
}

type GlanceShowImageResponseImagetypeEnum struct {
	GOLD    GlanceShowImageResponseImagetype
	PRIVATE GlanceShowImageResponseImagetype
	SHARED  GlanceShowImageResponseImagetype
	MARKET  GlanceShowImageResponseImagetype
}

func GetGlanceShowImageResponseImagetypeEnum() GlanceShowImageResponseImagetypeEnum {
	return GlanceShowImageResponseImagetypeEnum{
		GOLD: GlanceShowImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: GlanceShowImageResponseImagetype{
			value: "private",
		},
		SHARED: GlanceShowImageResponseImagetype{
			value: "shared",
		},
		MARKET: GlanceShowImageResponseImagetype{
			value: "market",
		},
	}
}

func (c GlanceShowImageResponseImagetype) Value() string {
	return c.value
}

func (c GlanceShowImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseImagetype) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseIsregistered struct {
	value string
}

type GlanceShowImageResponseIsregisteredEnum struct {
	TRUE  GlanceShowImageResponseIsregistered
	FALSE GlanceShowImageResponseIsregistered
}

func GetGlanceShowImageResponseIsregisteredEnum() GlanceShowImageResponseIsregisteredEnum {
	return GlanceShowImageResponseIsregisteredEnum{
		TRUE: GlanceShowImageResponseIsregistered{
			value: "true",
		},
		FALSE: GlanceShowImageResponseIsregistered{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseIsregistered) Value() string {
	return c.value
}

func (c GlanceShowImageResponseIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseIsregistered) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseOsBit struct {
	value string
}

type GlanceShowImageResponseOsBitEnum struct {
	E_32 GlanceShowImageResponseOsBit
	E_64 GlanceShowImageResponseOsBit
}

func GetGlanceShowImageResponseOsBitEnum() GlanceShowImageResponseOsBitEnum {
	return GlanceShowImageResponseOsBitEnum{
		E_32: GlanceShowImageResponseOsBit{
			value: "32",
		},
		E_64: GlanceShowImageResponseOsBit{
			value: "64",
		},
	}
}

func (c GlanceShowImageResponseOsBit) Value() string {
	return c.value
}

func (c GlanceShowImageResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseOsBit) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseOsType struct {
	value string
}

type GlanceShowImageResponseOsTypeEnum struct {
	LINUX   GlanceShowImageResponseOsType
	WINDOWS GlanceShowImageResponseOsType
	OTHER   GlanceShowImageResponseOsType
}

func GetGlanceShowImageResponseOsTypeEnum() GlanceShowImageResponseOsTypeEnum {
	return GlanceShowImageResponseOsTypeEnum{
		LINUX: GlanceShowImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: GlanceShowImageResponseOsType{
			value: "Windows",
		},
		OTHER: GlanceShowImageResponseOsType{
			value: "Other",
		},
	}
}

func (c GlanceShowImageResponseOsType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseOsType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponsePlatform struct {
	value string
}

type GlanceShowImageResponsePlatformEnum struct {
	WINDOWS      GlanceShowImageResponsePlatform
	UBUNTU       GlanceShowImageResponsePlatform
	RED_HAT      GlanceShowImageResponsePlatform
	SUSE         GlanceShowImageResponsePlatform
	CENT_OS      GlanceShowImageResponsePlatform
	DEBIAN       GlanceShowImageResponsePlatform
	OPEN_SUSE    GlanceShowImageResponsePlatform
	ORACLE_LINUX GlanceShowImageResponsePlatform
	FEDORA       GlanceShowImageResponsePlatform
	OTHER        GlanceShowImageResponsePlatform
	CORE_OS      GlanceShowImageResponsePlatform
	EULER_OS     GlanceShowImageResponsePlatform
}

func GetGlanceShowImageResponsePlatformEnum() GlanceShowImageResponsePlatformEnum {
	return GlanceShowImageResponsePlatformEnum{
		WINDOWS: GlanceShowImageResponsePlatform{
			value: "Windows",
		},
		UBUNTU: GlanceShowImageResponsePlatform{
			value: "Ubuntu",
		},
		RED_HAT: GlanceShowImageResponsePlatform{
			value: "RedHat",
		},
		SUSE: GlanceShowImageResponsePlatform{
			value: "SUSE",
		},
		CENT_OS: GlanceShowImageResponsePlatform{
			value: "CentOS",
		},
		DEBIAN: GlanceShowImageResponsePlatform{
			value: "Debian",
		},
		OPEN_SUSE: GlanceShowImageResponsePlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: GlanceShowImageResponsePlatform{
			value: "OracleLinux",
		},
		FEDORA: GlanceShowImageResponsePlatform{
			value: "Fedora",
		},
		OTHER: GlanceShowImageResponsePlatform{
			value: "Other",
		},
		CORE_OS: GlanceShowImageResponsePlatform{
			value: "CoreOS",
		},
		EULER_OS: GlanceShowImageResponsePlatform{
			value: "EulerOS",
		},
	}
}

func (c GlanceShowImageResponsePlatform) Value() string {
	return c.value
}

func (c GlanceShowImageResponsePlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponsePlatform) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseDiskFormat struct {
	value string
}

type GlanceShowImageResponseDiskFormatEnum struct {
	VHD   GlanceShowImageResponseDiskFormat
	ZVHD  GlanceShowImageResponseDiskFormat
	RAW   GlanceShowImageResponseDiskFormat
	QCOW2 GlanceShowImageResponseDiskFormat
	ZVHD2 GlanceShowImageResponseDiskFormat
}

func GetGlanceShowImageResponseDiskFormatEnum() GlanceShowImageResponseDiskFormatEnum {
	return GlanceShowImageResponseDiskFormatEnum{
		VHD: GlanceShowImageResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: GlanceShowImageResponseDiskFormat{
			value: "zvhd",
		},
		RAW: GlanceShowImageResponseDiskFormat{
			value: "raw",
		},
		QCOW2: GlanceShowImageResponseDiskFormat{
			value: "qcow2",
		},
		ZVHD2: GlanceShowImageResponseDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c GlanceShowImageResponseDiskFormat) Value() string {
	return c.value
}

func (c GlanceShowImageResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseStatus struct {
	value string
}

type GlanceShowImageResponseStatusEnum struct {
	QUEUED  GlanceShowImageResponseStatus
	SAVING  GlanceShowImageResponseStatus
	DELETED GlanceShowImageResponseStatus
	KILLED  GlanceShowImageResponseStatus
	ACTIVE  GlanceShowImageResponseStatus
}

func GetGlanceShowImageResponseStatusEnum() GlanceShowImageResponseStatusEnum {
	return GlanceShowImageResponseStatusEnum{
		QUEUED: GlanceShowImageResponseStatus{
			value: "queued",
		},
		SAVING: GlanceShowImageResponseStatus{
			value: "saving",
		},
		DELETED: GlanceShowImageResponseStatus{
			value: "deleted",
		},
		KILLED: GlanceShowImageResponseStatus{
			value: "killed",
		},
		ACTIVE: GlanceShowImageResponseStatus{
			value: "active",
		},
	}
}

func (c GlanceShowImageResponseStatus) Value() string {
	return c.value
}

func (c GlanceShowImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseVirtualEnvType struct {
	value string
}

type GlanceShowImageResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE GlanceShowImageResponseVirtualEnvType
	IRONIC         GlanceShowImageResponseVirtualEnvType
	DATA_IMAGE     GlanceShowImageResponseVirtualEnvType
}

func GetGlanceShowImageResponseVirtualEnvTypeEnum() GlanceShowImageResponseVirtualEnvTypeEnum {
	return GlanceShowImageResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: GlanceShowImageResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: GlanceShowImageResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: GlanceShowImageResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c GlanceShowImageResponseVirtualEnvType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseVisibility struct {
	value string
}

type GlanceShowImageResponseVisibilityEnum struct {
	PRIVATE GlanceShowImageResponseVisibility
	PUBLIC  GlanceShowImageResponseVisibility
	SHARED  GlanceShowImageResponseVisibility
}

func GetGlanceShowImageResponseVisibilityEnum() GlanceShowImageResponseVisibilityEnum {
	return GlanceShowImageResponseVisibilityEnum{
		PRIVATE: GlanceShowImageResponseVisibility{
			value: "private",
		},
		PUBLIC: GlanceShowImageResponseVisibility{
			value: "public",
		},
		SHARED: GlanceShowImageResponseVisibility{
			value: "shared",
		},
	}
}

func (c GlanceShowImageResponseVisibility) Value() string {
	return c.value
}

func (c GlanceShowImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseVisibility) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseSupportFcInject struct {
	value string
}

type GlanceShowImageResponseSupportFcInjectEnum struct {
	TRUE  GlanceShowImageResponseSupportFcInject
	FALSE GlanceShowImageResponseSupportFcInject
}

func GetGlanceShowImageResponseSupportFcInjectEnum() GlanceShowImageResponseSupportFcInjectEnum {
	return GlanceShowImageResponseSupportFcInjectEnum{
		TRUE: GlanceShowImageResponseSupportFcInject{
			value: "true",
		},
		FALSE: GlanceShowImageResponseSupportFcInject{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseSupportFcInject) Value() string {
	return c.value
}

func (c GlanceShowImageResponseSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseSupportFcInject) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseHwFirmwareType struct {
	value string
}

type GlanceShowImageResponseHwFirmwareTypeEnum struct {
	BIOS GlanceShowImageResponseHwFirmwareType
	UEFI GlanceShowImageResponseHwFirmwareType
}

func GetGlanceShowImageResponseHwFirmwareTypeEnum() GlanceShowImageResponseHwFirmwareTypeEnum {
	return GlanceShowImageResponseHwFirmwareTypeEnum{
		BIOS: GlanceShowImageResponseHwFirmwareType{
			value: "bios",
		},
		UEFI: GlanceShowImageResponseHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c GlanceShowImageResponseHwFirmwareType) Value() string {
	return c.value
}

func (c GlanceShowImageResponseHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseSupportArm struct {
	value string
}

type GlanceShowImageResponseSupportArmEnum struct {
	TRUE  GlanceShowImageResponseSupportArm
	FALSE GlanceShowImageResponseSupportArm
}

func GetGlanceShowImageResponseSupportArmEnum() GlanceShowImageResponseSupportArmEnum {
	return GlanceShowImageResponseSupportArmEnum{
		TRUE: GlanceShowImageResponseSupportArm{
			value: "true",
		},
		FALSE: GlanceShowImageResponseSupportArm{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseSupportArm) Value() string {
	return c.value
}

func (c GlanceShowImageResponseSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseSupportArm) UnmarshalJSON(b []byte) error {
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

type GlanceShowImageResponseIsOffshelved struct {
	value string
}

type GlanceShowImageResponseIsOffshelvedEnum struct {
	TRUE  GlanceShowImageResponseIsOffshelved
	FALSE GlanceShowImageResponseIsOffshelved
}

func GetGlanceShowImageResponseIsOffshelvedEnum() GlanceShowImageResponseIsOffshelvedEnum {
	return GlanceShowImageResponseIsOffshelvedEnum{
		TRUE: GlanceShowImageResponseIsOffshelved{
			value: "true",
		},
		FALSE: GlanceShowImageResponseIsOffshelved{
			value: "false",
		},
	}
}

func (c GlanceShowImageResponseIsOffshelved) Value() string {
	return c.value
}

func (c GlanceShowImageResponseIsOffshelved) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceShowImageResponseIsOffshelved) UnmarshalJSON(b []byte) error {
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
