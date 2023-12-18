package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowResourceQuotaResponse struct {
	Quotas         *AllQuotas `json:"quotas,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowResourceQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceQuotaResponse", string(data)}, " ")
}
