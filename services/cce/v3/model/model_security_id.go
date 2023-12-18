package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SecurityId struct {
	Id *string `json:"id,omitempty"`
}

func (o SecurityId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityId struct{}"
	}

	return strings.Join([]string{"SecurityId", string(data)}, " ")
}
