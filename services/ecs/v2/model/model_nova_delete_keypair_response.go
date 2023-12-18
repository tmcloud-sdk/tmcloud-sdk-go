package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaDeleteKeypairResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDeleteKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteKeypairResponse struct{}"
	}

	return strings.Join([]string{"NovaDeleteKeypairResponse", string(data)}, " ")
}
