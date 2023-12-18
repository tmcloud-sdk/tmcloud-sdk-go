package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateScalingTagInfoRequest struct {
	ResourceType CreateScalingTagInfoRequestResourceType `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Body *CreateTagsOption `json:"body,omitempty"`
}

func (o CreateScalingTagInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingTagInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingTagInfoRequest", string(data)}, " ")
}

type CreateScalingTagInfoRequestResourceType struct {
	value string
}

type CreateScalingTagInfoRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG CreateScalingTagInfoRequestResourceType
}

func GetCreateScalingTagInfoRequestResourceTypeEnum() CreateScalingTagInfoRequestResourceTypeEnum {
	return CreateScalingTagInfoRequestResourceTypeEnum{
		SCALING_GROUP_TAG: CreateScalingTagInfoRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c CreateScalingTagInfoRequestResourceType) Value() string {
	return c.value
}

func (c CreateScalingTagInfoRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingTagInfoRequestResourceType) UnmarshalJSON(b []byte) error {
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
