package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListQuotaDetailsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Quotas         *[]QuotaInfo `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
