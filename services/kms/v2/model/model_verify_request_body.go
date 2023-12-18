package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type VerifyRequestBody struct {
	KeyId string `json:"key_id"`

	Message string `json:"message"`

	Signature string `json:"signature"`

	SigningAlgorithm VerifyRequestBodySigningAlgorithm `json:"signing_algorithm"`

	MessageType *VerifyRequestBodyMessageType `json:"message_type,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o VerifyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyRequestBody struct{}"
	}

	return strings.Join([]string{"VerifyRequestBody", string(data)}, " ")
}

type VerifyRequestBodySigningAlgorithm struct {
	value string
}

type VerifyRequestBodySigningAlgorithmEnum struct {
	RSASSA_PSS_SHA_256        VerifyRequestBodySigningAlgorithm
	RSASSA_PSS_SHA_384        VerifyRequestBodySigningAlgorithm
	RSASSA_PSS_SHA_512        VerifyRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_256 VerifyRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_384 VerifyRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_512 VerifyRequestBodySigningAlgorithm
	ECDSA_SHA_256             VerifyRequestBodySigningAlgorithm
	ECDSA_SHA_384             VerifyRequestBodySigningAlgorithm
	ECDSA_SHA_512             VerifyRequestBodySigningAlgorithm
	SM2_DSA_SM3               VerifyRequestBodySigningAlgorithm
}

func GetVerifyRequestBodySigningAlgorithmEnum() VerifyRequestBodySigningAlgorithmEnum {
	return VerifyRequestBodySigningAlgorithmEnum{
		RSASSA_PSS_SHA_256: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_256",
		},
		RSASSA_PSS_SHA_384: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_384",
		},
		RSASSA_PSS_SHA_512: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_512",
		},
		RSASSA_PKCS1_V1_5_SHA_256: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_256",
		},
		RSASSA_PKCS1_V1_5_SHA_384: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_384",
		},
		RSASSA_PKCS1_V1_5_SHA_512: VerifyRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_512",
		},
		ECDSA_SHA_256: VerifyRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_256",
		},
		ECDSA_SHA_384: VerifyRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_384",
		},
		ECDSA_SHA_512: VerifyRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_512",
		},
		SM2_DSA_SM3: VerifyRequestBodySigningAlgorithm{
			value: "SM2DSA_SM3",
		},
	}
}

func (c VerifyRequestBodySigningAlgorithm) Value() string {
	return c.value
}

func (c VerifyRequestBodySigningAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VerifyRequestBodySigningAlgorithm) UnmarshalJSON(b []byte) error {
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

type VerifyRequestBodyMessageType struct {
	value string
}

type VerifyRequestBodyMessageTypeEnum struct {
	DIGEST VerifyRequestBodyMessageType
	RAW    VerifyRequestBodyMessageType
}

func GetVerifyRequestBodyMessageTypeEnum() VerifyRequestBodyMessageTypeEnum {
	return VerifyRequestBodyMessageTypeEnum{
		DIGEST: VerifyRequestBodyMessageType{
			value: "DIGEST",
		},
		RAW: VerifyRequestBodyMessageType{
			value: "RAW",
		},
	}
}

func (c VerifyRequestBodyMessageType) Value() string {
	return c.value
}

func (c VerifyRequestBodyMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VerifyRequestBodyMessageType) UnmarshalJSON(b []byte) error {
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
