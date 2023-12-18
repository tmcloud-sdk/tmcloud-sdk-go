package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListImagesRequest struct {
	Imagetype *ListImagesRequestImagetype `json:"__imagetype,omitempty"`

	Isregistered *ListImagesRequestIsregistered `json:"__isregistered,omitempty"`

	OsBit *ListImagesRequestOsBit `json:"__os_bit,omitempty"`

	OsType *ListImagesRequestOsType `json:"__os_type,omitempty"`

	Platform *ListImagesRequestPlatform `json:"__platform,omitempty"`

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

	DiskFormat *ListImagesRequestDiskFormat `json:"disk_format,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Id *string `json:"id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	MemberStatus *ListImagesRequestMemberStatus `json:"member_status,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	SortDir *ListImagesRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *ListImagesRequestSortKey `json:"sort_key,omitempty"`

	Status *ListImagesRequestStatus `json:"status,omitempty"`

	Tag *string `json:"tag,omitempty"`

	VirtualEnvType *ListImagesRequestVirtualEnvType `json:"virtual_env_type,omitempty"`

	Visibility *ListImagesRequestVisibility `json:"visibility,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	FlavorId *string `json:"flavor_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Architecture *ListImagesRequestArchitecture `json:"architecture,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}

type ListImagesRequestImagetype struct {
	value string
}

type ListImagesRequestImagetypeEnum struct {
	GOLD    ListImagesRequestImagetype
	PRIVATE ListImagesRequestImagetype
	SHARED  ListImagesRequestImagetype
	MARKET  ListImagesRequestImagetype
}

func GetListImagesRequestImagetypeEnum() ListImagesRequestImagetypeEnum {
	return ListImagesRequestImagetypeEnum{
		GOLD: ListImagesRequestImagetype{
			value: "gold",
		},
		PRIVATE: ListImagesRequestImagetype{
			value: "private",
		},
		SHARED: ListImagesRequestImagetype{
			value: "shared",
		},
		MARKET: ListImagesRequestImagetype{
			value: "market",
		},
	}
}

func (c ListImagesRequestImagetype) Value() string {
	return c.value
}

func (c ListImagesRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestImagetype) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestIsregistered struct {
	value string
}

type ListImagesRequestIsregisteredEnum struct {
	TRUE  ListImagesRequestIsregistered
	FALSE ListImagesRequestIsregistered
}

func GetListImagesRequestIsregisteredEnum() ListImagesRequestIsregisteredEnum {
	return ListImagesRequestIsregisteredEnum{
		TRUE: ListImagesRequestIsregistered{
			value: "true",
		},
		FALSE: ListImagesRequestIsregistered{
			value: "false",
		},
	}
}

func (c ListImagesRequestIsregistered) Value() string {
	return c.value
}

func (c ListImagesRequestIsregistered) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestIsregistered) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestOsBit struct {
	value string
}

type ListImagesRequestOsBitEnum struct {
	E_32 ListImagesRequestOsBit
	E_64 ListImagesRequestOsBit
}

func GetListImagesRequestOsBitEnum() ListImagesRequestOsBitEnum {
	return ListImagesRequestOsBitEnum{
		E_32: ListImagesRequestOsBit{
			value: "32",
		},
		E_64: ListImagesRequestOsBit{
			value: "64",
		},
	}
}

func (c ListImagesRequestOsBit) Value() string {
	return c.value
}

func (c ListImagesRequestOsBit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestOsBit) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestOsType struct {
	value string
}

type ListImagesRequestOsTypeEnum struct {
	LINUX   ListImagesRequestOsType
	WINDOWS ListImagesRequestOsType
	OTHER   ListImagesRequestOsType
}

func GetListImagesRequestOsTypeEnum() ListImagesRequestOsTypeEnum {
	return ListImagesRequestOsTypeEnum{
		LINUX: ListImagesRequestOsType{
			value: "Linux",
		},
		WINDOWS: ListImagesRequestOsType{
			value: "Windows",
		},
		OTHER: ListImagesRequestOsType{
			value: "Other",
		},
	}
}

func (c ListImagesRequestOsType) Value() string {
	return c.value
}

func (c ListImagesRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestOsType) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestPlatform struct {
	value string
}

type ListImagesRequestPlatformEnum struct {
	WINDOWS      ListImagesRequestPlatform
	UBUNTU       ListImagesRequestPlatform
	RED_HAT      ListImagesRequestPlatform
	SUSE         ListImagesRequestPlatform
	CENT_OS      ListImagesRequestPlatform
	DEBIAN       ListImagesRequestPlatform
	OPEN_SUSE    ListImagesRequestPlatform
	ORACLE_LINUX ListImagesRequestPlatform
	FEDORA       ListImagesRequestPlatform
	OTHER        ListImagesRequestPlatform
	CORE_OS      ListImagesRequestPlatform
	EULER_OS     ListImagesRequestPlatform
}

func GetListImagesRequestPlatformEnum() ListImagesRequestPlatformEnum {
	return ListImagesRequestPlatformEnum{
		WINDOWS: ListImagesRequestPlatform{
			value: "Windows",
		},
		UBUNTU: ListImagesRequestPlatform{
			value: "Ubuntu",
		},
		RED_HAT: ListImagesRequestPlatform{
			value: "RedHat",
		},
		SUSE: ListImagesRequestPlatform{
			value: "SUSE",
		},
		CENT_OS: ListImagesRequestPlatform{
			value: "CentOS",
		},
		DEBIAN: ListImagesRequestPlatform{
			value: "Debian",
		},
		OPEN_SUSE: ListImagesRequestPlatform{
			value: "OpenSUSE",
		},
		ORACLE_LINUX: ListImagesRequestPlatform{
			value: "Oracle Linux",
		},
		FEDORA: ListImagesRequestPlatform{
			value: "Fedora",
		},
		OTHER: ListImagesRequestPlatform{
			value: "Other",
		},
		CORE_OS: ListImagesRequestPlatform{
			value: "CoreOS",
		},
		EULER_OS: ListImagesRequestPlatform{
			value: "EulerOS",
		},
	}
}

func (c ListImagesRequestPlatform) Value() string {
	return c.value
}

func (c ListImagesRequestPlatform) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestPlatform) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestDiskFormat struct {
	value string
}

type ListImagesRequestDiskFormatEnum struct {
	VHD   ListImagesRequestDiskFormat
	ZVHD  ListImagesRequestDiskFormat
	RAW   ListImagesRequestDiskFormat
	QCOW2 ListImagesRequestDiskFormat
	ZVHD2 ListImagesRequestDiskFormat
}

func GetListImagesRequestDiskFormatEnum() ListImagesRequestDiskFormatEnum {
	return ListImagesRequestDiskFormatEnum{
		VHD: ListImagesRequestDiskFormat{
			value: "vhd",
		},
		ZVHD: ListImagesRequestDiskFormat{
			value: "zvhd",
		},
		RAW: ListImagesRequestDiskFormat{
			value: "raw",
		},
		QCOW2: ListImagesRequestDiskFormat{
			value: "qcow2",
		},
		ZVHD2: ListImagesRequestDiskFormat{
			value: "zvhd2",
		},
	}
}

func (c ListImagesRequestDiskFormat) Value() string {
	return c.value
}

func (c ListImagesRequestDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestDiskFormat) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestMemberStatus struct {
	value string
}

type ListImagesRequestMemberStatusEnum struct {
	ACCEPTED ListImagesRequestMemberStatus
	REJECTED ListImagesRequestMemberStatus
	PENDING  ListImagesRequestMemberStatus
}

func GetListImagesRequestMemberStatusEnum() ListImagesRequestMemberStatusEnum {
	return ListImagesRequestMemberStatusEnum{
		ACCEPTED: ListImagesRequestMemberStatus{
			value: "accepted",
		},
		REJECTED: ListImagesRequestMemberStatus{
			value: "rejected",
		},
		PENDING: ListImagesRequestMemberStatus{
			value: "pending",
		},
	}
}

func (c ListImagesRequestMemberStatus) Value() string {
	return c.value
}

func (c ListImagesRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestMemberStatus) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestSortDir struct {
	value string
}

type ListImagesRequestSortDirEnum struct {
	ASC  ListImagesRequestSortDir
	DESC ListImagesRequestSortDir
}

func GetListImagesRequestSortDirEnum() ListImagesRequestSortDirEnum {
	return ListImagesRequestSortDirEnum{
		ASC: ListImagesRequestSortDir{
			value: "asc",
		},
		DESC: ListImagesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListImagesRequestSortDir) Value() string {
	return c.value
}

func (c ListImagesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestSortKey struct {
	value string
}

type ListImagesRequestSortKeyEnum struct {
	CREATED_AT       ListImagesRequestSortKey
	NAME             ListImagesRequestSortKey
	CONTAINER_FORMAT ListImagesRequestSortKey
	DISK_FORMAT      ListImagesRequestSortKey
	STATUS           ListImagesRequestSortKey
	ID               ListImagesRequestSortKey
	SIZE             ListImagesRequestSortKey
}

func GetListImagesRequestSortKeyEnum() ListImagesRequestSortKeyEnum {
	return ListImagesRequestSortKeyEnum{
		CREATED_AT: ListImagesRequestSortKey{
			value: "created_at",
		},
		NAME: ListImagesRequestSortKey{
			value: "name",
		},
		CONTAINER_FORMAT: ListImagesRequestSortKey{
			value: "container_format",
		},
		DISK_FORMAT: ListImagesRequestSortKey{
			value: "disk_format",
		},
		STATUS: ListImagesRequestSortKey{
			value: "status ",
		},
		ID: ListImagesRequestSortKey{
			value: "id",
		},
		SIZE: ListImagesRequestSortKey{
			value: "size",
		},
	}
}

func (c ListImagesRequestSortKey) Value() string {
	return c.value
}

func (c ListImagesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestStatus struct {
	value string
}

type ListImagesRequestStatusEnum struct {
	QUEUED  ListImagesRequestStatus
	SAVING  ListImagesRequestStatus
	DELETED ListImagesRequestStatus
	KILLED  ListImagesRequestStatus
	ACTIVE  ListImagesRequestStatus
}

func GetListImagesRequestStatusEnum() ListImagesRequestStatusEnum {
	return ListImagesRequestStatusEnum{
		QUEUED: ListImagesRequestStatus{
			value: "queued",
		},
		SAVING: ListImagesRequestStatus{
			value: "saving",
		},
		DELETED: ListImagesRequestStatus{
			value: "deleted",
		},
		KILLED: ListImagesRequestStatus{
			value: "killed",
		},
		ACTIVE: ListImagesRequestStatus{
			value: "active",
		},
	}
}

func (c ListImagesRequestStatus) Value() string {
	return c.value
}

func (c ListImagesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestVirtualEnvType struct {
	value string
}

type ListImagesRequestVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ListImagesRequestVirtualEnvType
	IRONIC         ListImagesRequestVirtualEnvType
	DATA_IMAGE     ListImagesRequestVirtualEnvType
}

func GetListImagesRequestVirtualEnvTypeEnum() ListImagesRequestVirtualEnvTypeEnum {
	return ListImagesRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListImagesRequestVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ListImagesRequestVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ListImagesRequestVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ListImagesRequestVirtualEnvType) Value() string {
	return c.value
}

func (c ListImagesRequestVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestVisibility struct {
	value string
}

type ListImagesRequestVisibilityEnum struct {
	PUBLIC  ListImagesRequestVisibility
	PRIVATE ListImagesRequestVisibility
}

func GetListImagesRequestVisibilityEnum() ListImagesRequestVisibilityEnum {
	return ListImagesRequestVisibilityEnum{
		PUBLIC: ListImagesRequestVisibility{
			value: "public",
		},
		PRIVATE: ListImagesRequestVisibility{
			value: "private",
		},
	}
}

func (c ListImagesRequestVisibility) Value() string {
	return c.value
}

func (c ListImagesRequestVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestVisibility) UnmarshalJSON(b []byte) error {
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

type ListImagesRequestArchitecture struct {
	value string
}

type ListImagesRequestArchitectureEnum struct {
	X86 ListImagesRequestArchitecture
	ARM ListImagesRequestArchitecture
}

func GetListImagesRequestArchitectureEnum() ListImagesRequestArchitectureEnum {
	return ListImagesRequestArchitectureEnum{
		X86: ListImagesRequestArchitecture{
			value: "x86",
		},
		ARM: ListImagesRequestArchitecture{
			value: "arm",
		},
	}
}

func (c ListImagesRequestArchitecture) Value() string {
	return c.value
}

func (c ListImagesRequestArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImagesRequestArchitecture) UnmarshalJSON(b []byte) error {
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
