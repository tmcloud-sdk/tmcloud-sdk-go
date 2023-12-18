package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowStart struct {
	Enable bool `json:"enable"`

	Duration int32 `json:"duration"`
}

func (o SlowStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowStart struct{}"
	}

	return strings.Join([]string{"SlowStart", string(data)}, " ")
}
