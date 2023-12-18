package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type AllResources struct {
	Type *AllResourcesType `json:"type,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Quota *int32 `json:"quota,omitempty"`

	Max *int32 `json:"max,omitempty"`

	Min *int32 `json:"min,omitempty"`
}

func (o AllResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllResources struct{}"
	}

	return strings.Join([]string{"AllResources", string(data)}, " ")
}

type AllResourcesType struct {
	value string
}

type AllResourcesTypeEnum struct {
	SCALING_GROUP            AllResourcesType
	SCALING_CONFIG           AllResourcesType
	SCALING_POLICY           AllResourcesType
	SCALING_INSTANCE         AllResourcesType
	BANDWIDTH_SCALING_POLICY AllResourcesType
}

func GetAllResourcesTypeEnum() AllResourcesTypeEnum {
	return AllResourcesTypeEnum{
		SCALING_GROUP: AllResourcesType{
			value: "scaling_group",
		},
		SCALING_CONFIG: AllResourcesType{
			value: "scaling_config",
		},
		SCALING_POLICY: AllResourcesType{
			value: "scaling_Policy",
		},
		SCALING_INSTANCE: AllResourcesType{
			value: "scaling_Instance",
		},
		BANDWIDTH_SCALING_POLICY: AllResourcesType{
			value: "bandwidth_scaling_policy",
		},
	}
}

func (c AllResourcesType) Value() string {
	return c.value
}

func (c AllResourcesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AllResourcesType) UnmarshalJSON(b []byte) error {
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
