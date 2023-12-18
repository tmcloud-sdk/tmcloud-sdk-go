package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateCertificateOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Certificate string `json:"certificate"`

	Description *string `json:"description,omitempty"`

	Domain *string `json:"domain,omitempty"`

	Name *string `json:"name,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Type *CreateCertificateOptionType `json:"type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EncCertificate *string `json:"enc_certificate,omitempty"`

	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o CreateCertificateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateOption struct{}"
	}

	return strings.Join([]string{"CreateCertificateOption", string(data)}, " ")
}

type CreateCertificateOptionType struct {
	value string
}

type CreateCertificateOptionTypeEnum struct {
	SERVER CreateCertificateOptionType
	CLIENT CreateCertificateOptionType
}

func GetCreateCertificateOptionTypeEnum() CreateCertificateOptionTypeEnum {
	return CreateCertificateOptionTypeEnum{
		SERVER: CreateCertificateOptionType{
			value: "server",
		},
		CLIENT: CreateCertificateOptionType{
			value: "client",
		},
	}
}

func (c CreateCertificateOptionType) Value() string {
	return c.value
}

func (c CreateCertificateOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateOptionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
