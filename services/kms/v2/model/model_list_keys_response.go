package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListKeysResponse struct {
	Keys *[]string `json:"keys,omitempty"`

	KeyDetails *[]KeyDetails `json:"key_details,omitempty"`

	NextMarker *string `json:"next_marker,omitempty"`

	Truncated *ListKeysResponseTruncated `json:"truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeysResponse struct{}"
	}

	return strings.Join([]string{"ListKeysResponse", string(data)}, " ")
}

type ListKeysResponseTruncated struct {
	value string
}

type ListKeysResponseTruncatedEnum struct {
	TRUE  ListKeysResponseTruncated
	FALSE ListKeysResponseTruncated
}

func GetListKeysResponseTruncatedEnum() ListKeysResponseTruncatedEnum {
	return ListKeysResponseTruncatedEnum{
		TRUE: ListKeysResponseTruncated{
			value: "true",
		},
		FALSE: ListKeysResponseTruncated{
			value: "false",
		},
	}
}

func (c ListKeysResponseTruncated) Value() string {
	return c.value
}

func (c ListKeysResponseTruncated) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKeysResponseTruncated) UnmarshalJSON(b []byte) error {
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
