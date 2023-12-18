package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateDnsNameRequestBody struct {
	DnsType string `json:"dns_type"`
}

func (o CreateDnsNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDnsNameRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDnsNameRequestBody", string(data)}, " ")
}
