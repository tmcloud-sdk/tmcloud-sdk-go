package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ImageInfo struct {
	BackupId *string `json:"__backup_id,omitempty"`

	DataOrigin *string `json:"__data_origin,omitempty"`

	Description *string `json:"__description,omitempty"`

	ImageSize string `json:"__image_size"`

	ImageSourceType ImageInfoImageSourceType `json:"__image_source_type"`

	Imagetype ImageInfoImagetype `json:"__imagetype"`

	Isregistered ImageInfoIsregistered `json:"__isregistered"`

	Originalimagename *string `json:"__originalimagename,omitempty"`

	OsBit *ImageInfoOsBit `json:"__os_bit,omitempty"`

	OsType ImageInfoOsType `json:"__os_type"`

	OsVersion *string `json:"__os_version,omitempty"`

	Platform *ImageInfoPlatform `json:"__platform,omitempty"`

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

	SystemSupportMarket *bool `json:"__system_support_market,omitempty"`

	Checksum *string `json:"checksum,omitempty"`

	ContainerFormat string `json:"container_format"`

	CreatedAt string `json:"created_at"`

	DiskFormat *string `json:"disk_format,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	File *string `json:"file,omitempty"`

	Id string `json:"id"`

	MinDisk int32 `json:"min_disk"`

	MinRam int32 `json:"min_ram"`

	Name string `json:"name"`

	Owner string `json:"owner"`

	Protected bool `json:"protected"`

	Schema *string `json:"schema,omitempty"`

	Self string `json:"self"`

	Size *int32 `json:"size,omitempty"`

	Status ImageInfoStatus `json:"status"`

	Tags []string `json:"tags"`

	UpdatedAt string `json:"updated_at"`

	VirtualEnvType ImageInfoVirtualEnvType `json:"virtual_env_type"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Visibility ImageInfoVisibility `json:"visibility"`

	SupportFcInject *ImageInfoSupportFcInject `json:"__support_fc_inject,omitempty"`

	HwFirmwareType *ImageInfoHwFirmwareType `json:"hw_firmware_type,omitempty"`

	SupportArm *ImageInfoSupportArm `json:"__support_arm,omitempty"`

	MaxRam *string `json:"max_ram,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	OsFeatureList *string `json:"__os_feature_list,omitempty"`

	AccountCode *string `json:"__account_code,omitempty"`

	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty"`

	IsOffshelved *string `json:"__is_offshelved,omitempty"`

	Lazyloading *string `json:"__lazyloading,omitempty"`

	RootOrigin *string `json:"__root_origin,omitempty"`

	SequenceNum *string `json:"__sequence_num,omitempty"`

	ActiveAt string `json:"active_at"`

	SupportAgentList *string `json:"__support_agent_list,omitempty"`

	SupportAmd *string `json:"__support_amd,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}

type ImageInfoImageSourceType struct {
	value string
}

type ImageInfoImageSourceTypeEnum struct {
	UDS   ImageInfoImageSourceType
	SWIFT ImageInfoImageSourceType
}

func GetImageInfoImageSourceTypeEnum() ImageInfoImageSourceTypeEnum {
	return ImageInfoImageSourceTypeEnum{
		UDS: ImageInfoImageSourceType{
			value: "uds",
		},
		SWIFT: ImageInfoImageSourceType{
			value: "swift",
		},
	}
}

func (c ImageInfoImageSourceType) Value() string {
	return c.value
}

func (c ImageInfoImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoImageSourceType) UnmarshalJSON(b []byte) error {
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

type ImageInfoImagetype struct {
	value string
}

type ImageInfoImagetypeEnum struct {
	GOLD    ImageInfoImagetype
	PRIVATE ImageInfoImagetype
	SHARED  ImageInfoImagetype
}

func GetImageInfoImagetypeEnum() ImageInfoImagetypeEnum {
	return ImageInfoImagetypeEnum{
		GOLD: ImageInfoImagetype{
			value: "gold",
		},
		PRIVATE: ImageInfoImagetype{
			value: "private",
		},
		SHARED: ImageInfoImagetype{
			value: "shared",
		},
	}
}

func (c ImageInfoImagetype) Value() string {
	return c.value
}

func (c ImageInfoImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoImagetype) UnmarshalJSON(b []byte) error {
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

type ImageInfoIsregistered struct {
	value string
}

type ImageInfoIsregisteredEnum struct {
	TRUE  ImageInfoIsregistered
	FALSE ImageInfoIsregistered
}

func GetImageInfoIsregisteredEnum() ImageInfoIsregisteredEnum {
	return ImageInfoIsregisteredEnum{
		TRUE: ImageInfoIsregistered{
			value: "true",
		},
		FALSE: ImageInfoIsregistered{
			value: "false",
		},
	}
}

func (c ImageInfoIsregistered) Value() string {
	return c.value
}

func (c ImageInfoIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoIsregistered) UnmarshalJSON(b []byte) error {
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

type ImageInfoOsBit struct {
	value string
}

type ImageInfoOsBitEnum struct {
	E_32 ImageInfoOsBit
	E_64 ImageInfoOsBit
}

func GetImageInfoOsBitEnum() ImageInfoOsBitEnum {
	return ImageInfoOsBitEnum{
		E_32: ImageInfoOsBit{
			value: "32",
		},
		E_64: ImageInfoOsBit{
			value: "64",
		},
	}
}

func (c ImageInfoOsBit) Value() string {
	return c.value
}

func (c ImageInfoOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoOsBit) UnmarshalJSON(b []byte) error {
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

type ImageInfoOsType struct {
	value string
}

type ImageInfoOsTypeEnum struct {
	LINUX   ImageInfoOsType
	WINDOWS ImageInfoOsType
	OTHER   ImageInfoOsType
}

func GetImageInfoOsTypeEnum() ImageInfoOsTypeEnum {
	return ImageInfoOsTypeEnum{
		LINUX: ImageInfoOsType{
			value: "Linux",
		},
		WINDOWS: ImageInfoOsType{
			value: "Windows",
		},
		OTHER: ImageInfoOsType{
			value: "Other",
		},
	}
}

func (c ImageInfoOsType) Value() string {
	return c.value
}

func (c ImageInfoOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoOsType) UnmarshalJSON(b []byte) error {
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

type ImageInfoPlatform struct {
	value string
}

type ImageInfoPlatformEnum struct {
	WINDOWS      ImageInfoPlatform
	UBUNTU       ImageInfoPlatform
	RED_HAT      ImageInfoPlatform
	SUSE         ImageInfoPlatform
	CENT_OS      ImageInfoPlatform
	DEBIAN       ImageInfoPlatform
	OPEN_SUSE    ImageInfoPlatform
	ORACLE_LINUX ImageInfoPlatform
	FEDORA       ImageInfoPlatform
	OTHER        ImageInfoPlatform
	CORE_OS      ImageInfoPlatform
	EULER_OS     ImageInfoPlatform
}

func GetImageInfoPlatformEnum() ImageInfoPlatformEnum {
	return ImageInfoPlatformEnum{
		WINDOWS: ImageInfoPlatform{
			value: "Windows",
		},
		UBUNTU: ImageInfoPlatform{
			value: "Ubuntu",
		},
		RED_HAT: ImageInfoPlatform{
			value: "RedHat",
		},
		SUSE: ImageInfoPlatform{
			value: "SUSE",
		},
		CENT_OS: ImageInfoPlatform{
			value: "CentOS",
		},
		DEBIAN: ImageInfoPlatform{
			value: "Debian",
		},
		OPEN_SUSE: ImageInfoPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: ImageInfoPlatform{
			value: "Oracle Linux",
		},
		FEDORA: ImageInfoPlatform{
			value: "Fedora",
		},
		OTHER: ImageInfoPlatform{
			value: "Other",
		},
		CORE_OS: ImageInfoPlatform{
			value: "CoreOS",
		},
		EULER_OS: ImageInfoPlatform{
			value: "EulerOS",
		},
	}
}

func (c ImageInfoPlatform) Value() string {
	return c.value
}

func (c ImageInfoPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoPlatform) UnmarshalJSON(b []byte) error {
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

type ImageInfoStatus struct {
	value string
}

type ImageInfoStatusEnum struct {
	QUEUED  ImageInfoStatus
	SAVING  ImageInfoStatus
	DELETED ImageInfoStatus
	KILLED  ImageInfoStatus
	ACTIVE  ImageInfoStatus
}

func GetImageInfoStatusEnum() ImageInfoStatusEnum {
	return ImageInfoStatusEnum{
		QUEUED: ImageInfoStatus{
			value: "queued",
		},
		SAVING: ImageInfoStatus{
			value: "saving",
		},
		DELETED: ImageInfoStatus{
			value: "deleted",
		},
		KILLED: ImageInfoStatus{
			value: "killed",
		},
		ACTIVE: ImageInfoStatus{
			value: "active",
		},
	}
}

func (c ImageInfoStatus) Value() string {
	return c.value
}

func (c ImageInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoStatus) UnmarshalJSON(b []byte) error {
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

type ImageInfoVirtualEnvType struct {
	value string
}

type ImageInfoVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ImageInfoVirtualEnvType
	IRONIC         ImageInfoVirtualEnvType
	DATA_IMAGE     ImageInfoVirtualEnvType
}

func GetImageInfoVirtualEnvTypeEnum() ImageInfoVirtualEnvTypeEnum {
	return ImageInfoVirtualEnvTypeEnum{
		FUSION_COMPUTE: ImageInfoVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ImageInfoVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ImageInfoVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ImageInfoVirtualEnvType) Value() string {
	return c.value
}

func (c ImageInfoVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ImageInfoVisibility struct {
	value string
}

type ImageInfoVisibilityEnum struct {
	PRIVATE ImageInfoVisibility
	PUBLIC  ImageInfoVisibility
}

func GetImageInfoVisibilityEnum() ImageInfoVisibilityEnum {
	return ImageInfoVisibilityEnum{
		PRIVATE: ImageInfoVisibility{
			value: "private",
		},
		PUBLIC: ImageInfoVisibility{
			value: "public",
		},
	}
}

func (c ImageInfoVisibility) Value() string {
	return c.value
}

func (c ImageInfoVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoVisibility) UnmarshalJSON(b []byte) error {
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

type ImageInfoSupportFcInject struct {
	value string
}

type ImageInfoSupportFcInjectEnum struct {
	TRUE  ImageInfoSupportFcInject
	FALSE ImageInfoSupportFcInject
}

func GetImageInfoSupportFcInjectEnum() ImageInfoSupportFcInjectEnum {
	return ImageInfoSupportFcInjectEnum{
		TRUE: ImageInfoSupportFcInject{
			value: "true",
		},
		FALSE: ImageInfoSupportFcInject{
			value: "false",
		},
	}
}

func (c ImageInfoSupportFcInject) Value() string {
	return c.value
}

func (c ImageInfoSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoSupportFcInject) UnmarshalJSON(b []byte) error {
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

type ImageInfoHwFirmwareType struct {
	value string
}

type ImageInfoHwFirmwareTypeEnum struct {
	BIOS ImageInfoHwFirmwareType
	UEFI ImageInfoHwFirmwareType
}

func GetImageInfoHwFirmwareTypeEnum() ImageInfoHwFirmwareTypeEnum {
	return ImageInfoHwFirmwareTypeEnum{
		BIOS: ImageInfoHwFirmwareType{
			value: "bios",
		},
		UEFI: ImageInfoHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c ImageInfoHwFirmwareType) Value() string {
	return c.value
}

func (c ImageInfoHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type ImageInfoSupportArm struct {
	value string
}

type ImageInfoSupportArmEnum struct {
	TRUE  ImageInfoSupportArm
	FALSE ImageInfoSupportArm
}

func GetImageInfoSupportArmEnum() ImageInfoSupportArmEnum {
	return ImageInfoSupportArmEnum{
		TRUE: ImageInfoSupportArm{
			value: "true",
		},
		FALSE: ImageInfoSupportArm{
			value: "false",
		},
	}
}

func (c ImageInfoSupportArm) Value() string {
	return c.value
}

func (c ImageInfoSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageInfoSupportArm) UnmarshalJSON(b []byte) error {
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
