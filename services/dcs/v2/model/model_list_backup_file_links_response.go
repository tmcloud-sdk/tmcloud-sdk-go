package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListBackupFileLinksResponse struct {
	FilePath *string `json:"file_path,omitempty"`

	BucketName *string `json:"bucket_name,omitempty"`

	Links          *[]LinksItem `json:"links,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListBackupFileLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupFileLinksResponse struct{}"
	}

	return strings.Join([]string{"ListBackupFileLinksResponse", string(data)}, " ")
}
