package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListScalingTagInfosByResourceIdRequest struct {
	ResourceType ListScalingTagInfosByResourceIdRequestResourceType `json:"resource_type"`

	ResourceId string `json:"resource_id"`
}

func (o ListScalingTagInfosByResourceIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByResourceIdRequest struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByResourceIdRequest", string(data)}, " ")
}

type ListScalingTagInfosByResourceIdRequestResourceType struct {
	value string
}

type ListScalingTagInfosByResourceIdRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG ListScalingTagInfosByResourceIdRequestResourceType
}

func GetListScalingTagInfosByResourceIdRequestResourceTypeEnum() ListScalingTagInfosByResourceIdRequestResourceTypeEnum {
	return ListScalingTagInfosByResourceIdRequestResourceTypeEnum{
		SCALING_GROUP_TAG: ListScalingTagInfosByResourceIdRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c ListScalingTagInfosByResourceIdRequestResourceType) Value() string {
	return c.value
}

func (c ListScalingTagInfosByResourceIdRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingTagInfosByResourceIdRequestResourceType) UnmarshalJSON(b []byte) error {
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
