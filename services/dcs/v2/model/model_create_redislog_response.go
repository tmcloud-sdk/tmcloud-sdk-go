package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRedislogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRedislogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogResponse struct{}"
	}

	return strings.Join([]string{"CreateRedislogResponse", string(data)}, " ")
}
