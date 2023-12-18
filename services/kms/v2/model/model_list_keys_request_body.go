package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListKeysRequestBody struct {
	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	KeyState *string `json:"key_state,omitempty"`

	KeySpec *ListKeysRequestBodyKeySpec `json:"key_spec,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ListKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeysRequestBody struct{}"
	}

	return strings.Join([]string{"ListKeysRequestBody", string(data)}, " ")
}

type ListKeysRequestBodyKeySpec struct {
	value string
}

type ListKeysRequestBodyKeySpecEnum struct {
	AES_256  ListKeysRequestBodyKeySpec
	SM4      ListKeysRequestBodyKeySpec
	RSA_2048 ListKeysRequestBodyKeySpec
	RSA_3072 ListKeysRequestBodyKeySpec
	RSA_4096 ListKeysRequestBodyKeySpec
	EC_P256  ListKeysRequestBodyKeySpec
	EC_P384  ListKeysRequestBodyKeySpec
	SM2      ListKeysRequestBodyKeySpec
	ALL      ListKeysRequestBodyKeySpec
}

func GetListKeysRequestBodyKeySpecEnum() ListKeysRequestBodyKeySpecEnum {
	return ListKeysRequestBodyKeySpecEnum{
		AES_256: ListKeysRequestBodyKeySpec{
			value: "AES_256",
		},
		SM4: ListKeysRequestBodyKeySpec{
			value: "SM4",
		},
		RSA_2048: ListKeysRequestBodyKeySpec{
			value: "RSA_2048",
		},
		RSA_3072: ListKeysRequestBodyKeySpec{
			value: "RSA_3072",
		},
		RSA_4096: ListKeysRequestBodyKeySpec{
			value: "RSA_4096",
		},
		EC_P256: ListKeysRequestBodyKeySpec{
			value: "EC_P256",
		},
		EC_P384: ListKeysRequestBodyKeySpec{
			value: "EC_P384",
		},
		SM2: ListKeysRequestBodyKeySpec{
			value: "SM2",
		},
		ALL: ListKeysRequestBodyKeySpec{
			value: "ALL",
		},
	}
}

func (c ListKeysRequestBodyKeySpec) Value() string {
	return c.value
}

func (c ListKeysRequestBodyKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKeysRequestBodyKeySpec) UnmarshalJSON(b []byte) error {
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
