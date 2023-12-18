package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ImportKeyMaterialRequestBody struct {
	KeyId string `json:"key_id"`

	ImportToken string `json:"import_token"`

	EncryptedKeyMaterial string `json:"encrypted_key_material"`

	EncryptedPrivatekey *string `json:"encrypted_privatekey,omitempty"`

	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ImportKeyMaterialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportKeyMaterialRequestBody struct{}"
	}

	return strings.Join([]string{"ImportKeyMaterialRequestBody", string(data)}, " ")
}
