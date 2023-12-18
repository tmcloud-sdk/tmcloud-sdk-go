package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowBinlogClearPolicyResponse struct {
	BinlogRetentionHours *int32 `json:"binlog_retention_hours,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowBinlogClearPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogClearPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowBinlogClearPolicyResponse", string(data)}, " ")
}
