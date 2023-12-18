package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetBinlogClearPolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *BinlogClearPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBinlogClearPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBinlogClearPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBinlogClearPolicyRequest", string(data)}, " ")
}
