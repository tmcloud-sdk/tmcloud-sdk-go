package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListVolumesByTagsRequestBody struct {
	Action ListVolumesByTagsRequestBodyAction `json:"action"`

	Limit *int32 `json:"limit,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Tags []TagsForListVolumes `json:"tags"`
}

func (o ListVolumesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsRequestBody", string(data)}, " ")
}

type ListVolumesByTagsRequestBodyAction struct {
	value string
}

type ListVolumesByTagsRequestBodyActionEnum struct {
	FILTER ListVolumesByTagsRequestBodyAction
}

func GetListVolumesByTagsRequestBodyActionEnum() ListVolumesByTagsRequestBodyActionEnum {
	return ListVolumesByTagsRequestBodyActionEnum{
		FILTER: ListVolumesByTagsRequestBodyAction{
			value: "filter",
		},
	}
}

func (c ListVolumesByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListVolumesByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVolumesByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
