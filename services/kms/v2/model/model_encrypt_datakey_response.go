package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EncryptDatakeyResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	CipherText *string `json:"cipher_text,omitempty"`

	DatakeyLength  *string `json:"datakey_length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EncryptDatakeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDatakeyResponse struct{}"
	}

	return strings.Join([]string{"EncryptDatakeyResponse", string(data)}, " ")
}
