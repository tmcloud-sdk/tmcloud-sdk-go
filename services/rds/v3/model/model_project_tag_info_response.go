package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ProjectTagInfoResponse struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o ProjectTagInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTagInfoResponse struct{}"
	}

	return strings.Join([]string{"ProjectTagInfoResponse", string(data)}, " ")
}
