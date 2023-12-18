package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListImageByTagsRequestBody struct {
	Action ListImageByTagsRequestBodyAction `json:"action"`

	Tags *[]Tags `json:"tags,omitempty"`

	TagsAny *[]Tags `json:"tags_any,omitempty"`

	NotTags *[]Tags `json:"not_tags,omitempty"`

	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Matches *[]TagKeyValue `json:"matches,omitempty"`

	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o ListImageByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListImageByTagsRequestBody", string(data)}, " ")
}

type ListImageByTagsRequestBodyAction struct {
	value string
}

type ListImageByTagsRequestBodyActionEnum struct {
	FILTER ListImageByTagsRequestBodyAction
	COUNT  ListImageByTagsRequestBodyAction
}

func GetListImageByTagsRequestBodyActionEnum() ListImageByTagsRequestBodyActionEnum {
	return ListImageByTagsRequestBodyActionEnum{
		FILTER: ListImageByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListImageByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListImageByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListImageByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
