package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyRedisConfigBody struct {
	RedisConfig *[]RedisConfig `json:"redis_config,omitempty"`
}

func (o ModifyRedisConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRedisConfigBody struct{}"
	}

	return strings.Join([]string{"ModifyRedisConfigBody", string(data)}, " ")
}
