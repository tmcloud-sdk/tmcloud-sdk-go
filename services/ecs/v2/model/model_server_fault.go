package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerFault struct {
	Code *int32 `json:"code,omitempty"`

	Created *string `json:"created,omitempty"`

	Message *string `json:"message,omitempty"`

	Details *string `json:"details,omitempty"`
}

func (o ServerFault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerFault struct{}"
	}

	return strings.Join([]string{"ServerFault", string(data)}, " ")
}
