package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DatabaseForCreation struct {
	Name string `json:"name"`

	CharacterSet string `json:"character_set"`

	Comment *string `json:"comment,omitempty"`
}

func (o DatabaseForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseForCreation struct{}"
	}

	return strings.Join([]string{"DatabaseForCreation", string(data)}, " ")
}
