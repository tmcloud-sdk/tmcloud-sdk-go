package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryProgressResp struct {
	JobId *string `json:"job_id,omitempty"`

	Progress *string `json:"progress,omitempty"`

	IncreTransDelay *string `json:"incre_trans_delay,omitempty"`

	TaskMode *QueryProgressRespTaskMode `json:"task_mode,omitempty"`

	TransferStatus *string `json:"transfer_status,omitempty"`

	ProcessTime *string `json:"process_time,omitempty"`

	RemainingTime *string `json:"remaining_time,omitempty"`

	ProgressMap map[string]ProgressInfo `json:"progress_map,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o QueryProgressResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryProgressResp struct{}"
	}

	return strings.Join([]string{"QueryProgressResp", string(data)}, " ")
}

type QueryProgressRespTaskMode struct {
	value string
}

type QueryProgressRespTaskModeEnum struct {
	FULL_TRANS      QueryProgressRespTaskMode
	FULL_INCR_TRANS QueryProgressRespTaskMode
	INCR_TRANS      QueryProgressRespTaskMode
}

func GetQueryProgressRespTaskModeEnum() QueryProgressRespTaskModeEnum {
	return QueryProgressRespTaskModeEnum{
		FULL_TRANS: QueryProgressRespTaskMode{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: QueryProgressRespTaskMode{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: QueryProgressRespTaskMode{
			value: "INCR_TRANS",
		},
	}
}

func (c QueryProgressRespTaskMode) Value() string {
	return c.value
}

func (c QueryProgressRespTaskMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryProgressRespTaskMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
