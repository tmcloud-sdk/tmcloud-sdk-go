package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EncryptDatakeyRequestBody struct {
	KeyId string `json:"key_id"`

	PlainText string `json:"plain_text"`

	DatakeyPlainLength string `json:"datakey_plain_length"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o EncryptDatakeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDatakeyRequestBody struct{}"
	}

	return strings.Join([]string{"EncryptDatakeyRequestBody", string(data)}, " ")
}
