package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TagsSingleValue struct {
	Key string `json:"key"`

	Value *string `json:"value,omitempty"`
}

func (o TagsSingleValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsSingleValue struct{}"
	}

	return strings.Join([]string{"TagsSingleValue", string(data)}, " ")
}
