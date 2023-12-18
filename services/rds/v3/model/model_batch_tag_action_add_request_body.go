package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchTagActionAddRequestBody struct {
	Action string `json:"action"`

	Tags []TagWithKeyValue `json:"tags"`
}

func (o BatchTagActionAddRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionAddRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionAddRequestBody", string(data)}, " ")
}
