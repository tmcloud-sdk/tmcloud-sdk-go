package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowLogFile struct {
	FileName string `json:"file_name"`

	FileSize string `json:"file_size"`
}

func (o SlowLogFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogFile struct{}"
	}

	return strings.Join([]string{"SlowLogFile", string(data)}, " ")
}
