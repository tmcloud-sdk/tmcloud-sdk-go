package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateParametersForImportResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	ImportToken *string `json:"import_token,omitempty"`

	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	PublicKey      *string `json:"public_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateParametersForImportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateParametersForImportResponse struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportResponse", string(data)}, " ")
}
