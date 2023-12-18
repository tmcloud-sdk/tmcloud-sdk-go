package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteKeyResponse struct {
	KeyId *string `json:"key_id,omitempty"`

	KeyState       *string `json:"key_state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeyResponse", string(data)}, " ")
}
