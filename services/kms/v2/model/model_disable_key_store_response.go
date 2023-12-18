package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DisableKeyStoreResponse struct {
	Keystore       *KeyStoreStateInfo `json:"keystore,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o DisableKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyStoreResponse", string(data)}, " ")
}
