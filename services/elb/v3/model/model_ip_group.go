package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type IpGroup struct {
	CreatedAt string `json:"created_at"`

	Description string `json:"description"`

	Id string `json:"id"`

	IpList []IpInfo `json:"ip_list"`

	Listeners []ListenerRef `json:"listeners"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	UpdatedAt string `json:"updated_at"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
