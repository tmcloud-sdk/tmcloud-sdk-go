package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DomainNameEntity struct {
	DomainName *string `json:"domain_name,omitempty"`

	IsReadonly *bool `json:"is_readonly,omitempty"`
}

func (o DomainNameEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNameEntity struct{}"
	}

	return strings.Join([]string{"DomainNameEntity", string(data)}, " ")
}
