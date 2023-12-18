package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreatePortResponse struct {
	Port           *Port `json:"port,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortResponse struct{}"
	}

	return strings.Join([]string{"CreatePortResponse", string(data)}, " ")
}
