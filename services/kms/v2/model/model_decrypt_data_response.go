package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DecryptDataResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	PlainText *string `json:"plain_text,omitempty"`

	PlainTextBase64 *string `json:"plain_text_base64,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o DecryptDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataResponse struct{}"
	}

	return strings.Join([]string{"DecryptDataResponse", string(data)}, " ")
}
