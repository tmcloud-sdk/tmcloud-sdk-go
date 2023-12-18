package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SignRequestBody struct {
	KeyId string `json:"key_id"`

	Message string `json:"message"`

	SigningAlgorithm SignRequestBodySigningAlgorithm `json:"signing_algorithm"`

	MessageType *SignRequestBodyMessageType `json:"message_type,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o SignRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignRequestBody struct{}"
	}

	return strings.Join([]string{"SignRequestBody", string(data)}, " ")
}

type SignRequestBodySigningAlgorithm struct {
	value string
}

type SignRequestBodySigningAlgorithmEnum struct {
	RSASSA_PSS_SHA_256        SignRequestBodySigningAlgorithm
	RSASSA_PSS_SHA_384        SignRequestBodySigningAlgorithm
	RSASSA_PSS_SHA_512        SignRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_256 SignRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_384 SignRequestBodySigningAlgorithm
	RSASSA_PKCS1_V1_5_SHA_512 SignRequestBodySigningAlgorithm
	ECDSA_SHA_256             SignRequestBodySigningAlgorithm
	ECDSA_SHA_384             SignRequestBodySigningAlgorithm
	ECDSA_SHA_512             SignRequestBodySigningAlgorithm
	SM2_DSA_SM3               SignRequestBodySigningAlgorithm
}

func GetSignRequestBodySigningAlgorithmEnum() SignRequestBodySigningAlgorithmEnum {
	return SignRequestBodySigningAlgorithmEnum{
		RSASSA_PSS_SHA_256: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_256",
		},
		RSASSA_PSS_SHA_384: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_384",
		},
		RSASSA_PSS_SHA_512: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PSS_SHA_512",
		},
		RSASSA_PKCS1_V1_5_SHA_256: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_256",
		},
		RSASSA_PKCS1_V1_5_SHA_384: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_384",
		},
		RSASSA_PKCS1_V1_5_SHA_512: SignRequestBodySigningAlgorithm{
			value: "RSASSA_PKCS1_V1_5_SHA_512",
		},
		ECDSA_SHA_256: SignRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_256",
		},
		ECDSA_SHA_384: SignRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_384",
		},
		ECDSA_SHA_512: SignRequestBodySigningAlgorithm{
			value: "ECDSA_SHA_512",
		},
		SM2_DSA_SM3: SignRequestBodySigningAlgorithm{
			value: "SM2DSA_SM3",
		},
	}
}

func (c SignRequestBodySigningAlgorithm) Value() string {
	return c.value
}

func (c SignRequestBodySigningAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignRequestBodySigningAlgorithm) UnmarshalJSON(b []byte) error {
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

type SignRequestBodyMessageType struct {
	value string
}

type SignRequestBodyMessageTypeEnum struct {
	DIGEST SignRequestBodyMessageType
	RAW    SignRequestBodyMessageType
}

func GetSignRequestBodyMessageTypeEnum() SignRequestBodyMessageTypeEnum {
	return SignRequestBodyMessageTypeEnum{
		DIGEST: SignRequestBodyMessageType{
			value: "DIGEST",
		},
		RAW: SignRequestBodyMessageType{
			value: "RAW",
		},
	}
}

func (c SignRequestBodyMessageType) Value() string {
	return c.value
}

func (c SignRequestBodyMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignRequestBodyMessageType) UnmarshalJSON(b []byte) error {
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
