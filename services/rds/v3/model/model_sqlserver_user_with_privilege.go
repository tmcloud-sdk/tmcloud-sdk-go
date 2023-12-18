package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SqlserverUserWithPrivilege struct {
	Name string `json:"name"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (o SqlserverUserWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverUserWithPrivilege struct{}"
	}

	return strings.Join([]string{"SqlserverUserWithPrivilege", string(data)}, " ")
}
