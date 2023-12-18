package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowKmsTagsRequest struct {
	KeyId string `json:"key_id"`
}

func (o ShowKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKmsTagsRequest", string(data)}, " ")
}
