package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Context struct {
	Cluster *string `json:"cluster,omitempty"`

	User *string `json:"user,omitempty"`
}

func (o Context) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Context struct{}"
	}

	return strings.Join([]string{"Context", string(data)}, " ")
}
