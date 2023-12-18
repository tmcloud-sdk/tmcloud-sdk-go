package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FlavorLink struct {
	Href string `json:"href"`

	Rel string `json:"rel"`

	Type string `json:"type"`
}

func (o FlavorLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorLink struct{}"
	}

	return strings.Join([]string{"FlavorLink", string(data)}, " ")
}
