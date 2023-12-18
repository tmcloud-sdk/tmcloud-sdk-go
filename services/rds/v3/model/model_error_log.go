package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ErrorLog struct {
	Time string `json:"time"`

	Level string `json:"level"`

	Content string `json:"content"`
}

func (o ErrorLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorLog struct{}"
	}

	return strings.Join([]string{"ErrorLog", string(data)}, " ")
}
