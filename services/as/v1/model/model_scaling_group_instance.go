package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingGroupInstance struct {
	InstanceId *string `json:"instance_id,omitempty"`

	InstanceName *string `json:"instance_name,omitempty"`

	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	ScalingGroupName *string `json:"scaling_group_name,omitempty"`

	LifeCycleState *ScalingGroupInstanceLifeCycleState `json:"life_cycle_state,omitempty"`

	HealthStatus *ScalingGroupInstanceHealthStatus `json:"health_status,omitempty"`

	ScalingConfigurationName *string `json:"scaling_configuration_name,omitempty"`

	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	ProtectFromScalingDown *bool `json:"protect_from_scaling_down,omitempty"`
}

func (o ScalingGroupInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingGroupInstance struct{}"
	}

	return strings.Join([]string{"ScalingGroupInstance", string(data)}, " ")
}

type ScalingGroupInstanceLifeCycleState struct {
	value string
}

type ScalingGroupInstanceLifeCycleStateEnum struct {
	INSERVICE     ScalingGroupInstanceLifeCycleState
	PENDING       ScalingGroupInstanceLifeCycleState
	REMOVING      ScalingGroupInstanceLifeCycleState
	PENDING_WAIT  ScalingGroupInstanceLifeCycleState
	REMOVING_WAIT ScalingGroupInstanceLifeCycleState
}

func GetScalingGroupInstanceLifeCycleStateEnum() ScalingGroupInstanceLifeCycleStateEnum {
	return ScalingGroupInstanceLifeCycleStateEnum{
		INSERVICE: ScalingGroupInstanceLifeCycleState{
			value: "INSERVICE",
		},
		PENDING: ScalingGroupInstanceLifeCycleState{
			value: "PENDING",
		},
		REMOVING: ScalingGroupInstanceLifeCycleState{
			value: "REMOVING",
		},
		PENDING_WAIT: ScalingGroupInstanceLifeCycleState{
			value: "PENDING_WAIT",
		},
		REMOVING_WAIT: ScalingGroupInstanceLifeCycleState{
			value: "REMOVING_WAIT",
		},
	}
}

func (c ScalingGroupInstanceLifeCycleState) Value() string {
	return c.value
}

func (c ScalingGroupInstanceLifeCycleState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingGroupInstanceLifeCycleState) UnmarshalJSON(b []byte) error {
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

type ScalingGroupInstanceHealthStatus struct {
	value string
}

type ScalingGroupInstanceHealthStatusEnum struct {
	NORMAL       ScalingGroupInstanceHealthStatus
	ERROR        ScalingGroupInstanceHealthStatus
	INITAILIZING ScalingGroupInstanceHealthStatus
}

func GetScalingGroupInstanceHealthStatusEnum() ScalingGroupInstanceHealthStatusEnum {
	return ScalingGroupInstanceHealthStatusEnum{
		NORMAL: ScalingGroupInstanceHealthStatus{
			value: "NORMAL",
		},
		ERROR: ScalingGroupInstanceHealthStatus{
			value: "ERROR",
		},
		INITAILIZING: ScalingGroupInstanceHealthStatus{
			value: "INITAILIZING",
		},
	}
}

func (c ScalingGroupInstanceHealthStatus) Value() string {
	return c.value
}

func (c ScalingGroupInstanceHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingGroupInstanceHealthStatus) UnmarshalJSON(b []byte) error {
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
