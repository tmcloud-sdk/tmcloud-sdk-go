package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ValidateSignatureResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	SignatureValid *string `json:"signature_valid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSignatureResponse struct{}"
	}

	return strings.Join([]string{"ValidateSignatureResponse", string(data)}, " ")
}
