package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateGrantResponse struct {
	GrantId        *string `json:"grant_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantResponse struct{}"
	}

	return strings.Join([]string{"CreateGrantResponse", string(data)}, " ")
}
