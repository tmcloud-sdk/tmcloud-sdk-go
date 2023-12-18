package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateDataImageRequestBody struct {
	CmkId *string `json:"cmk_id,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ImageTags *[]ImageTag `json:"image_tags,omitempty"`

	ImageUrl string `json:"image_url"`

	MinDisk int32 `json:"min_disk"`

	Name string `json:"name"`

	OsType *CreateDataImageRequestBodyOsType `json:"os_type,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateDataImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequestBody", string(data)}, " ")
}

type CreateDataImageRequestBodyOsType struct {
	value string
}

type CreateDataImageRequestBodyOsTypeEnum struct {
	WINDOWS CreateDataImageRequestBodyOsType
	LINUX   CreateDataImageRequestBodyOsType
}

func GetCreateDataImageRequestBodyOsTypeEnum() CreateDataImageRequestBodyOsTypeEnum {
	return CreateDataImageRequestBodyOsTypeEnum{
		WINDOWS: CreateDataImageRequestBodyOsType{
			value: "Windows",
		},
		LINUX: CreateDataImageRequestBodyOsType{
			value: "Linux",
		},
	}
}

func (c CreateDataImageRequestBodyOsType) Value() string {
	return c.value
}

func (c CreateDataImageRequestBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataImageRequestBodyOsType) UnmarshalJSON(b []byte) error {
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
