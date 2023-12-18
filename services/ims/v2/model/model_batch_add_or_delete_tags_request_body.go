package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchAddOrDeleteTagsRequestBody struct {
	Action BatchAddOrDeleteTagsRequestBodyAction `json:"action"`

	Tags []ResourceTag `json:"tags"`
}

func (o BatchAddOrDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsRequestBody", string(data)}, " ")
}

type BatchAddOrDeleteTagsRequestBodyAction struct {
	value string
}

type BatchAddOrDeleteTagsRequestBodyActionEnum struct {
	CREATE BatchAddOrDeleteTagsRequestBodyAction
	DELETE BatchAddOrDeleteTagsRequestBodyAction
}

func GetBatchAddOrDeleteTagsRequestBodyActionEnum() BatchAddOrDeleteTagsRequestBodyActionEnum {
	return BatchAddOrDeleteTagsRequestBodyActionEnum{
		CREATE: BatchAddOrDeleteTagsRequestBodyAction{
			value: "create",
		},
		DELETE: BatchAddOrDeleteTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchAddOrDeleteTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchAddOrDeleteTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddOrDeleteTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
