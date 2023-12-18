package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Features struct {
	SupportAcl *bool `json:"support_acl,omitempty"`

	SupportTransparentClientIp *bool `json:"support_transparent_client_ip,omitempty"`

	SupportSsl *bool `json:"support_ssl,omitempty"`
}

func (o Features) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Features struct{}"
	}

	return strings.Join([]string{"Features", string(data)}, " ")
}
