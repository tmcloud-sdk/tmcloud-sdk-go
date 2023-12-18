package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowImageQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaResponse", string(data)}, " ")
}
