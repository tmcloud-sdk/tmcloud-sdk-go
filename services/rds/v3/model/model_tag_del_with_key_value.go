package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TagDelWithKeyValue struct {
	Key string `json:"key"`

	Value *string `json:"value,omitempty"`
}

func (o TagDelWithKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDelWithKeyValue struct{}"
	}

	return strings.Join([]string{"TagDelWithKeyValue", string(data)}, " ")
}
