package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SimpleFlavor struct {
	Id string `json:"id"`

	Links []Link `json:"links"`
}

func (o SimpleFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleFlavor struct{}"
	}

	return strings.Join([]string{"SimpleFlavor", string(data)}, " ")
}
