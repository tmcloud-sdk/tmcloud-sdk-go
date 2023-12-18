package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateImageResponse struct {
	BackupId *string `json:"__backup_id,omitempty"`

	DataOrigin *string `json:"__data_origin,omitempty"`

	Description *string `json:"__description,omitempty"`

	ImageSize *string `json:"__image_size,omitempty"`

	ImageSourceType *UpdateImageResponseImageSourceType `json:"__image_source_type,omitempty"`

	Imagetype *UpdateImageResponseImagetype `json:"__imagetype,omitempty"`

	Isregistered *UpdateImageResponseIsregistered `json:"__isregistered,omitempty"`

	Originalimagename *string `json:"__originalimagename,omitempty"`

	OsBit *UpdateImageResponseOsBit `json:"__os_bit,omitempty"`

	OsType *UpdateImageResponseOsType `json:"__os_type,omitempty"`

	OsVersion *string `json:"__os_version,omitempty"`

	Platform *UpdateImageResponsePlatform `json:"__platform,omitempty"`

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

	ContainerFormat *string `json:"container_format,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	DiskFormat *string `json:"disk_format,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	File *string `json:"file,omitempty"`

	Id *string `json:"id,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	Schema *string `json:"schema,omitempty"`

	Self *string `json:"self,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Status *UpdateImageResponseStatus `json:"status,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	VirtualEnvType *UpdateImageResponseVirtualEnvType `json:"virtual_env_type,omitempty"`

	VirtualSize *int32 `json:"virtual_size,omitempty"`

	Visibility *UpdateImageResponseVisibility `json:"visibility,omitempty"`

	SupportFcInject *UpdateImageResponseSupportFcInject `json:"__support_fc_inject,omitempty"`

	HwFirmwareType *UpdateImageResponseHwFirmwareType `json:"hw_firmware_type,omitempty"`

	SupportArm *UpdateImageResponseSupportArm `json:"__support_arm,omitempty"`

	MaxRam *string `json:"max_ram,omitempty"`

	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	OsFeatureList *string `json:"__os_feature_list,omitempty"`

	AccountCode *string `json:"__account_code,omitempty"`

	HwVifMultiqueueEnabled *string `json:"hw_vif_multiqueue_enabled,omitempty"`

	IsOffshelved *string `json:"__is_offshelved,omitempty"`

	Lazyloading *string `json:"__lazyloading,omitempty"`

	RootOrigin *string `json:"__root_origin,omitempty"`

	SequenceNum *string `json:"__sequence_num,omitempty"`

	ActiveAt *string `json:"active_at,omitempty"`

	SupportAgentList *string `json:"__support_agent_list,omitempty"`

	SupportAmd     *string `json:"__support_amd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageResponse struct{}"
	}

	return strings.Join([]string{"UpdateImageResponse", string(data)}, " ")
}

type UpdateImageResponseImageSourceType struct {
	value string
}

type UpdateImageResponseImageSourceTypeEnum struct {
	UDS   UpdateImageResponseImageSourceType
	SWIFT UpdateImageResponseImageSourceType
}

func GetUpdateImageResponseImageSourceTypeEnum() UpdateImageResponseImageSourceTypeEnum {
	return UpdateImageResponseImageSourceTypeEnum{
		UDS: UpdateImageResponseImageSourceType{
			value: "uds",
		},
		SWIFT: UpdateImageResponseImageSourceType{
			value: "swift",
		},
	}
}

func (c UpdateImageResponseImageSourceType) Value() string {
	return c.value
}

func (c UpdateImageResponseImageSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseImageSourceType) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseImagetype struct {
	value string
}

type UpdateImageResponseImagetypeEnum struct {
	GOLD    UpdateImageResponseImagetype
	PRIVATE UpdateImageResponseImagetype
	SHARED  UpdateImageResponseImagetype
}

func GetUpdateImageResponseImagetypeEnum() UpdateImageResponseImagetypeEnum {
	return UpdateImageResponseImagetypeEnum{
		GOLD: UpdateImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: UpdateImageResponseImagetype{
			value: "private",
		},
		SHARED: UpdateImageResponseImagetype{
			value: "shared",
		},
	}
}

func (c UpdateImageResponseImagetype) Value() string {
	return c.value
}

func (c UpdateImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseImagetype) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseIsregistered struct {
	value string
}

type UpdateImageResponseIsregisteredEnum struct {
	TRUE  UpdateImageResponseIsregistered
	FALSE UpdateImageResponseIsregistered
}

func GetUpdateImageResponseIsregisteredEnum() UpdateImageResponseIsregisteredEnum {
	return UpdateImageResponseIsregisteredEnum{
		TRUE: UpdateImageResponseIsregistered{
			value: "true",
		},
		FALSE: UpdateImageResponseIsregistered{
			value: "false",
		},
	}
}

func (c UpdateImageResponseIsregistered) Value() string {
	return c.value
}

func (c UpdateImageResponseIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseIsregistered) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseOsBit struct {
	value string
}

type UpdateImageResponseOsBitEnum struct {
	E_32 UpdateImageResponseOsBit
	E_64 UpdateImageResponseOsBit
}

func GetUpdateImageResponseOsBitEnum() UpdateImageResponseOsBitEnum {
	return UpdateImageResponseOsBitEnum{
		E_32: UpdateImageResponseOsBit{
			value: "32",
		},
		E_64: UpdateImageResponseOsBit{
			value: "64",
		},
	}
}

func (c UpdateImageResponseOsBit) Value() string {
	return c.value
}

func (c UpdateImageResponseOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseOsBit) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseOsType struct {
	value string
}

type UpdateImageResponseOsTypeEnum struct {
	LINUX   UpdateImageResponseOsType
	WINDOWS UpdateImageResponseOsType
	OTHER   UpdateImageResponseOsType
}

func GetUpdateImageResponseOsTypeEnum() UpdateImageResponseOsTypeEnum {
	return UpdateImageResponseOsTypeEnum{
		LINUX: UpdateImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: UpdateImageResponseOsType{
			value: "Windows",
		},
		OTHER: UpdateImageResponseOsType{
			value: "Other",
		},
	}
}

func (c UpdateImageResponseOsType) Value() string {
	return c.value
}

func (c UpdateImageResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseOsType) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponsePlatform struct {
	value string
}

type UpdateImageResponsePlatformEnum struct {
	WINDOWS      UpdateImageResponsePlatform
	UBUNTU       UpdateImageResponsePlatform
	RED_HAT      UpdateImageResponsePlatform
	SUSE         UpdateImageResponsePlatform
	CENT_OS      UpdateImageResponsePlatform
	DEBIAN       UpdateImageResponsePlatform
	OPEN_SUSE    UpdateImageResponsePlatform
	ORACLE_LINUX UpdateImageResponsePlatform
	FEDORA       UpdateImageResponsePlatform
	OTHER        UpdateImageResponsePlatform
	CORE_OS      UpdateImageResponsePlatform
	EULER_OS     UpdateImageResponsePlatform
}

func GetUpdateImageResponsePlatformEnum() UpdateImageResponsePlatformEnum {
	return UpdateImageResponsePlatformEnum{
		WINDOWS: UpdateImageResponsePlatform{
			value: "Windows",
		},
		UBUNTU: UpdateImageResponsePlatform{
			value: "Ubuntu",
		},
		RED_HAT: UpdateImageResponsePlatform{
			value: "RedHat",
		},
		SUSE: UpdateImageResponsePlatform{
			value: "SUSE",
		},
		CENT_OS: UpdateImageResponsePlatform{
			value: "CentOS",
		},
		DEBIAN: UpdateImageResponsePlatform{
			value: "Debian",
		},
		OPEN_SUSE: UpdateImageResponsePlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: UpdateImageResponsePlatform{
			value: "Oracle Linux",
		},
		FEDORA: UpdateImageResponsePlatform{
			value: "Fedora",
		},
		OTHER: UpdateImageResponsePlatform{
			value: "Other",
		},
		CORE_OS: UpdateImageResponsePlatform{
			value: "CoreOS",
		},
		EULER_OS: UpdateImageResponsePlatform{
			value: "EulerOS",
		},
	}
}

func (c UpdateImageResponsePlatform) Value() string {
	return c.value
}

func (c UpdateImageResponsePlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponsePlatform) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseStatus struct {
	value string
}

type UpdateImageResponseStatusEnum struct {
	QUEUED  UpdateImageResponseStatus
	SAVING  UpdateImageResponseStatus
	DELETED UpdateImageResponseStatus
	KILLED  UpdateImageResponseStatus
	ACTIVE  UpdateImageResponseStatus
}

func GetUpdateImageResponseStatusEnum() UpdateImageResponseStatusEnum {
	return UpdateImageResponseStatusEnum{
		QUEUED: UpdateImageResponseStatus{
			value: "queued",
		},
		SAVING: UpdateImageResponseStatus{
			value: "saving",
		},
		DELETED: UpdateImageResponseStatus{
			value: "deleted",
		},
		KILLED: UpdateImageResponseStatus{
			value: "killed",
		},
		ACTIVE: UpdateImageResponseStatus{
			value: "active",
		},
	}
}

func (c UpdateImageResponseStatus) Value() string {
	return c.value
}

func (c UpdateImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseVirtualEnvType struct {
	value string
}

type UpdateImageResponseVirtualEnvTypeEnum struct {
	FUSION_COMPUTE UpdateImageResponseVirtualEnvType
	IRONIC         UpdateImageResponseVirtualEnvType
	DATA_IMAGE     UpdateImageResponseVirtualEnvType
}

func GetUpdateImageResponseVirtualEnvTypeEnum() UpdateImageResponseVirtualEnvTypeEnum {
	return UpdateImageResponseVirtualEnvTypeEnum{
		FUSION_COMPUTE: UpdateImageResponseVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: UpdateImageResponseVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: UpdateImageResponseVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c UpdateImageResponseVirtualEnvType) Value() string {
	return c.value
}

func (c UpdateImageResponseVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseVisibility struct {
	value string
}

type UpdateImageResponseVisibilityEnum struct {
	PRIVATE UpdateImageResponseVisibility
	PUBLIC  UpdateImageResponseVisibility
}

func GetUpdateImageResponseVisibilityEnum() UpdateImageResponseVisibilityEnum {
	return UpdateImageResponseVisibilityEnum{
		PRIVATE: UpdateImageResponseVisibility{
			value: "private",
		},
		PUBLIC: UpdateImageResponseVisibility{
			value: "public",
		},
	}
}

func (c UpdateImageResponseVisibility) Value() string {
	return c.value
}

func (c UpdateImageResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseVisibility) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseSupportFcInject struct {
	value string
}

type UpdateImageResponseSupportFcInjectEnum struct {
	TRUE  UpdateImageResponseSupportFcInject
	FALSE UpdateImageResponseSupportFcInject
}

func GetUpdateImageResponseSupportFcInjectEnum() UpdateImageResponseSupportFcInjectEnum {
	return UpdateImageResponseSupportFcInjectEnum{
		TRUE: UpdateImageResponseSupportFcInject{
			value: "true",
		},
		FALSE: UpdateImageResponseSupportFcInject{
			value: "false",
		},
	}
}

func (c UpdateImageResponseSupportFcInject) Value() string {
	return c.value
}

func (c UpdateImageResponseSupportFcInject) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseSupportFcInject) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseHwFirmwareType struct {
	value string
}

type UpdateImageResponseHwFirmwareTypeEnum struct {
	BIOS UpdateImageResponseHwFirmwareType
	UEFI UpdateImageResponseHwFirmwareType
}

func GetUpdateImageResponseHwFirmwareTypeEnum() UpdateImageResponseHwFirmwareTypeEnum {
	return UpdateImageResponseHwFirmwareTypeEnum{
		BIOS: UpdateImageResponseHwFirmwareType{
			value: "bios",
		},
		UEFI: UpdateImageResponseHwFirmwareType{
			value: "uefi",
		},
	}
}

func (c UpdateImageResponseHwFirmwareType) Value() string {
	return c.value
}

func (c UpdateImageResponseHwFirmwareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseHwFirmwareType) UnmarshalJSON(b []byte) error {
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

type UpdateImageResponseSupportArm struct {
	value string
}

type UpdateImageResponseSupportArmEnum struct {
	TRUE  UpdateImageResponseSupportArm
	FALSE UpdateImageResponseSupportArm
}

func GetUpdateImageResponseSupportArmEnum() UpdateImageResponseSupportArmEnum {
	return UpdateImageResponseSupportArmEnum{
		TRUE: UpdateImageResponseSupportArm{
			value: "true",
		},
		FALSE: UpdateImageResponseSupportArm{
			value: "false",
		},
	}
}

func (c UpdateImageResponseSupportArm) Value() string {
	return c.value
}

func (c UpdateImageResponseSupportArm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageResponseSupportArm) UnmarshalJSON(b []byte) error {
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
