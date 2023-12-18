package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSlowlogResponse struct {
	Count *int32 `json:"count,omitempty"`

	Slowlogs       *[]SlowlogItem `json:"slowlogs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSlowlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogResponse", string(data)}, " ")
}
