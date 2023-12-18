package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdatePasswordResponse struct {
	LockTime *string `json:"lock_time,omitempty"`

	Result *UpdatePasswordResponseResult `json:"result,omitempty"`

	LockTimeLeft *string `json:"lock_time_left,omitempty"`

	RetryTimesLeft *string `json:"retry_times_left,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdatePasswordResponse", string(data)}, " ")
}

type UpdatePasswordResponseResult struct {
	value string
}

type UpdatePasswordResponseResultEnum struct {
	SUCCESS         UpdatePasswordResponseResult
	PASSWORD_FAILED UpdatePasswordResponseResult
	LOCKED          UpdatePasswordResponseResult
	FAILED          UpdatePasswordResponseResult
}

func GetUpdatePasswordResponseResultEnum() UpdatePasswordResponseResultEnum {
	return UpdatePasswordResponseResultEnum{
		SUCCESS: UpdatePasswordResponseResult{
			value: "success",
		},
		PASSWORD_FAILED: UpdatePasswordResponseResult{
			value: "passwordFailed",
		},
		LOCKED: UpdatePasswordResponseResult{
			value: "locked",
		},
		FAILED: UpdatePasswordResponseResult{
			value: "failed",
		},
	}
}

func (c UpdatePasswordResponseResult) Value() string {
	return c.value
}

func (c UpdatePasswordResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePasswordResponseResult) UnmarshalJSON(b []byte) error {
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
