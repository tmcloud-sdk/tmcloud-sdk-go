package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaServerImage struct {
	Id string `json:"id"`

	Links []NovaLink `json:"links"`
}

func (o NovaServerImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerImage struct{}"
	}

	return strings.Join([]string{"NovaServerImage", string(data)}, " ")
}
