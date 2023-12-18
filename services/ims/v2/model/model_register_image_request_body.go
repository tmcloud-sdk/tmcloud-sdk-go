package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RegisterImageRequestBody struct {
	ImageUrl string `json:"image_url"`
}

func (o RegisterImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterImageRequestBody", string(data)}, " ")
}
