package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteScalingTagInfoRequest struct {
	ResourceType DeleteScalingTagInfoRequestResourceType `json:"resource_type"`

	ResourceId string `json:"resource_id"`

	Body *DeleteTagsOption `json:"body,omitempty"`
}

func (o DeleteScalingTagInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingTagInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingTagInfoRequest", string(data)}, " ")
}

type DeleteScalingTagInfoRequestResourceType struct {
	value string
}

type DeleteScalingTagInfoRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG DeleteScalingTagInfoRequestResourceType
}

func GetDeleteScalingTagInfoRequestResourceTypeEnum() DeleteScalingTagInfoRequestResourceTypeEnum {
	return DeleteScalingTagInfoRequestResourceTypeEnum{
		SCALING_GROUP_TAG: DeleteScalingTagInfoRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c DeleteScalingTagInfoRequestResourceType) Value() string {
	return c.value
}

func (c DeleteScalingTagInfoRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteScalingTagInfoRequestResourceType) UnmarshalJSON(b []byte) error {
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
