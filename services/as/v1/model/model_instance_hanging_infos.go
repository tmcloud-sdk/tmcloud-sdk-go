package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type InstanceHangingInfos struct {
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	LifecycleActionKey *string `json:"lifecycle_action_key,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	LifecycleHookStatus *InstanceHangingInfosLifecycleHookStatus `json:"lifecycle_hook_status,omitempty"`

	Timeout *string `json:"timeout,omitempty"`

	DefaultResult *string `json:"default_result,omitempty"`
}

func (o InstanceHangingInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceHangingInfos struct{}"
	}

	return strings.Join([]string{"InstanceHangingInfos", string(data)}, " ")
}

type InstanceHangingInfosLifecycleHookStatus struct {
	value string
}

type InstanceHangingInfosLifecycleHookStatusEnum struct {
	HANGING  InstanceHangingInfosLifecycleHookStatus
	CONTINUE InstanceHangingInfosLifecycleHookStatus
	ABANDON  InstanceHangingInfosLifecycleHookStatus
}

func GetInstanceHangingInfosLifecycleHookStatusEnum() InstanceHangingInfosLifecycleHookStatusEnum {
	return InstanceHangingInfosLifecycleHookStatusEnum{
		HANGING: InstanceHangingInfosLifecycleHookStatus{
			value: "HANGING",
		},
		CONTINUE: InstanceHangingInfosLifecycleHookStatus{
			value: "CONTINUE",
		},
		ABANDON: InstanceHangingInfosLifecycleHookStatus{
			value: "ABANDON",
		},
	}
}

func (c InstanceHangingInfosLifecycleHookStatus) Value() string {
	return c.value
}

func (c InstanceHangingInfosLifecycleHookStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceHangingInfosLifecycleHookStatus) UnmarshalJSON(b []byte) error {
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
