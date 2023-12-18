package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type KeyDetails struct {
	KeyId *string `json:"key_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	KeyAlias *string `json:"key_alias,omitempty"`

	Realm *string `json:"realm,omitempty"`

	KeySpec *KeyDetailsKeySpec `json:"key_spec,omitempty"`

	KeyUsage *KeyDetailsKeyUsage `json:"key_usage,omitempty"`

	KeyDescription *string `json:"key_description,omitempty"`

	CreationDate *string `json:"creation_date,omitempty"`

	ScheduledDeletionDate *string `json:"scheduled_deletion_date,omitempty"`

	KeyState *string `json:"key_state,omitempty"`

	DefaultKeyFlag *string `json:"default_key_flag,omitempty"`

	KeyType *string `json:"key_type,omitempty"`

	ExpirationTime *string `json:"expiration_time,omitempty"`

	Origin *KeyDetailsOrigin `json:"origin,omitempty"`

	KeyRotationEnabled *string `json:"key_rotation_enabled,omitempty"`

	SysEnterpriseProjectId *string `json:"sys_enterprise_project_id,omitempty"`

	KeystoreId *string `json:"keystore_id,omitempty"`

	KeyLabel *string `json:"key_label,omitempty"`
}

func (o KeyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyDetails struct{}"
	}

	return strings.Join([]string{"KeyDetails", string(data)}, " ")
}

type KeyDetailsKeySpec struct {
	value string
}

type KeyDetailsKeySpecEnum struct {
	AES_256  KeyDetailsKeySpec
	SM4      KeyDetailsKeySpec
	RSA_2048 KeyDetailsKeySpec
	RSA_3072 KeyDetailsKeySpec
	RSA_4096 KeyDetailsKeySpec
	EC_P256  KeyDetailsKeySpec
	EC_P384  KeyDetailsKeySpec
	SM2      KeyDetailsKeySpec
}

func GetKeyDetailsKeySpecEnum() KeyDetailsKeySpecEnum {
	return KeyDetailsKeySpecEnum{
		AES_256: KeyDetailsKeySpec{
			value: "AES_256",
		},
		SM4: KeyDetailsKeySpec{
			value: "SM4",
		},
		RSA_2048: KeyDetailsKeySpec{
			value: "RSA_2048",
		},
		RSA_3072: KeyDetailsKeySpec{
			value: "RSA_3072",
		},
		RSA_4096: KeyDetailsKeySpec{
			value: "RSA_4096",
		},
		EC_P256: KeyDetailsKeySpec{
			value: "EC_P256",
		},
		EC_P384: KeyDetailsKeySpec{
			value: "EC_P384",
		},
		SM2: KeyDetailsKeySpec{
			value: "SM2",
		},
	}
}

func (c KeyDetailsKeySpec) Value() string {
	return c.value
}

func (c KeyDetailsKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyDetailsKeySpec) UnmarshalJSON(b []byte) error {
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

type KeyDetailsKeyUsage struct {
	value string
}

type KeyDetailsKeyUsageEnum struct {
	ENCRYPT_DECRYPT KeyDetailsKeyUsage
	SIGN_VERIFY     KeyDetailsKeyUsage
}

func GetKeyDetailsKeyUsageEnum() KeyDetailsKeyUsageEnum {
	return KeyDetailsKeyUsageEnum{
		ENCRYPT_DECRYPT: KeyDetailsKeyUsage{
			value: "ENCRYPT_DECRYPT",
		},
		SIGN_VERIFY: KeyDetailsKeyUsage{
			value: "SIGN_VERIFY",
		},
	}
}

func (c KeyDetailsKeyUsage) Value() string {
	return c.value
}

func (c KeyDetailsKeyUsage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyDetailsKeyUsage) UnmarshalJSON(b []byte) error {
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

type KeyDetailsOrigin struct {
	value string
}

type KeyDetailsOriginEnum struct {
	KMS      KeyDetailsOrigin
	EXTERNAL KeyDetailsOrigin
}

func GetKeyDetailsOriginEnum() KeyDetailsOriginEnum {
	return KeyDetailsOriginEnum{
		KMS: KeyDetailsOrigin{
			value: "kms",
		},
		EXTERNAL: KeyDetailsOrigin{
			value: "external",
		},
	}
}

func (c KeyDetailsOrigin) Value() string {
	return c.value
}

func (c KeyDetailsOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeyDetailsOrigin) UnmarshalJSON(b []byte) error {
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
