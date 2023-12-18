package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListSubnetsByTagsRequestBody struct {
	Action ListSubnetsByTagsRequestBodyAction `json:"action"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`

	Tags *[]ListTag `json:"tags,omitempty"`
}

func (o ListSubnetsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsRequestBody", string(data)}, " ")
}

type ListSubnetsByTagsRequestBodyAction struct {
	value string
}

type ListSubnetsByTagsRequestBodyActionEnum struct {
	FILTER ListSubnetsByTagsRequestBodyAction
	COUNT  ListSubnetsByTagsRequestBodyAction
}

func GetListSubnetsByTagsRequestBodyActionEnum() ListSubnetsByTagsRequestBodyActionEnum {
	return ListSubnetsByTagsRequestBodyActionEnum{
		FILTER: ListSubnetsByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListSubnetsByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListSubnetsByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListSubnetsByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubnetsByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
