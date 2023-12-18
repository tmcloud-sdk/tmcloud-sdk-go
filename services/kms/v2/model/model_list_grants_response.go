package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListGrantsResponse struct {
	Grants *[]Grants `json:"grants,omitempty"`

	NextMarker *string `json:"next_marker,omitempty"`

	Truncated *ListGrantsResponseTruncated `json:"truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGrantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGrantsResponse struct{}"
	}

	return strings.Join([]string{"ListGrantsResponse", string(data)}, " ")
}

type ListGrantsResponseTruncated struct {
	value string
}

type ListGrantsResponseTruncatedEnum struct {
	TRUE  ListGrantsResponseTruncated
	FALSE ListGrantsResponseTruncated
}

func GetListGrantsResponseTruncatedEnum() ListGrantsResponseTruncatedEnum {
	return ListGrantsResponseTruncatedEnum{
		TRUE: ListGrantsResponseTruncated{
			value: "true",
		},
		FALSE: ListGrantsResponseTruncated{
			value: "false",
		},
	}
}

func (c ListGrantsResponseTruncated) Value() string {
	return c.value
}

func (c ListGrantsResponseTruncated) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGrantsResponseTruncated) UnmarshalJSON(b []byte) error {
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
