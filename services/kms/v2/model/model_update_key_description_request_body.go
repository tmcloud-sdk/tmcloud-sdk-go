package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateKeyDescriptionRequestBody struct {
	KeyId string `json:"key_id"`

	KeyDescription string `json:"key_description"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o UpdateKeyDescriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyDescriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeyDescriptionRequestBody", string(data)}, " ")
}
