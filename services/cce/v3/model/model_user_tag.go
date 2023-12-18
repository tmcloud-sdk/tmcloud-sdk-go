package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UserTag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

func (o UserTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTag struct{}"
	}

	return strings.Join([]string{"UserTag", string(data)}, " ")
}
