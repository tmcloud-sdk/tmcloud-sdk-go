package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Resources struct {
	Type *ResourcesType `json:"type,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Quota *int32 `json:"quota,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}

type ResourcesType struct {
	value string
}

type ResourcesTypeEnum struct {
	CMK           ResourcesType
	GRANT_PER_CMK ResourcesType
}

func GetResourcesTypeEnum() ResourcesTypeEnum {
	return ResourcesTypeEnum{
		CMK: ResourcesType{
			value: "CMK",
		},
		GRANT_PER_CMK: ResourcesType{
			value: "grant_per_CMK",
		},
	}
}

func (c ResourcesType) Value() string {
	return c.value
}

func (c ResourcesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcesType) UnmarshalJSON(b []byte) error {
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
