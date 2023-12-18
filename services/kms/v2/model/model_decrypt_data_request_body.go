package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DecryptDataRequestBody struct {
	CipherText string `json:"cipher_text"`

	EncryptionAlgorithm *DecryptDataRequestBodyEncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	KeyId *string `json:"key_id,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o DecryptDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataRequestBody struct{}"
	}

	return strings.Join([]string{"DecryptDataRequestBody", string(data)}, " ")
}

type DecryptDataRequestBodyEncryptionAlgorithm struct {
	value string
}

type DecryptDataRequestBodyEncryptionAlgorithmEnum struct {
	SYMMETRIC_DEFAULT  DecryptDataRequestBodyEncryptionAlgorithm
	RSAES_OAEP_SHA_256 DecryptDataRequestBodyEncryptionAlgorithm
	SM2_ENCRYPT        DecryptDataRequestBodyEncryptionAlgorithm
}

func GetDecryptDataRequestBodyEncryptionAlgorithmEnum() DecryptDataRequestBodyEncryptionAlgorithmEnum {
	return DecryptDataRequestBodyEncryptionAlgorithmEnum{
		SYMMETRIC_DEFAULT: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "SYMMETRIC_DEFAULT",
		},
		RSAES_OAEP_SHA_256: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "RSAES_OAEP_SHA_256",
		},
		SM2_ENCRYPT: DecryptDataRequestBodyEncryptionAlgorithm{
			value: "SM2_ENCRYPT",
		},
	}
}

func (c DecryptDataRequestBodyEncryptionAlgorithm) Value() string {
	return c.value
}

func (c DecryptDataRequestBodyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DecryptDataRequestBodyEncryptionAlgorithm) UnmarshalJSON(b []byte) error {
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
