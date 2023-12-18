package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKmsTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateKmsTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagResponse struct{}"
	}

	return strings.Join([]string{"CreateKmsTagResponse", string(data)}, " ")
}
