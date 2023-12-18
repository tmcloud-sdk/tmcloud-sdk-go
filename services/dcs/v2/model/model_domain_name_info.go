package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DomainNameInfo struct {
	SupportPublicResolve *bool `json:"support_public_resolve,omitempty"`

	IsLatestRules *bool `json:"is_latest_rules,omitempty"`

	ZoneName *string `json:"zone_name,omitempty"`

	HistoryDomainNames *[]DomainNameEntity `json:"history_domain_names,omitempty"`
}

func (o DomainNameInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNameInfo struct{}"
	}

	return strings.Join([]string{"DomainNameInfo", string(data)}, " ")
}
