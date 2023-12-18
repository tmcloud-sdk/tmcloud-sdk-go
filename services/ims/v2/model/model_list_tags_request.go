package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListTagsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Page *int32 `json:"page,omitempty"`

	Imagetype *ListTagsRequestImagetype `json:"__imagetype,omitempty"`

	Id *string `json:"id,omitempty"`

	Status *ListTagsRequestStatus `json:"status,omitempty"`

	Name *string `json:"name,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	Platform *string `json:"__platform,omitempty"`

	OsType *ListTagsRequestOsType `json:"__os_type,omitempty"`

	MemberStatus *ListTagsRequestMemberStatus `json:"member_status,omitempty"`

	VirtualEnvType *ListTagsRequestVirtualEnvType `json:"virtual_env_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Architecture *ListTagsRequestArchitecture `json:"architecture,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ListTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

type ListTagsRequestImagetype struct {
	value string
}

type ListTagsRequestImagetypeEnum struct {
	GOLD    ListTagsRequestImagetype
	PRIVATE ListTagsRequestImagetype
	SHARED  ListTagsRequestImagetype
	MARKET  ListTagsRequestImagetype
}

func GetListTagsRequestImagetypeEnum() ListTagsRequestImagetypeEnum {
	return ListTagsRequestImagetypeEnum{
		GOLD: ListTagsRequestImagetype{
			value: "gold",
		},
		PRIVATE: ListTagsRequestImagetype{
			value: "private",
		},
		SHARED: ListTagsRequestImagetype{
			value: "shared",
		},
		MARKET: ListTagsRequestImagetype{
			value: "market",
		},
	}
}

func (c ListTagsRequestImagetype) Value() string {
	return c.value
}

func (c ListTagsRequestImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestImagetype) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestStatus struct {
	value string
}

type ListTagsRequestStatusEnum struct {
	QUEUED  ListTagsRequestStatus
	SAVING  ListTagsRequestStatus
	DELETED ListTagsRequestStatus
	KILLED  ListTagsRequestStatus
	ACTIVE  ListTagsRequestStatus
}

func GetListTagsRequestStatusEnum() ListTagsRequestStatusEnum {
	return ListTagsRequestStatusEnum{
		QUEUED: ListTagsRequestStatus{
			value: "queued",
		},
		SAVING: ListTagsRequestStatus{
			value: "saving",
		},
		DELETED: ListTagsRequestStatus{
			value: "deleted",
		},
		KILLED: ListTagsRequestStatus{
			value: "killed",
		},
		ACTIVE: ListTagsRequestStatus{
			value: "active",
		},
	}
}

func (c ListTagsRequestStatus) Value() string {
	return c.value
}

func (c ListTagsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestOsType struct {
	value string
}

type ListTagsRequestOsTypeEnum struct {
	LINUX   ListTagsRequestOsType
	WINDOWS ListTagsRequestOsType
	OTHER   ListTagsRequestOsType
}

func GetListTagsRequestOsTypeEnum() ListTagsRequestOsTypeEnum {
	return ListTagsRequestOsTypeEnum{
		LINUX: ListTagsRequestOsType{
			value: "Linux",
		},
		WINDOWS: ListTagsRequestOsType{
			value: "Windows",
		},
		OTHER: ListTagsRequestOsType{
			value: "Other",
		},
	}
}

func (c ListTagsRequestOsType) Value() string {
	return c.value
}

func (c ListTagsRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestOsType) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestMemberStatus struct {
	value string
}

type ListTagsRequestMemberStatusEnum struct {
	ACCEPTED ListTagsRequestMemberStatus
	REJECTED ListTagsRequestMemberStatus
	PENDING  ListTagsRequestMemberStatus
}

func GetListTagsRequestMemberStatusEnum() ListTagsRequestMemberStatusEnum {
	return ListTagsRequestMemberStatusEnum{
		ACCEPTED: ListTagsRequestMemberStatus{
			value: "accepted",
		},
		REJECTED: ListTagsRequestMemberStatus{
			value: "rejected",
		},
		PENDING: ListTagsRequestMemberStatus{
			value: "pending",
		},
	}
}

func (c ListTagsRequestMemberStatus) Value() string {
	return c.value
}

func (c ListTagsRequestMemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestMemberStatus) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestVirtualEnvType struct {
	value string
}

type ListTagsRequestVirtualEnvTypeEnum struct {
	FUSION_COMPUTE ListTagsRequestVirtualEnvType
	IRONIC         ListTagsRequestVirtualEnvType
	DATA_IMAGE     ListTagsRequestVirtualEnvType
}

func GetListTagsRequestVirtualEnvTypeEnum() ListTagsRequestVirtualEnvTypeEnum {
	return ListTagsRequestVirtualEnvTypeEnum{
		FUSION_COMPUTE: ListTagsRequestVirtualEnvType{
			value: "FusionCompute",
		},
		IRONIC: ListTagsRequestVirtualEnvType{
			value: "Ironic",
		},
		DATA_IMAGE: ListTagsRequestVirtualEnvType{
			value: "DataImage",
		},
	}
}

func (c ListTagsRequestVirtualEnvType) Value() string {
	return c.value
}

func (c ListTagsRequestVirtualEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestVirtualEnvType) UnmarshalJSON(b []byte) error {
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

type ListTagsRequestArchitecture struct {
	value string
}

type ListTagsRequestArchitectureEnum struct {
	X86 ListTagsRequestArchitecture
	ARM ListTagsRequestArchitecture
}

func GetListTagsRequestArchitectureEnum() ListTagsRequestArchitectureEnum {
	return ListTagsRequestArchitectureEnum{
		X86: ListTagsRequestArchitecture{
			value: "x86",
		},
		ARM: ListTagsRequestArchitecture{
			value: "arm",
		},
	}
}

func (c ListTagsRequestArchitecture) Value() string {
	return c.value
}

func (c ListTagsRequestArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagsRequestArchitecture) UnmarshalJSON(b []byte) error {
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
