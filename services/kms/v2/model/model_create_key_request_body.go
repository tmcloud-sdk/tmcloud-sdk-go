package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateKeyRequestBody struct {
	KeyAlias string `json:"key_alias"`

	KeySpec *CreateKeyRequestBodyKeySpec `json:"key_spec,omitempty"`

	KeyUsage *CreateKeyRequestBodyKeyUsage `json:"key_usage,omitempty"`

	KeyDescription *string `json:"key_description,omitempty"`

	Origin *CreateKeyRequestBodyOrigin `json:"origin,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Sequence *string `json:"sequence,omitempty"`

	KeystoreId *string `json:"keystore_id,omitempty"`
}

func (o CreateKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyRequestBody", string(data)}, " ")
}

type CreateKeyRequestBodyKeySpec struct {
	value string
}

type CreateKeyRequestBodyKeySpecEnum struct {
	AES_256  CreateKeyRequestBodyKeySpec
	SM4      CreateKeyRequestBodyKeySpec
	RSA_2048 CreateKeyRequestBodyKeySpec
	RSA_3072 CreateKeyRequestBodyKeySpec
	RSA_4096 CreateKeyRequestBodyKeySpec
	EC_P256  CreateKeyRequestBodyKeySpec
	EC_P384  CreateKeyRequestBodyKeySpec
	SM2      CreateKeyRequestBodyKeySpec
}

func GetCreateKeyRequestBodyKeySpecEnum() CreateKeyRequestBodyKeySpecEnum {
	return CreateKeyRequestBodyKeySpecEnum{
		AES_256: CreateKeyRequestBodyKeySpec{
			value: "AES_256",
		},
		SM4: CreateKeyRequestBodyKeySpec{
			value: "SM4",
		},
		RSA_2048: CreateKeyRequestBodyKeySpec{
			value: "RSA_2048",
		},
		RSA_3072: CreateKeyRequestBodyKeySpec{
			value: "RSA_3072",
		},
		RSA_4096: CreateKeyRequestBodyKeySpec{
			value: "RSA_4096",
		},
		EC_P256: CreateKeyRequestBodyKeySpec{
			value: "EC_P256",
		},
		EC_P384: CreateKeyRequestBodyKeySpec{
			value: "EC_P384",
		},
		SM2: CreateKeyRequestBodyKeySpec{
			value: "SM2",
		},
	}
}

func (c CreateKeyRequestBodyKeySpec) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyKeySpec) UnmarshalJSON(b []byte) error {
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

type CreateKeyRequestBodyKeyUsage struct {
	value string
}

type CreateKeyRequestBodyKeyUsageEnum struct {
	ENCRYPT_DECRYPT CreateKeyRequestBodyKeyUsage
	SIGN_VERIFY     CreateKeyRequestBodyKeyUsage
}

func GetCreateKeyRequestBodyKeyUsageEnum() CreateKeyRequestBodyKeyUsageEnum {
	return CreateKeyRequestBodyKeyUsageEnum{
		ENCRYPT_DECRYPT: CreateKeyRequestBodyKeyUsage{
			value: "ENCRYPT_DECRYPT",
		},
		SIGN_VERIFY: CreateKeyRequestBodyKeyUsage{
			value: "SIGN_VERIFY",
		},
	}
}

func (c CreateKeyRequestBodyKeyUsage) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyKeyUsage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyKeyUsage) UnmarshalJSON(b []byte) error {
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

type CreateKeyRequestBodyOrigin struct {
	value string
}

type CreateKeyRequestBodyOriginEnum struct {
	KMS      CreateKeyRequestBodyOrigin
	EXTERNAL CreateKeyRequestBodyOrigin
}

func GetCreateKeyRequestBodyOriginEnum() CreateKeyRequestBodyOriginEnum {
	return CreateKeyRequestBodyOriginEnum{
		KMS: CreateKeyRequestBodyOrigin{
			value: "kms",
		},
		EXTERNAL: CreateKeyRequestBodyOrigin{
			value: "external",
		},
	}
}

func (c CreateKeyRequestBodyOrigin) Value() string {
	return c.value
}

func (c CreateKeyRequestBodyOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateKeyRequestBodyOrigin) UnmarshalJSON(b []byte) error {
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
