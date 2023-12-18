package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AuthenticatingProxy struct {
	Ca *string `json:"ca,omitempty"`

	Cert *string `json:"cert,omitempty"`

	PrivateKey *string `json:"privateKey,omitempty"`
}

func (o AuthenticatingProxy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthenticatingProxy struct{}"
	}

	return strings.Join([]string{"AuthenticatingProxy", string(data)}, " ")
}
