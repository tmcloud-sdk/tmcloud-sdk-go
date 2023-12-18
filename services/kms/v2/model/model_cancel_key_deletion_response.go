package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CancelKeyDeletionResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	KeyState       *string `json:"key_state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelKeyDeletionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelKeyDeletionResponse struct{}"
	}

	return strings.Join([]string{"CancelKeyDeletionResponse", string(data)}, " ")
}
