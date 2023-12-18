package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PersonalityResult struct {
	Path *string `json:"path,omitempty"`

	Content *string `json:"content,omitempty"`
}

func (o PersonalityResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersonalityResult struct{}"
	}

	return strings.Join([]string{"PersonalityResult", string(data)}, " ")
}
