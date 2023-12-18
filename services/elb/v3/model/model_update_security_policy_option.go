package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateSecurityPolicyOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Protocols *[]string `json:"protocols,omitempty"`

	Ciphers *[]UpdateSecurityPolicyOptionCiphers `json:"ciphers,omitempty"`
}

func (o UpdateSecurityPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyOption", string(data)}, " ")
}

type UpdateSecurityPolicyOptionCiphers struct {
	value string
}

type UpdateSecurityPolicyOptionCiphersEnum struct {
	ECDHE_RSA_AES256_GCM_SHA384   UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_GCM_SHA256   UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_GCM_SHA384 UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_GCM_SHA256 UpdateSecurityPolicyOptionCiphers
	AES128_GCM_SHA256             UpdateSecurityPolicyOptionCiphers
	AES256_GCM_SHA384             UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA256     UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA256       UpdateSecurityPolicyOptionCiphers
	AES128_SHA256                 UpdateSecurityPolicyOptionCiphers
	AES256_SHA256                 UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA384     UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA384       UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES128_SHA        UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES128_SHA          UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_AES256_SHA          UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_AES256_SHA        UpdateSecurityPolicyOptionCiphers
	AES128_SHA                    UpdateSecurityPolicyOptionCiphers
	AES256_SHA                    UpdateSecurityPolicyOptionCiphers
	CAMELLIA128_SHA               UpdateSecurityPolicyOptionCiphers
	DES_CBC3_SHA                  UpdateSecurityPolicyOptionCiphers
	CAMELLIA256_SHA               UpdateSecurityPolicyOptionCiphers
	ECDHE_RSA_CHACHA20_POLY1305   UpdateSecurityPolicyOptionCiphers
	ECDHE_ECDSA_CHACHA20_POLY1305 UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_GCM_SHA256        UpdateSecurityPolicyOptionCiphers
	TLS_AES_256_GCM_SHA384        UpdateSecurityPolicyOptionCiphers
	TLS_CHACHA20_POLY1305_SHA256  UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_SHA256        UpdateSecurityPolicyOptionCiphers
	TLS_AES_128_CCM_8_SHA256      UpdateSecurityPolicyOptionCiphers
}

func GetUpdateSecurityPolicyOptionCiphersEnum() UpdateSecurityPolicyOptionCiphersEnum {
	return UpdateSecurityPolicyOptionCiphersEnum{
		ECDHE_RSA_AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-GCM-SHA384",
		},
		ECDHE_RSA_AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-GCM-SHA256",
		},
		ECDHE_ECDSA_AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-GCM-SHA256",
		},
		AES128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-GCM-SHA256",
		},
		AES256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-GCM-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA256",
		},
		ECDHE_RSA_AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA256",
		},
		AES128_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-SHA256",
		},
		AES256_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-SHA256",
		},
		ECDHE_ECDSA_AES256_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA384",
		},
		ECDHE_RSA_AES256_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA384",
		},
		ECDHE_ECDSA_AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES128-SHA",
		},
		ECDHE_RSA_AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES128-SHA",
		},
		ECDHE_RSA_AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-AES256-SHA",
		},
		ECDHE_ECDSA_AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-AES256-SHA",
		},
		AES128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "AES128-SHA",
		},
		AES256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "AES256-SHA",
		},
		CAMELLIA128_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "CAMELLIA128-SHA",
		},
		DES_CBC3_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "DES-CBC3-SHA",
		},
		CAMELLIA256_SHA: UpdateSecurityPolicyOptionCiphers{
			value: "CAMELLIA256-SHA",
		},
		ECDHE_RSA_CHACHA20_POLY1305: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-RSA-CHACHA20-POLY1305",
		},
		ECDHE_ECDSA_CHACHA20_POLY1305: UpdateSecurityPolicyOptionCiphers{
			value: "ECDHE-ECDSA-CHACHA20-POLY1305",
		},
		TLS_AES_128_GCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_GCM_SHA256",
		},
		TLS_AES_256_GCM_SHA384: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_256_GCM_SHA384",
		},
		TLS_CHACHA20_POLY1305_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_CHACHA20_POLY1305_SHA256",
		},
		TLS_AES_128_CCM_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_SHA256",
		},
		TLS_AES_128_CCM_8_SHA256: UpdateSecurityPolicyOptionCiphers{
			value: "TLS_AES_128_CCM_8_SHA256",
		},
	}
}

func (c UpdateSecurityPolicyOptionCiphers) Value() string {
	return c.value
}

func (c UpdateSecurityPolicyOptionCiphers) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPolicyOptionCiphers) UnmarshalJSON(b []byte) error {
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
