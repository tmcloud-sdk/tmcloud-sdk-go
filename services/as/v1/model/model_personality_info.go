package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PersonalityInfo struct {
	Path string `json:"path"`

	Content string `json:"content"`
}

func (o PersonalityInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersonalityInfo struct{}"
	}

	return strings.Join([]string{"PersonalityInfo", string(data)}, " ")
}
