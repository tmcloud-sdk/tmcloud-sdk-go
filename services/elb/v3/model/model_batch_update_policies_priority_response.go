package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdatePoliciesPriorityResponse struct {
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdatePoliciesPriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityResponse", string(data)}, " ")
}
