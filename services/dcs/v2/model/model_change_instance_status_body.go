package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeInstanceStatusBody struct {
	Instances *[]string `json:"instances,omitempty"`

	Action *string `json:"action,omitempty"`
}

func (o ChangeInstanceStatusBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceStatusBody struct{}"
	}

	return strings.Join([]string{"ChangeInstanceStatusBody", string(data)}, " ")
}
