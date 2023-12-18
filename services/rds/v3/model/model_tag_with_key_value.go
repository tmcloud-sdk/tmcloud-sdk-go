package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TagWithKeyValue struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o TagWithKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagWithKeyValue struct{}"
	}

	return strings.Join([]string{"TagWithKeyValue", string(data)}, " ")
}
