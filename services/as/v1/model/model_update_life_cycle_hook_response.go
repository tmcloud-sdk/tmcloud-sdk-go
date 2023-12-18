package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateLifeCycleHookResponse struct {
	LifecycleHookName *string `json:"lifecycle_hook_name,omitempty"`

	LifecycleHookType *UpdateLifeCycleHookResponseLifecycleHookType `json:"lifecycle_hook_type,omitempty"`

	DefaultResult *UpdateLifeCycleHookResponseDefaultResult `json:"default_result,omitempty"`

	DefaultTimeout *int32 `json:"default_timeout,omitempty"`

	NotificationTopicUrn *string `json:"notification_topic_urn,omitempty"`

	NotificationTopicName *string `json:"notification_topic_name,omitempty"`

	NotificationMetadata *string `json:"notification_metadata,omitempty"`

	CreateTime     *string `json:"create_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLifeCycleHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLifeCycleHookResponse struct{}"
	}

	return strings.Join([]string{"UpdateLifeCycleHookResponse", string(data)}, " ")
}

type UpdateLifeCycleHookResponseLifecycleHookType struct {
	value string
}

type UpdateLifeCycleHookResponseLifecycleHookTypeEnum struct {
	INSTANCE_TERMINATING UpdateLifeCycleHookResponseLifecycleHookType
	INSTANCE_LAUNCHING   UpdateLifeCycleHookResponseLifecycleHookType
}

func GetUpdateLifeCycleHookResponseLifecycleHookTypeEnum() UpdateLifeCycleHookResponseLifecycleHookTypeEnum {
	return UpdateLifeCycleHookResponseLifecycleHookTypeEnum{
		INSTANCE_TERMINATING: UpdateLifeCycleHookResponseLifecycleHookType{
			value: "INSTANCE_TERMINATING",
		},
		INSTANCE_LAUNCHING: UpdateLifeCycleHookResponseLifecycleHookType{
			value: "INSTANCE_LAUNCHING",
		},
	}
}

func (c UpdateLifeCycleHookResponseLifecycleHookType) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookResponseLifecycleHookType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookResponseLifecycleHookType) UnmarshalJSON(b []byte) error {
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

type UpdateLifeCycleHookResponseDefaultResult struct {
	value string
}

type UpdateLifeCycleHookResponseDefaultResultEnum struct {
	ABANDON  UpdateLifeCycleHookResponseDefaultResult
	CONTINUE UpdateLifeCycleHookResponseDefaultResult
}

func GetUpdateLifeCycleHookResponseDefaultResultEnum() UpdateLifeCycleHookResponseDefaultResultEnum {
	return UpdateLifeCycleHookResponseDefaultResultEnum{
		ABANDON: UpdateLifeCycleHookResponseDefaultResult{
			value: "ABANDON",
		},
		CONTINUE: UpdateLifeCycleHookResponseDefaultResult{
			value: "CONTINUE",
		},
	}
}

func (c UpdateLifeCycleHookResponseDefaultResult) Value() string {
	return c.value
}

func (c UpdateLifeCycleHookResponseDefaultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLifeCycleHookResponseDefaultResult) UnmarshalJSON(b []byte) error {
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
