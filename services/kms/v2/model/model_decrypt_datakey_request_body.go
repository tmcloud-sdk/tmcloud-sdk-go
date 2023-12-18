package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DecryptDatakeyRequestBody struct {
	KeyId string `json:"key_id"`

	CipherText string `json:"cipher_text"`

	DatakeyCipherLength string `json:"datakey_cipher_length"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o DecryptDatakeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyRequestBody struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyRequestBody", string(data)}, " ")
}
