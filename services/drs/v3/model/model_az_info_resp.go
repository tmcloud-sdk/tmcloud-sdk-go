package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AzInfoResp struct {
	Code *string `json:"code,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o AzInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfoResp struct{}"
	}

	return strings.Join([]string{"AzInfoResp", string(data)}, " ")
}
