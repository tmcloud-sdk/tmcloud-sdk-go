package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListRetirableGrantsResponse struct {
	Grants *[]Grants `json:"grants,omitempty"`

	NextMarker *string `json:"next_marker,omitempty"`

	Truncated      *ListRetirableGrantsResponseTruncated `json:"truncated,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListRetirableGrantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsResponse struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsResponse", string(data)}, " ")
}

type ListRetirableGrantsResponseTruncated struct {
	value string
}

type ListRetirableGrantsResponseTruncatedEnum struct {
	TRUE  ListRetirableGrantsResponseTruncated
	FALSE ListRetirableGrantsResponseTruncated
}

func GetListRetirableGrantsResponseTruncatedEnum() ListRetirableGrantsResponseTruncatedEnum {
	return ListRetirableGrantsResponseTruncatedEnum{
		TRUE: ListRetirableGrantsResponseTruncated{
			value: "true",
		},
		FALSE: ListRetirableGrantsResponseTruncated{
			value: "false",
		},
	}
}

func (c ListRetirableGrantsResponseTruncated) Value() string {
	return c.value
}

func (c ListRetirableGrantsResponseTruncated) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRetirableGrantsResponseTruncated) UnmarshalJSON(b []byte) error {
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
