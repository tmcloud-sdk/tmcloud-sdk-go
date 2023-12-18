package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type TransformationInfo struct {
	TransformationType TransformationInfoTransformationType `json:"transformation_type"`

	Value string `json:"value"`
}

func (o TransformationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformationInfo struct{}"
	}

	return strings.Join([]string{"TransformationInfo", string(data)}, " ")
}

type TransformationInfoTransformationType struct {
	value string
}

type TransformationInfoTransformationTypeEnum struct {
	CONTENT_CONDITIONAL_FILTER TransformationInfoTransformationType
	CONFIG_CONDITIONAL_FILTER  TransformationInfoTransformationType
}

func GetTransformationInfoTransformationTypeEnum() TransformationInfoTransformationTypeEnum {
	return TransformationInfoTransformationTypeEnum{
		CONTENT_CONDITIONAL_FILTER: TransformationInfoTransformationType{
			value: "contentConditionalFilter",
		},
		CONFIG_CONDITIONAL_FILTER: TransformationInfoTransformationType{
			value: "configConditionalFilter",
		},
	}
}

func (c TransformationInfoTransformationType) Value() string {
	return c.value
}

func (c TransformationInfoTransformationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransformationInfoTransformationType) UnmarshalJSON(b []byte) error {
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
