package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Files struct {
	FileName string `json:"file_name"`

	Size *string `json:"size,omitempty"`

	UpdateAt *string `json:"update_at,omitempty"`
}

func (o Files) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Files struct{}"
	}

	return strings.Join([]string{"Files", string(data)}, " ")
}
