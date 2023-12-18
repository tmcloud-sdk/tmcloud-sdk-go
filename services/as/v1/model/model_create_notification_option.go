package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateNotificationOption struct {
	TopicUrn string `json:"topic_urn"`

	TopicScene []CreateNotificationOptionTopicScene `json:"topic_scene"`
}

func (o CreateNotificationOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationOption struct{}"
	}

	return strings.Join([]string{"CreateNotificationOption", string(data)}, " ")
}

type CreateNotificationOptionTopicScene struct {
	value string
}

type CreateNotificationOptionTopicSceneEnum struct {
	SCALING_UP             CreateNotificationOptionTopicScene
	SCALING_UP_FAIL        CreateNotificationOptionTopicScene
	SCALING_DOWN           CreateNotificationOptionTopicScene
	SCALING_DOWN_FAIL      CreateNotificationOptionTopicScene
	SCALING_GROUP_ABNORMAL CreateNotificationOptionTopicScene
}

func GetCreateNotificationOptionTopicSceneEnum() CreateNotificationOptionTopicSceneEnum {
	return CreateNotificationOptionTopicSceneEnum{
		SCALING_UP: CreateNotificationOptionTopicScene{
			value: "SCALING_UP",
		},
		SCALING_UP_FAIL: CreateNotificationOptionTopicScene{
			value: "SCALING_UP_FAIL",
		},
		SCALING_DOWN: CreateNotificationOptionTopicScene{
			value: "SCALING_DOWN",
		},
		SCALING_DOWN_FAIL: CreateNotificationOptionTopicScene{
			value: "SCALING_DOWN_FAIL",
		},
		SCALING_GROUP_ABNORMAL: CreateNotificationOptionTopicScene{
			value: "SCALING_GROUP_ABNORMAL",
		},
	}
}

func (c CreateNotificationOptionTopicScene) Value() string {
	return c.value
}

func (c CreateNotificationOptionTopicScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationOptionTopicScene) UnmarshalJSON(b []byte) error {
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
