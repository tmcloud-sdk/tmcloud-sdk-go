package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerTag struct {
	Key string `json:"key"`

	Value *string `json:"value,omitempty"`
}

func (o ServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerTag struct{}"
	}

	return strings.Join([]string{"ServerTag", string(data)}, " ")
}
