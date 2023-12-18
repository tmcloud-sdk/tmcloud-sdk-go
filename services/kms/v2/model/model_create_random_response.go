package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRandomResponse struct {
	RandomData     *string `json:"random_data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRandomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRandomResponse struct{}"
	}

	return strings.Join([]string{"CreateRandomResponse", string(data)}, " ")
}
