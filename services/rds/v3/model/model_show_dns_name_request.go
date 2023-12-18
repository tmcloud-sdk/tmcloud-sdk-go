package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDnsNameRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *string `json:"X-Language,omitempty"`

	DnsType string `json:"dns_type"`
}

func (o ShowDnsNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDnsNameRequest struct{}"
	}

	return strings.Join([]string{"ShowDnsNameRequest", string(data)}, " ")
}
