package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UserPassword struct {
	Username *string `json:"username,omitempty"`

	Password string `json:"password"`
}

func (o UserPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPassword struct{}"
	}

	return strings.Join([]string{"UserPassword", string(data)}, " ")
}
