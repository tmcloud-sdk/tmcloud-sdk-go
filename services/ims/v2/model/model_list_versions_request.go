package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListVersionsRequest struct {
}

func (o ListVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListVersionsRequest", string(data)}, " ")
}
