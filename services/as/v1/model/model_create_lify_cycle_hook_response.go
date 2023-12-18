package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateLifyCycleHookResponse struct {
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	LifecycleHookType *CreateLifyCycleHookResponseLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	DefaultResult *CreateLifyCycleHookResponseDefaultResult `json:"default_result,omitempty"`

	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	NotificationTopicName *string `json:"notification_topic_name,omitempty"`

	NotificationMetadata *string `json:"notification_metadata,omitempty"`

	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLifyCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLifyCycleHookResponse struct{}"
	}

	return strings.Join([]string{"CreateLifyCycleHookResponse", string(data)}, " ")
}

type CreateLifyCycleHookResponseLifecycleHookType struct {
	value string
}

type CreateLifyCycleHookResponseLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING CreateLifyCycleHookResponseLifecycleHookType
	INSTANCE_LAUNCHING   CreateLifyCycleHookResponseLifecycleHookType
}

func GetCreateLifyCycleHookResponseLifecycleHookTypeEnum() CreateLifyCycleHookResponseLifecycleHookTypeEnum {
	return CreateLifyCycleHookResponseLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: CreateLifyCycleHookResponseLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: CreateLifyCycleHookResponseLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c CreateLifyCycleHookResponseLifecycleHookType) Value() string {
	return c.value
}

func (c CreateLifyCycleHookResponseLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifyCycleHookResponseLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type CreateLifyCycleHookResponseDefaultResult struct {
	value string
}

type CreateLifyCycleHookResponseDefaultResultEnum struct {
	ABANDON  CreateLifyCycleHookResponseDefaultResult
	CONTINUE CreateLifyCycleHookResponseDefaultResult
}

func GetCreateLifyCycleHookResponseDefaultResultEnum() CreateLifyCycleHookResponseDefaultResultEnum {
	return CreateLifyCycleHookResponseDefaultResultEnum{
		ABANDON: CreateLifyCycleHookResponseDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: CreateLifyCycleHookResponseDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c CreateLifyCycleHookResponseDefaultResult) Value() string {
	return c.value
}

func (c CreateLifyCycleHookResponseDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLifyCycleHookResponseDefaultResult) UnmarshalJSON(b []byte) error {
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
