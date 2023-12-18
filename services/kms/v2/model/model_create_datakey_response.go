package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateDatakeyResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	PlainText *string `json:"plain_text,omitempty"`

	CipherText     *string `json:"cipher_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatakeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyResponse struct{}"
	}

	return strings.Join([]string{"CreateDatakeyResponse", string(data)}, " ")
}
