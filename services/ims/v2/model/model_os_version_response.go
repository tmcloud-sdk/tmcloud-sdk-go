package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type OsVersionResponse struct {
	Status string `json:"status"`

	Id *string `json:"id,omitempty"`

	Links *[]Links `json:"links,omitempty"`
}

func (o OsVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersionResponse struct{}"
	}

	return strings.Join([]string{"OsVersionResponse", string(data)}, " ")
}
