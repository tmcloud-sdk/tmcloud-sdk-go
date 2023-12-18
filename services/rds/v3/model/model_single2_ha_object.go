package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Single2HaObject struct {
	AzCodeNewNode string `json:"az_code_new_node"`

	DsspoolId *string `json:"dsspool_id,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	AdDomainInfo *AdDomainInfo `json:"ad_domain_info,omitempty"`
}

func (o Single2HaObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Single2HaObject struct{}"
	}

	return strings.Join([]string{"Single2HaObject", string(data)}, " ")
}
