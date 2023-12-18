package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CertificateInfo struct {
	AdminStateUp bool `json:"admin_state_up"`

	Certificate string `json:"certificate"`

	Description string `json:"description"`

	Domain string `json:"domain"`

	Id string `json:"id"`

	Name string `json:"name"`

	PrivateKey string `json:"private_key"`

	Type string `json:"type"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	ExpireTime string `json:"expire_time"`

	ProjectId string `json:"project_id"`

	EncCertificate *string `json:"enc_certificate,omitempty"`

	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o CertificateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateInfo struct{}"
	}

	return strings.Join([]string{"CertificateInfo", string(data)}, " ")
}
