package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EnableKeyResponse struct {
	KeyInfo        *KeyStatusInfo `json:"key_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o EnableKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyResponse struct{}"
	}

	return strings.Join([]string{"EnableKeyResponse", string(data)}, " ")
}
