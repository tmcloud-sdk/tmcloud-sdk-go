package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GetParametersForImportRequestBody struct {
	KeyId string `json:"key_id"`

	WrappingAlgorithm GetParametersForImportRequestBodyWrappingAlgorithm `json:"wrapping_algorithm"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o GetParametersForImportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetParametersForImportRequestBody struct{}"
	}

	return strings.Join([]string{"GetParametersForImportRequestBody", string(data)}, " ")
}

type GetParametersForImportRequestBodyWrappingAlgorithm struct {
	value string
}

type GetParametersForImportRequestBodyWrappingAlgorithmEnum struct {
	RSAES_OAEP_SHA_256 GetParametersForImportRequestBodyWrappingAlgorithm
	SM2_ENCRYPT        GetParametersForImportRequestBodyWrappingAlgorithm
}

func GetGetParametersForImportRequestBodyWrappingAlgorithmEnum() GetParametersForImportRequestBodyWrappingAlgorithmEnum {
	return GetParametersForImportRequestBodyWrappingAlgorithmEnum{
		RSAES_OAEP_SHA_256: GetParametersForImportRequestBodyWrappingAlgorithm{
			value: "RSAES_OAEP_SHA_256",
		},
		SM2_ENCRYPT: GetParametersForImportRequestBodyWrappingAlgorithm{
			value: "SM2_ENCRYPT",
		},
	}
}

func (c GetParametersForImportRequestBodyWrappingAlgorithm) Value() string {
	return c.value
}

func (c GetParametersForImportRequestBodyWrappingAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetParametersForImportRequestBodyWrappingAlgorithm) UnmarshalJSON(b []byte) error {
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
