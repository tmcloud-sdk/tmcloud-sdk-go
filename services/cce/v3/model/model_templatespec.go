package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Templatespec struct {
	Type string `json:"type"`

	Require *bool `json:"require,omitempty"`

	Labels []string `json:"labels"`

	LogoURL string `json:"logoURL"`

	ReadmeURL string `json:"readmeURL"`

	Description string `json:"description"`

	Versions []Versions `json:"versions"`
}

func (o Templatespec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Templatespec struct{}"
	}

	return strings.Join([]string{"Templatespec", string(data)}, " ")
}
