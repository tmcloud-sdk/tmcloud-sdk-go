package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListScalingTagInfosByTenantIdRequest struct {
	ResourceType ListScalingTagInfosByTenantIdRequestResourceType `json:"resource_type"`
}

func (o ListScalingTagInfosByTenantIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingTagInfosByTenantIdRequest struct{}"
	}

	return strings.Join([]string{"ListScalingTagInfosByTenantIdRequest", string(data)}, " ")
}

type ListScalingTagInfosByTenantIdRequestResourceType struct {
	value string
}

type ListScalingTagInfosByTenantIdRequestResourceTypeEnum struct {
	SCALING_GROUP_TAG ListScalingTagInfosByTenantIdRequestResourceType
}

func GetListScalingTagInfosByTenantIdRequestResourceTypeEnum() ListScalingTagInfosByTenantIdRequestResourceTypeEnum {
	return ListScalingTagInfosByTenantIdRequestResourceTypeEnum{
		SCALING_GROUP_TAG: ListScalingTagInfosByTenantIdRequestResourceType{
			value: "scaling_group_tag",
		},
	}
}

func (c ListScalingTagInfosByTenantIdRequestResourceType) Value() string {
	return c.value
}

func (c ListScalingTagInfosByTenantIdRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingTagInfosByTenantIdRequestResourceType) UnmarshalJSON(b []byte) error {
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
