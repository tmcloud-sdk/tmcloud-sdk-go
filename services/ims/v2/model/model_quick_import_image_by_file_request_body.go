package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QuickImportImageByFileRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	OsVersion string `json:"os_version"`

	ImageUrl string `json:"image_url"`

	MinDisk int32 `json:"min_disk"`

	Tags *[]string `json:"tags,omitempty"`

	Type *QuickImportImageByFileRequestBodyType `json:"type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Architecture *QuickImportImageByFileRequestBodyArchitecture `json:"architecture,omitempty"`

	OsType *QuickImportImageByFileRequestBodyOsType `json:"os_type,omitempty"`

	ImageTags *[]ResourceTag `json:"image_tags,omitempty"`
}

func (o QuickImportImageByFileRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuickImportImageByFileRequestBody struct{}"
	}

	return strings.Join([]string{"QuickImportImageByFileRequestBody", string(data)}, " ")
}

type QuickImportImageByFileRequestBodyType struct {
	value string
}

type QuickImportImageByFileRequestBodyTypeEnum struct {
	ECS        QuickImportImageByFileRequestBodyType
	BMS        QuickImportImageByFileRequestBodyType
	DATA_IMAGE QuickImportImageByFileRequestBodyType
}

func GetQuickImportImageByFileRequestBodyTypeEnum() QuickImportImageByFileRequestBodyTypeEnum {
	return QuickImportImageByFileRequestBodyTypeEnum{
		ECS: QuickImportImageByFileRequestBodyType{
			value: "ECS",
		},
		BMS: QuickImportImageByFileRequestBodyType{
			value: "BMS",
		},
		DATA_IMAGE: QuickImportImageByFileRequestBodyType{
			value: "DataImage",
		},
	}
}

func (c QuickImportImageByFileRequestBodyType) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyType) UnmarshalJSON(b []byte) error {
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

type QuickImportImageByFileRequestBodyArchitecture struct {
	value string
}

type QuickImportImageByFileRequestBodyArchitectureEnum struct {
	X86 QuickImportImageByFileRequestBodyArchitecture
	ARM QuickImportImageByFileRequestBodyArchitecture
}

func GetQuickImportImageByFileRequestBodyArchitectureEnum() QuickImportImageByFileRequestBodyArchitectureEnum {
	return QuickImportImageByFileRequestBodyArchitectureEnum{
		X86: QuickImportImageByFileRequestBodyArchitecture{
			value: "x86",
		},
		ARM: QuickImportImageByFileRequestBodyArchitecture{
			value: "arm",
		},
	}
}

func (c QuickImportImageByFileRequestBodyArchitecture) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyArchitecture) UnmarshalJSON(b []byte) error {
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

type QuickImportImageByFileRequestBodyOsType struct {
	value string
}

type QuickImportImageByFileRequestBodyOsTypeEnum struct {
	LINUX   QuickImportImageByFileRequestBodyOsType
	WINDOWS QuickImportImageByFileRequestBodyOsType
}

func GetQuickImportImageByFileRequestBodyOsTypeEnum() QuickImportImageByFileRequestBodyOsTypeEnum {
	return QuickImportImageByFileRequestBodyOsTypeEnum{
		LINUX: QuickImportImageByFileRequestBodyOsType{
			value: "Linux",
		},
		WINDOWS: QuickImportImageByFileRequestBodyOsType{
			value: "Windows",
		},
	}
}

func (c QuickImportImageByFileRequestBodyOsType) Value() string {
	return c.value
}

func (c QuickImportImageByFileRequestBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuickImportImageByFileRequestBodyOsType) UnmarshalJSON(b []byte) error {
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
