package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListConfigHistoriesResponse struct {
	HistoryNum *int32 `json:"history_num,omitempty"`

	Histories      *[]HistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListConfigHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigHistoriesResponse", string(data)}, " ")
}
