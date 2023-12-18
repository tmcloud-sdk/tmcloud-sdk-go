package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ScalingPolicyActionV2 struct {
	Operation *ScalingPolicyActionV2Operation `json:"operation,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Percentage *int32 `json:"percentage,omitempty"`

	Limits *int32 `json:"limits,omitempty"`
}

func (o ScalingPolicyActionV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyActionV2 struct{}"
	}

	return strings.Join([]string{"ScalingPolicyActionV2", string(data)}, " ")
}

type ScalingPolicyActionV2Operation struct {
	value string
}

type ScalingPolicyActionV2OperationEnum struct {
	ADD    ScalingPolicyActionV2Operation
	REMOVE ScalingPolicyActionV2Operation
	REDUCE ScalingPolicyActionV2Operation
	SET    ScalingPolicyActionV2Operation
}

func GetScalingPolicyActionV2OperationEnum() ScalingPolicyActionV2OperationEnum {
	return ScalingPolicyActionV2OperationEnum{
		ADD: ScalingPolicyActionV2Operation{
			value: "ADD",
		},
		REMOVE: ScalingPolicyActionV2Operation{
			value: "REMOVE",
		},
		REDUCE: ScalingPolicyActionV2Operation{
			value: "REDUCE",
		},
		SET: ScalingPolicyActionV2Operation{
			value: "SET",
		},
	}
}

func (c ScalingPolicyActionV2Operation) Value() string {
	return c.value
}

func (c ScalingPolicyActionV2Operation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyActionV2Operation) UnmarshalJSON(b []byte) error {
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
