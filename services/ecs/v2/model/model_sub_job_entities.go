package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SubJobEntities struct {
	ServerId *string `json:"server_id,omitempty"`

	NicId *string `json:"nic_id,omitempty"`

	ErrorcodeMessage *string `json:"errorcode_message,omitempty"`
}

func (o SubJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobEntities struct{}"
	}

	return strings.Join([]string{"SubJobEntities", string(data)}, " ")
}
