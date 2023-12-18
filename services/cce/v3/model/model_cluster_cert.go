package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterCert struct {
	Server *string `json:"server,omitempty"`

	CertificateAuthorityData *string `json:"certificate-authority-data,omitempty"`

	InsecureSkipTlsVerify *bool `json:"insecure-skip-tls-verify,omitempty"`
}

func (o ClusterCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCert struct{}"
	}

	return strings.Join([]string{"ClusterCert", string(data)}, " ")
}
