package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerMetadataRequestBody struct {
	Metadata map[string]string `json:"metadata"`
}

func (o UpdateServerMetadataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataRequestBody", string(data)}, " ")
}
