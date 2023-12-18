package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingActivityLogV2 struct {
	Status *ScalingActivityLogV2Status `json:"status,omitempty"`

	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	Id *string `json:"id,omitempty"`

	InstanceRemovedList *[]ScalingInstance `json:"instance_removed_list,omitempty"`

	InstanceDeletedList *[]ScalingInstance `json:"instance_deleted_list,omitempty"`

	InstanceAddedList *[]ScalingInstance `json:"instance_added_list,omitempty"`

	InstanceFailedList *[]ScalingInstance `json:"instance_failed_list,omitempty"`

	InstanceStandbyList *[]ScalingInstance `json:"instance_standby_list,omitempty"`

	ScalingValue *string `json:"scaling_value,omitempty"`

	Description *string `json:"description,omitempty"`

	InstanceValue *int32 `json:"instance_value,omitempty"`

	DesireValue *int32 `json:"desire_value,omitempty"`

	LbBindSuccessList *[]ModifyLb `json:"lb_bind_success_list,omitempty"`

	LbBindFailedList *[]ModifyLb `json:"lb_bind_failed_list,omitempty"`

	LbUnbindSuccessList *[]ModifyLb `json:"lb_unbind_success_list,omitempty"`

	LbUnbindFailedList *[]ModifyLb `json:"lb_unbind_failed_list,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ScalingActivityLogV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingActivityLogV2 struct{}"
	}

	return strings.Join([]string{"ScalingActivityLogV2", string(data)}, " ")
}

type ScalingActivityLogV2Status struct {
	value string
}

type ScalingActivityLogV2StatusEnum struct {
	SUCCESS ScalingActivityLogV2Status
	FAIL    ScalingActivityLogV2Status
	DING    ScalingActivityLogV2Status
}

func GetScalingActivityLogV2StatusEnum() ScalingActivityLogV2StatusEnum {
	return ScalingActivityLogV2StatusEnum{
		SUCCESS: ScalingActivityLogV2Status{
			value: "SUCCESS",
		},
		FAIL: ScalingActivityLogV2Status{
			value: "FAIL",
		},
		DING: ScalingActivityLogV2Status{
			value: "DING",
		},
	}
}

func (c ScalingActivityLogV2Status) Value() string {
	return c.value
}

func (c ScalingActivityLogV2Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingActivityLogV2Status) UnmarshalJSON(b []byte) error {
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
