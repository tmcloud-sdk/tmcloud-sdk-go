package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateKeyAliasRequestBody struct {
	KeyId string `json:"key_id"`

	KeyAlias string `json:"key_alias"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o UpdateKeyAliasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyAliasRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeyAliasRequestBody", string(data)}, " ")
}
