package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UserForCreation struct {
	Name string `json:"name"`

	Password string `json:"password"`

	Comment *string `json:"comment,omitempty"`

	Hosts *[]string `json:"hosts,omitempty"`

	Databases *[]DatabaseWithPrivilegeObject `json:"databases,omitempty"`
}

func (o UserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserForCreation struct{}"
	}

	return strings.Join([]string{"UserForCreation", string(data)}, " ")
}
