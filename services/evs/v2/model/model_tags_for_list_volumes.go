package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type TagsForListVolumes struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagsForListVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsForListVolumes struct{}"
	}

	return strings.Join([]string{"TagsForListVolumes", string(data)}, " ")
}
