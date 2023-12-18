package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateServerGroupResult struct {
	Id string `json:"id"`

	Members []string `json:"members"`

	Metadata map[string]string `json:"metadata"`

	Name string `json:"name"`

	Policies []string `json:"policies"`
}

func (o CreateServerGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupResult struct{}"
	}

	return strings.Join([]string{"CreateServerGroupResult", string(data)}, " ")
}
