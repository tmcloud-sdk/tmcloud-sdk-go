package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceParamHistoriesResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Histories      *[]ParamGroupHistoryResult `json:"histories,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListInstanceParamHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceParamHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceParamHistoriesResponse", string(data)}, " ")
}
