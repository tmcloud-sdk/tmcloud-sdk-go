package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type EncryptDataRequestBody struct {
	KeyId string `json:"key_id"`

	PlainText string `json:"plain_text"`

	EncryptionAlgorithm *EncryptDataRequestBodyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o EncryptDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataRequestBody struct{}"
	}

	return strings.Join([]string{"EncryptDataRequestBody", string(data)}, " ")
}

type EncryptDataRequestBodyEncryptionAlgorithm struct {
	value string
}

type EncryptDataRequestBodyEncryptionAlgorithmEnum struct {
	SYMMETRIC_DEFAULT  EncryptDataRequestBodyEncryptionAlgorithm
	RSAES_OAEP_SHA_256 EncryptDataRequestBodyEncryptionAlgorithm
	SM2_ENCRYPT        EncryptDataRequestBodyEncryptionAlgorithm
}

func GetEncryptDataRequestBodyEncryptionAlgorithmEnum() EncryptDataRequestBodyEncryptionAlgorithmEnum {
	return EncryptDataRequestBodyEncryptionAlgorithmEnum{
		SYMMETRIC_DEFAULT: EncryptDataRequestBodyEncryptionAlgorithm{
			value: "SYMMETRIC_DEFAULT",
		},
		RSAES_OAEP_SHA_256: EncryptDataRequestBodyEncryptionAlgorithm{
			value: "RSAES_OAEP_SHA_256",
		},
		SM2_ENCRYPT: EncryptDataRequestBodyEncryptionAlgorithm{
			value: "SM2_ENCRYPT",
		},
	}
}

func (c EncryptDataRequestBodyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c EncryptDataRequestBodyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncryptDataRequestBodyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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
