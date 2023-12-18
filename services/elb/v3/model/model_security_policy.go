package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SecurityPolicy struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Listeners []ListenerRef `json:"listeners"`

	Protocols []string `json:"protocols"`

	Ciphers []string `json:"ciphers"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`
}

func (o SecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityPolicy struct{}"
	}

	return strings.Join([]string{"SecurityPolicy", string(data)}, " ")
}
