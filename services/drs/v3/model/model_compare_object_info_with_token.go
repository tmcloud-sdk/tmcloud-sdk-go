package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CompareObjectInfoWithToken struct {
	DbName string `json:"db_name"`

	TableNameWithToken *[]CompareTableInfoWithToken `json:"table_name_with_token,omitempty"`
}

func (o CompareObjectInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareObjectInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareObjectInfoWithToken", string(data)}, " ")
}
