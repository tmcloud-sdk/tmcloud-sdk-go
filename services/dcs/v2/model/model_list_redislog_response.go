package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRedislogResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	FileList       *[]RunlogItem `json:"file_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRedislogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedislogResponse struct{}"
	}

	return strings.Join([]string{"ListRedislogResponse", string(data)}, " ")
}
