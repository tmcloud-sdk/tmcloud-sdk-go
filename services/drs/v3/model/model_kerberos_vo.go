package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type KerberosVo struct {
	Krb5ConfFile *string `json:"krb5_conf_file,omitempty"`

	KeyTabFile *string `json:"key_tab_file,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	UserPrincipal *string `json:"user_principal,omitempty"`
}

func (o KerberosVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KerberosVo struct{}"
	}

	return strings.Join([]string{"KerberosVo", string(data)}, " ")
}
