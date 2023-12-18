package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type KeystoreInfo struct {
	KeystoreId *string `json:"keystore_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`
}

func (o KeystoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoreInfo struct{}"
	}

	return strings.Join([]string{"KeystoreInfo", string(data)}, " ")
}
