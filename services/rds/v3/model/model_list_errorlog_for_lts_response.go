package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListErrorlogForLtsResponse struct {
	ErrorLogList   *[]ErrorLogItem `json:"error_log_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListErrorlogForLtsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorlogForLtsResponse struct{}"
	}

	return strings.Join([]string{"ListErrorlogForLtsResponse", string(data)}, " ")
}
