package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SyncPolicyResp struct {
	Id *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o SyncPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyResp struct{}"
	}

	return strings.Join([]string{"SyncPolicyResp", string(data)}, " ")
}
