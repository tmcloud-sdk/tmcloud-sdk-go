package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type KeyStoreStateInfo struct {
	KeystoreId *string `json:"keystore_id,omitempty"`

	KeystoreState *string `json:"keystore_state,omitempty"`
}

func (o KeyStoreStateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyStoreStateInfo struct{}"
	}

	return strings.Join([]string{"KeyStoreStateInfo", string(data)}, " ")
}
