package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type KeystoreDetails struct {
	KeystoreId *string `json:"keystore_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	KeystoreAlias *string `json:"keystore_alias,omitempty"`

	KeystoreType *string `json:"keystore_type,omitempty"`

	HsmClusterId *string `json:"hsm_cluster_id,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`
}

func (o KeystoreDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoreDetails struct{}"
	}

	return strings.Join([]string{"KeystoreDetails", string(data)}, " ")
}
