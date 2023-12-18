package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RedisConfig struct {
	ParamValue string `json:"param_value"`

	ParamName string `json:"param_name"`

	ParamId string `json:"param_id"`
}

func (o RedisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConfig struct{}"
	}

	return strings.Join([]string{"RedisConfig", string(data)}, " ")
}
