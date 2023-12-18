package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Versions struct {
	Version string `json:"version"`

	Input *interface{} `json:"input"`

	Stable bool `json:"stable"`

	Translate *interface{} `json:"translate"`

	SupportVersions []SupportVersions `json:"supportVersions"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	UpdateTimestamp string `json:"updateTimestamp"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
