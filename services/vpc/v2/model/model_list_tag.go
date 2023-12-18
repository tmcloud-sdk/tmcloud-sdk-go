package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListTag struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o ListTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTag struct{}"
	}

	return strings.Join([]string{"ListTag", string(data)}, " ")
}
