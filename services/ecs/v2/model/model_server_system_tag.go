package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerSystemTag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o ServerSystemTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerSystemTag struct{}"
	}

	return strings.Join([]string{"ServerSystemTag", string(data)}, " ")
}
