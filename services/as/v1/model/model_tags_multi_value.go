package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TagsMultiValue struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagsMultiValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsMultiValue struct{}"
	}

	return strings.Join([]string{"TagsMultiValue", string(data)}, " ")
}
