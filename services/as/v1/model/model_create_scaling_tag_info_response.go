package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScalingTagInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingTagInfoResponse", string(data)}, " ")
}
