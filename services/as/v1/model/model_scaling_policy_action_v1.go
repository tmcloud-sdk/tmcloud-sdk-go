package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ScalingPolicyActionV1 struct {
	Operation *ScalingPolicyActionV1Operation `json:"operation,omitempty"`

	InstanceNumber *int32 `json:"instance_number,omitempty"`

	InstancePercentage *int32 `json:"instance_percentage,omitempty"`
}

func (o ScalingPolicyActionV1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyActionV1 struct{}"
	}

	return strings.Join([]string{"ScalingPolicyActionV1", string(data)}, " ")
}

type ScalingPolicyActionV1Operation struct {
	value string
}

type ScalingPolicyActionV1OperationEnum struct {
	ADD    ScalingPolicyActionV1Operation
	REMOVE ScalingPolicyActionV1Operation
	REDUCE ScalingPolicyActionV1Operation
	SET    ScalingPolicyActionV1Operation
}

func GetScalingPolicyActionV1OperationEnum() ScalingPolicyActionV1OperationEnum {
	return ScalingPolicyActionV1OperationEnum{
		ADD: ScalingPolicyActionV1Operation{
			value: "ADD",
		},
		REMOVE: ScalingPolicyActionV1Operation{
			value: "REMOVE",
		},
		REDUCE: ScalingPolicyActionV1Operation{
			value: "REDUCE",
		},
		SET: ScalingPolicyActionV1Operation{
			value: "SET",
		},
	}
}

func (c ScalingPolicyActionV1Operation) Value() string {
	return c.value
}

func (c ScalingPolicyActionV1Operation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyActionV1Operation) UnmarshalJSON(b []byte) error {
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
