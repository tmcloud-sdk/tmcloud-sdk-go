package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Versions struct {
	Id string `json:"id"`

	Links []Link `json:"links"`

	MediaTypes []MediaTypes `json:"media-types"`

	MinVersion *string `json:"min_version,omitempty"`

	Status string `json:"status"`

	Updated string `json:"updated"`

	Version string `json:"version"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
