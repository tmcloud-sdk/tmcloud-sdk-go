package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetPolicyResponse struct {
	Results *[]SyncPolicyResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSetPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetPolicyResponse", string(data)}, " ")
}
