package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DisableKeyStoreRequest struct {
	KeystoreId string `json:"keystore_id"`
}

func (o DisableKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"DisableKeyStoreRequest", string(data)}, " ")
}
