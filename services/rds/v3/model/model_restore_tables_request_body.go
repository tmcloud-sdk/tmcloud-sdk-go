package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RestoreTablesRequestBody struct {
	RestoreTime int64 `json:"restoreTime"`

	RestoreTables []RestoreDatabasesInfo `json:"restoreTables"`
}

func (o RestoreTablesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreTablesRequestBody", string(data)}, " ")
}
