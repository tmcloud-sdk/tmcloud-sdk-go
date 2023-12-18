package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainNameRequest struct {
	InstanceId string `json:"instance_id"`

	DnsType string `json:"dns_type"`

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainNameRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainNameRequest", string(data)}, " ")
}
