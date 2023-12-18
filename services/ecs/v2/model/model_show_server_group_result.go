package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowServerGroupResult struct {
	Id string `json:"id"`

	Members []string `json:"members"`

	Metadata map[string]string `json:"metadata"`

	Name string `json:"name"`

	Policies []string `json:"policies"`
}

func (o ShowServerGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupResult struct{}"
	}

	return strings.Join([]string{"ShowServerGroupResult", string(data)}, " ")
}
