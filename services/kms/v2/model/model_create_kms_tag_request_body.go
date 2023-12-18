package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKmsTagRequestBody struct {
	Tag *TagItem `json:"tag,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o CreateKmsTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKmsTagRequestBody", string(data)}, " ")
}
