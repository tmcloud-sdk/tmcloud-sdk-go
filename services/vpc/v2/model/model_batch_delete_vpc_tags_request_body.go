package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteVpcTagsRequestBody struct {
	Action BatchDeleteVpcTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteVpcTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpcTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpcTagsRequestBody", string(data)}, " ")
}

type BatchDeleteVpcTagsRequestBodyAction struct {
	value string
}

type BatchDeleteVpcTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteVpcTagsRequestBodyAction
}

func GetBatchDeleteVpcTagsRequestBodyActionEnum() BatchDeleteVpcTagsRequestBodyActionEnum {
	return BatchDeleteVpcTagsRequestBodyActionEnum{
		DELETE: BatchDeleteVpcTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteVpcTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteVpcTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteVpcTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
