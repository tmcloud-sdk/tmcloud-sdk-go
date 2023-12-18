package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Matches struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o Matches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Matches struct{}"
	}

	return strings.Join([]string{"Matches", string(data)}, " ")
}
