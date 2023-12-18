package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListVpcsByTagsRequestBody struct {
	Action ListVpcsByTagsRequestBodyAction `json:"action"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`

	Tags *[]ListTag `json:"tags,omitempty"`
}

func (o ListVpcsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListVpcsByTagsRequestBody", string(data)}, " ")
}

type ListVpcsByTagsRequestBodyAction struct {
	value string
}

type ListVpcsByTagsRequestBodyActionEnum struct {
	FILTER ListVpcsByTagsRequestBodyAction
	COUNT  ListVpcsByTagsRequestBodyAction
}

func GetListVpcsByTagsRequestBodyActionEnum() ListVpcsByTagsRequestBodyActionEnum {
	return ListVpcsByTagsRequestBodyActionEnum{
		FILTER: ListVpcsByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListVpcsByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListVpcsByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListVpcsByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVpcsByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
