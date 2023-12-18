package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SqlserverUserForCreation struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

func (o SqlserverUserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverUserForCreation struct{}"
	}

	return strings.Join([]string{"SqlserverUserForCreation", string(data)}, " ")
}
