package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	ForceDelete *DeleteScalingGroupRequestForceDelete `json:"force_delete,omitempty"`
}

func (o DeleteScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingGroupRequest", string(data)}, " ")
}

type DeleteScalingGroupRequestForceDelete struct {
	value string
}

type DeleteScalingGroupRequestForceDeleteEnum struct {
	YES DeleteScalingGroupRequestForceDelete
	NO  DeleteScalingGroupRequestForceDelete
}

func GetDeleteScalingGroupRequestForceDeleteEnum() DeleteScalingGroupRequestForceDeleteEnum {
	return DeleteScalingGroupRequestForceDeleteEnum{
		YES: DeleteScalingGroupRequestForceDelete{
			value: "yes",
		},
		NO: DeleteScalingGroupRequestForceDelete{
			value: "no",
		},
	}
}

func (c DeleteScalingGroupRequestForceDelete) Value() string {
	return c.value
}

func (c DeleteScalingGroupRequestForceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScalingGroupRequestForceDelete) UnmarshalJSON(b []byte) error {
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
