package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowHotkeyTaskDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *ShowHotkeyTaskDetailsResponseStatus `json:"status,omitempty"`

	ScanType *ShowHotkeyTaskDetailsResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt *string `json:"finished_at,omitempty"`

	Num *int32 `json:"num,omitempty"`

	Keys           *[]HotkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHotkeyTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHotkeyTaskDetailsResponse", string(data)}, " ")
}

type ShowHotkeyTaskDetailsResponseStatus struct {
	value string
}

type ShowHotkeyTaskDetailsResponseStatusEnum struct {
	WAITING ShowHotkeyTaskDetailsResponseStatus
	RUNNING ShowHotkeyTaskDetailsResponseStatus
	SUCCESS ShowHotkeyTaskDetailsResponseStatus
	FAILED  ShowHotkeyTaskDetailsResponseStatus
}

func GetShowHotkeyTaskDetailsResponseStatusEnum() ShowHotkeyTaskDetailsResponseStatusEnum {
	return ShowHotkeyTaskDetailsResponseStatusEnum{
		WAITING: ShowHotkeyTaskDetailsResponseStatus{
			value: "waiting",
		},
		RUNNING: ShowHotkeyTaskDetailsResponseStatus{
			value: "running",
		},
		SUCCESS: ShowHotkeyTaskDetailsResponseStatus{
			value: "success",
		},
		FAILED: ShowHotkeyTaskDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowHotkeyTaskDetailsResponseStatus) Value() string {
	return c.value
}

func (c ShowHotkeyTaskDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHotkeyTaskDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowHotkeyTaskDetailsResponseScanType struct {
	value string
}

type ShowHotkeyTaskDetailsResponseScanTypeEnum struct {
	MANUAL ShowHotkeyTaskDetailsResponseScanType
	AUTO   ShowHotkeyTaskDetailsResponseScanType
}

func GetShowHotkeyTaskDetailsResponseScanTypeEnum() ShowHotkeyTaskDetailsResponseScanTypeEnum {
	return ShowHotkeyTaskDetailsResponseScanTypeEnum{
		MANUAL: ShowHotkeyTaskDetailsResponseScanType{
			value: "manual",
		},
		AUTO: ShowHotkeyTaskDetailsResponseScanType{
			value: "auto",
		},
	}
}

func (c ShowHotkeyTaskDetailsResponseScanType) Value() string {
	return c.value
}

func (c ShowHotkeyTaskDetailsResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHotkeyTaskDetailsResponseScanType) UnmarshalJSON(b []byte) error {
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
