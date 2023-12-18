package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type KafkaSecurity struct {
	TrustStoreKeyName *string `json:"trust_store_key_name,omitempty"`

	TrustStoreKey *string `json:"trust_store_key,omitempty"`

	TrustStorePassword *string `json:"trust_store_password,omitempty"`

	Type *KafkaSecurityType `json:"type,omitempty"`
}

func (o KafkaSecurity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaSecurity struct{}"
	}

	return strings.Join([]string{"KafkaSecurity", string(data)}, " ")
}

type KafkaSecurityType struct {
	value string
}

type KafkaSecurityTypeEnum struct {
	PLAINTEXT KafkaSecurityType
	SASL_SSL  KafkaSecurityType
}

func GetKafkaSecurityTypeEnum() KafkaSecurityTypeEnum {
	return KafkaSecurityTypeEnum{
		PLAINTEXT: KafkaSecurityType{
			value: "PLAINTEXT",
		},
		SASL_SSL: KafkaSecurityType{
			value: "SASL_SSL",
		},
	}
}

func (c KafkaSecurityType) Value() string {
	return c.value
}

func (c KafkaSecurityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaSecurityType) UnmarshalJSON(b []byte) error {
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
