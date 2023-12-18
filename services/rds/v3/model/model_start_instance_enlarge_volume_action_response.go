package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StartInstanceEnlargeVolumeActionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceEnlargeVolumeActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceEnlargeVolumeActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceEnlargeVolumeActionResponse", string(data)}, " ")
}
