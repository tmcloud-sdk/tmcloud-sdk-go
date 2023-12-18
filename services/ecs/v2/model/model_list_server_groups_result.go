package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListServerGroupsResult struct {
	Id string `json:"id"`

	Members []string `json:"members"`

	Metadata map[string]string `json:"metadata"`

	Name string `json:"name"`

	Policies []string `json:"policies"`
}

func (o ListServerGroupsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupsResult struct{}"
	}

	return strings.Join([]string{"ListServerGroupsResult", string(data)}, " ")
}
