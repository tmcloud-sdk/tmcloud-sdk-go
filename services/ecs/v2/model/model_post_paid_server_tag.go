package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o PostPaidServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerTag struct{}"
	}

	return strings.Join([]string{"PostPaidServerTag", string(data)}, " ")
}
