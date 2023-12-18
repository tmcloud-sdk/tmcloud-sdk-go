package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerFlavor struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Disk string `json:"disk"`

	Vcpus string `json:"vcpus"`

	Ram string `json:"ram"`
}

func (o ServerFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerFlavor struct{}"
	}

	return strings.Join([]string{"ServerFlavor", string(data)}, " ")
}
