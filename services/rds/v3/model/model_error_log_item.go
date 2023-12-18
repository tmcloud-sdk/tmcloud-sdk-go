package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ErrorLogItem struct {
	Time *string `json:"time,omitempty"`

	Level *string `json:"level,omitempty"`

	Content *string `json:"content,omitempty"`

	LineNum *string `json:"line_num,omitempty"`
}

func (o ErrorLogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorLogItem struct{}"
	}

	return strings.Join([]string{"ErrorLogItem", string(data)}, " ")
}
