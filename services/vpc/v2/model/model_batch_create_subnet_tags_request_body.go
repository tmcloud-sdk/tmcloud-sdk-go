package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateSubnetTagsRequestBody struct {
	Action BatchCreateSubnetTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateSubnetTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsRequestBody", string(data)}, " ")
}

type BatchCreateSubnetTagsRequestBodyAction struct {
	value string
}

type BatchCreateSubnetTagsRequestBodyActionEnum struct {
	CREATE BatchCreateSubnetTagsRequestBodyAction
}

func GetBatchCreateSubnetTagsRequestBodyActionEnum() BatchCreateSubnetTagsRequestBodyActionEnum {
	return BatchCreateSubnetTagsRequestBodyActionEnum{
		CREATE: BatchCreateSubnetTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateSubnetTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateSubnetTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateSubnetTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
