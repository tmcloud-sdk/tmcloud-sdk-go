package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CallbackLifeCycleHookOption struct {
	LifecycleActionKey *string `json:"lifecycle_action_key,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	LifecycleActionResult CallbackLifeCycleHookOptionLifecycleActionResult `json:"lifecycle_action_result"`
}

func (o CallbackLifeCycleHookOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CallbackLifeCycleHookOption struct{}"
	}

	return strings.Join([]string{"CallbackLifeCycleHookOption", string(data)}, " ")
}

type CallbackLifeCycleHookOptionLifecycleActionResult struct {
	value string
}

type CallbackLifeCycleHookOptionLifecycleActionResultEnum struct {
	ABANDON  CallbackLifeCycleHookOptionLifecycleActionResult
	CONTINUE CallbackLifeCycleHookOptionLifecycleActionResult
	EXTEND   CallbackLifeCycleHookOptionLifecycleActionResult
}

func GetCallbackLifeCycleHookOptionLifecycleActionResultEnum() CallbackLifeCycleHookOptionLifecycleActionResultEnum {
	return CallbackLifeCycleHookOptionLifecycleActionResultEnum{
		ABANDON: CallbackLifeCycleHookOptionLifecycleActionResult{
			value: "ABANDON",
		},
		CONTINUE: CallbackLifeCycleHookOptionLifecycleActionResult{
			value: "CONTINUE",
		},
		EXTEND: CallbackLifeCycleHookOptionLifecycleActionResult{
			value: "EXTEND",
		},
	}
}

func (c CallbackLifeCycleHookOptionLifecycleActionResult) Value() string {
	return c.value
}

func (c CallbackLifeCycleHookOptionLifecycleActionResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CallbackLifeCycleHookOptionLifecycleActionResult) UnmarshalJSON(b []byte) error {
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
