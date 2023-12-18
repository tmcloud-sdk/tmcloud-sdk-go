package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryTagsOption struct {
	Tags *[]TagsMultiValue `json:"tags,omitempty"`

	TagsAny *[]TagsMultiValue `json:"tags_any,omitempty"`

	NotTags *[]TagsMultiValue `json:"not_tags,omitempty"`

	NotTagsAny *[]TagsMultiValue `json:"not_tags_any,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Action QueryTagsOptionAction `json:"action"`

	Offset *string `json:"offset,omitempty"`

	Matches *[]Matches `json:"matches,omitempty"`
}

func (o QueryTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTagsOption struct{}"
	}

	return strings.Join([]string{"QueryTagsOption", string(data)}, " ")
}

type QueryTagsOptionAction struct {
	value string
}

type QueryTagsOptionActionEnum struct {
	FILTER QueryTagsOptionAction
	COUNT  QueryTagsOptionAction
}

func GetQueryTagsOptionActionEnum() QueryTagsOptionActionEnum {
	return QueryTagsOptionActionEnum{
		FILTER: QueryTagsOptionAction{
			value: "filter",
		},
		COUNT: QueryTagsOptionAction{
			value: "count",
		},
	}
}

func (c QueryTagsOptionAction) Value() string {
	return c.value
}

func (c QueryTagsOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryTagsOptionAction) UnmarshalJSON(b []byte) error {
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
