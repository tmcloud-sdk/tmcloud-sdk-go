package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchCreateServerTagsRequestBody struct {
	Action BatchCreateServerTagsRequestBodyAction `json:"action"`

	Tags []ServerTag `json:"tags"`
}

func (o BatchCreateServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsRequestBody", string(data)}, " ")
}

type BatchCreateServerTagsRequestBodyAction struct {
	value string
}

type BatchCreateServerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateServerTagsRequestBodyAction
}

func GetBatchCreateServerTagsRequestBodyActionEnum() BatchCreateServerTagsRequestBodyActionEnum {
	return BatchCreateServerTagsRequestBodyActionEnum{
		CREATE: BatchCreateServerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
