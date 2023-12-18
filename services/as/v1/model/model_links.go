package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Links struct {
	Href *string `json:"href,omitempty"`

	Rel *string `json:"rel,omitempty"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
