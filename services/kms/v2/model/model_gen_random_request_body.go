package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GenRandomRequestBody struct {
	RandomDataLength string `json:"random_data_length"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o GenRandomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenRandomRequestBody struct{}"
	}

	return strings.Join([]string{"GenRandomRequestBody", string(data)}, " ")
}
