package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CommandTimeTaken struct {
	CallsSum int32 `json:"calls_sum"`

	UsecSum float64 `json:"usec_sum"`

	CommandName string `json:"command_name"`

	PerUsec string `json:"per_usec"`

	AverageUsec float64 `json:"average_usec"`
}

func (o CommandTimeTaken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandTimeTaken struct{}"
	}

	return strings.Join([]string{"CommandTimeTaken", string(data)}, " ")
}
