package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteServerTagsRequestBody struct {
	Action BatchDeleteServerTagsRequestBodyAction `json:"action"`

	Tags []ServerTag `json:"tags"`
}

func (o BatchDeleteServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteServerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteServerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteServerTagsRequestBodyAction
}

func GetBatchDeleteServerTagsRequestBodyActionEnum() BatchDeleteServerTagsRequestBodyActionEnum {
	return BatchDeleteServerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteServerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
