package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKeyStoreRequestBody struct {
	KeystoreAlias string `json:"keystore_alias"`

	HsmClusterId string `json:"hsm_cluster_id"`

	HsmCaCert string `json:"hsm_ca_cert"`
}

func (o CreateKeyStoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreRequestBody", string(data)}, " ")
}
