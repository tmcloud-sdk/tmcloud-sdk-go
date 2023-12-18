package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateCertificateOption struct {
	Certificate *string `json:"certificate,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Domain *string `json:"domain,omitempty"`

	EncCertificate *string `json:"enc_certificate,omitempty"`

	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o UpdateCertificateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateOption struct{}"
	}

	return strings.Join([]string{"UpdateCertificateOption", string(data)}, " ")
}
