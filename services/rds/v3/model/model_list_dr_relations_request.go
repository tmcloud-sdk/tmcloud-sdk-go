package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListDrRelationsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListDrRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrRelationsRequest struct{}"
	}

	return strings.Join([]string{"ListDrRelationsRequest", string(data)}, " ")
}
