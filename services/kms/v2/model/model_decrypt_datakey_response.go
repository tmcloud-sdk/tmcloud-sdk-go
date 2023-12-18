package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DecryptDatakeyResponse struct {
	DataKey *string `json:"data_key,omitempty"`

	DatakeyLength *string `json:"datakey_length,omitempty"`

	DatakeyDgst    *string `json:"datakey_dgst,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DecryptDatakeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyResponse struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyResponse", string(data)}, " ")
}
