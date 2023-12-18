package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateLifeCycleHookOption struct {
	LifecycleHookType *UpdateLifeCycleHookOptionLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	DefaultResult *UpdateLifeCycleHookOptionDefaultResult `json:"default_result,omitempty"`

	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	NotificationMetadata *string `json:"notification_metadata,omitempty"`
}

func (o UpdateLifeCycleHookOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLifeCycleHookOption struct{}"
	}

	return strings.Join([]string{"UpdateLifeCycleHookOption", string(data)}, " ")
}

type UpdateLifeCycleHookOptionLifecycleHookType struct {
	value string
}

type UpdateLifeCycleHookOptionLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING UpdateLifeCycleHookOptionLifecycleHookType
	INSTANCE_LAUNCHING   UpdateLifeCycleHookOptionLifecycleHookType
}

func GetUpdateLifeCycleHookOptionLifecycleHookTypeEnum() UpdateLifeCycleHookOptionLifecycleHookTypeEnum {
	return UpdateLifeCycleHookOptionLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: UpdateLifeCycleHookOptionLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: UpdateLifeCycleHookOptionLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c UpdateLifeCycleHookOptionLifecycleHookType) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookOptionLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookOptionLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type UpdateLifeCycleHookOptionDefaultResult struct {
	value string
}

type UpdateLifeCycleHookOptionDefaultResultEnum struct {
	ABANDON  UpdateLifeCycleHookOptionDefaultResult
	CONTINUE UpdateLifeCycleHookOptionDefaultResult
}

func GetUpdateLifeCycleHookOptionDefaultResultEnum() UpdateLifeCycleHookOptionDefaultResultEnum {
	return UpdateLifeCycleHookOptionDefaultResultEnum{
		ABANDON: UpdateLifeCycleHookOptionDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: UpdateLifeCycleHookOptionDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c UpdateLifeCycleHookOptionDefaultResult) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookOptionDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookOptionDefaultResult) UnmarshalJSON(b []byte) error {
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
