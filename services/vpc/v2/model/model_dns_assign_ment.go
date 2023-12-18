package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DnsAssignMent struct {
	Hostname *string `json:"hostname,omitempty"`

	IpAddress *string `json:"ip_address,omitempty"`

	Fqdn *string `json:"fqdn,omitempty"`
}

func (o DnsAssignMent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsAssignMent struct{}"
	}

	return strings.Join([]string{"DnsAssignMent", string(data)}, " ")
}
