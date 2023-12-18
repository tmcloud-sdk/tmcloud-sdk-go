package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TargetInstanceBody struct {
	Id string `json:"id"`

	Name *string `json:"name,omitempty"`

	Password *string `json:"password,omitempty"`
}

func (o TargetInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetInstanceBody struct{}"
	}

	return strings.Join([]string{"TargetInstanceBody", string(data)}, " ")
}
