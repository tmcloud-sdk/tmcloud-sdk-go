package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteScalingInstanceRequest struct {
	InstanceId string `json:"instance_id"`

	InstanceDelete *DeleteScalingInstanceRequestInstanceDelete `json:"instance_delete,omitempty"`
}

func (o DeleteScalingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingInstanceRequest", string(data)}, " ")
}

type DeleteScalingInstanceRequestInstanceDelete struct {
	value string
}

type DeleteScalingInstanceRequestInstanceDeleteEnum struct {
	YES DeleteScalingInstanceRequestInstanceDelete
	NO  DeleteScalingInstanceRequestInstanceDelete
}

func GetDeleteScalingInstanceRequestInstanceDeleteEnum() DeleteScalingInstanceRequestInstanceDeleteEnum {
	return DeleteScalingInstanceRequestInstanceDeleteEnum{
		YES: DeleteScalingInstanceRequestInstanceDelete{
			value: "yes",
		},
		NO: DeleteScalingInstanceRequestInstanceDelete{
			value: "no",
		},
	}
}

func (c DeleteScalingInstanceRequestInstanceDelete) Value() string {
	return c.value
}

func (c DeleteScalingInstanceRequestInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScalingInstanceRequestInstanceDelete) UnmarshalJSON(b []byte) error {
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
