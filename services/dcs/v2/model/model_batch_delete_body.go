package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteBody struct {
	Instances *[]string `json:"instances,omitempty"`
}

func (o BatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBody", string(data)}, " ")
}
