package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DisableKeyResponse struct {
	KeyInfo        *KeyStatusInfo `json:"key_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DisableKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyResponse", string(data)}, " ")
}
