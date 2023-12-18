package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateLifeCycleHookOption struct {
	LifecycleHookName string `json:"lifecycle_hook_name"`

	LifecycleHookType CreateLifeCycleHookOptionLifecycleHookType `json:"lifecycle_hook_type"`

	DefaultResult *CreateLifeCycleHookOptionDefaultResult `json:"default_result,omitempty"`

	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	NotificationTopicUrn string `json:"notification_topic_urn"`

	NotificationMetadata *string `json:"notification_metadata,omitempty"`
}

func (o CreateLifeCycleHookOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLifeCycleHookOption struct{}"
	}

	return strings.Join([]string{"CreateLifeCycleHookOption", string(data)}, " ")
}

type CreateLifeCycleHookOptionLifecycleHookType struct {
	value string
}

type CreateLifeCycleHookOptionLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING CreateLifeCycleHookOptionLifecycleHookType
	INSTANCE_LAUNCHING   CreateLifeCycleHookOptionLifecycleHookType
}

func GetCreateLifeCycleHookOptionLifecycleHookTypeEnum() CreateLifeCycleHookOptionLifecycleHookTypeEnum {
	return CreateLifeCycleHookOptionLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: CreateLifeCycleHookOptionLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: CreateLifeCycleHookOptionLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c CreateLifeCycleHookOptionLifecycleHookType) Value() string {
	return c.value
}

func (c CreateLifeCycleHookOptionLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifeCycleHookOptionLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type CreateLifeCycleHookOptionDefaultResult struct {
	value string
}

type CreateLifeCycleHookOptionDefaultResultEnum struct {
	ABANDON  CreateLifeCycleHookOptionDefaultResult
	CONTINUE CreateLifeCycleHookOptionDefaultResult
}

func GetCreateLifeCycleHookOptionDefaultResultEnum() CreateLifeCycleHookOptionDefaultResultEnum {
	return CreateLifeCycleHookOptionDefaultResultEnum{
		ABANDON: CreateLifeCycleHookOptionDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: CreateLifeCycleHookOptionDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c CreateLifeCycleHookOptionDefaultResult) Value() string {
	return c.value
}

func (c CreateLifeCycleHookOptionDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifeCycleHookOptionDefaultResult) UnmarshalJSON(b []byte) error {
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
