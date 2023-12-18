package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateImageRequestBody struct {
	DataImages *[]CreateDataImage `json:"data_images,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ImageTags *[]TagKeyValue `json:"image_tags,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Name string `json:"name"`

	Tags *[]string `json:"tags,omitempty"`

	MaxRam *int32 `json:"max_ram,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	OsVersion *string `json:"os_version,omitempty"`

	ImageUrl *string `json:"image_url,omitempty"`

	MinDisk *int32 `json:"min_disk,omitempty"`

	IsConfig *bool `json:"is_config,omitempty"`

	CmkId *string `json:"cmk_id,omitempty"`

	Type *CreateImageRequestBodyType `json:"type,omitempty"`

	IsQuickImport *bool `json:"is_quick_import,omitempty"`

	Architecture *CreateImageRequestBodyArchitecture `json:"architecture,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`
}

func (o CreateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageRequestBody", string(data)}, " ")
}

type CreateImageRequestBodyType struct {
	value string
}

type CreateImageRequestBodyTypeEnum struct {
	ECS            CreateImageRequestBodyType
	BMS            CreateImageRequestBodyType
	FUSION_COMPUTE CreateImageRequestBodyType
	IRONIC         CreateImageRequestBodyType
	ISO_IMAGE      CreateImageRequestBodyType
}

func GetCreateImageRequestBodyTypeEnum() CreateImageRequestBodyTypeEnum {
	return CreateImageRequestBodyTypeEnum{
		ECS: CreateImageRequestBodyType{
			value: "ECS",
		},
		BMS: CreateImageRequestBodyType{
			value: "BMS",
		},
		FUSION_COMPUTE: CreateImageRequestBodyType{
			value: "FusionCompute",
		},
		IRONIC: CreateImageRequestBodyType{
			value: "Ironic",
		},
		ISO_IMAGE: CreateImageRequestBodyType{
			value: "IsoImage",
		},
	}
}

func (c CreateImageRequestBodyType) Value() string {
	return c.value
}

func (c CreateImageRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageRequestBodyType) UnmarshalJSON(b []byte) error {
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

type CreateImageRequestBodyArchitecture struct {
	value string
}

type CreateImageRequestBodyArchitectureEnum struct {
	X86 CreateImageRequestBodyArchitecture
	ARM CreateImageRequestBodyArchitecture
}

func GetCreateImageRequestBodyArchitectureEnum() CreateImageRequestBodyArchitectureEnum {
	return CreateImageRequestBodyArchitectureEnum{
		X86: CreateImageRequestBodyArchitecture{
			value: "x86",
		},
		ARM: CreateImageRequestBodyArchitecture{
			value: "arm",
		},
	}
}

func (c CreateImageRequestBodyArchitecture) Value() string {
	return c.value
}

func (c CreateImageRequestBodyArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateImageRequestBodyArchitecture) UnmarshalJSON(b []byte) error {
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
