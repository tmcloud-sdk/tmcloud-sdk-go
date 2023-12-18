package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RenameCommandResp struct {
	Command *string `json:"command,omitempty"`

	Flushall *string `json:"flushall,omitempty"`

	Flushdb *string `json:"flushdb,omitempty"`

	Hgetall *string `json:"hgetall,omitempty"`

	Keys *string `json:"keys,omitempty"`
}

func (o RenameCommandResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameCommandResp struct{}"
	}

	return strings.Join([]string{"RenameCommandResp", string(data)}, " ")
}
