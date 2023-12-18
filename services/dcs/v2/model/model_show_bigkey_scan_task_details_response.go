package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowBigkeyScanTaskDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Status *ShowBigkeyScanTaskDetailsResponseStatus `json:"status,omitempty"`

	ScanType *ShowBigkeyScanTaskDetailsResponseScanType `json:"scan_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	StartedAt *string `json:"started_at,omitempty"`

	FinishedAt *string `json:"finished_at,omitempty"`

	Num *int32 `json:"num,omitempty"`

	Keys           *[]BigkeysBody `json:"keys,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowBigkeyScanTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyScanTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowBigkeyScanTaskDetailsResponse", string(data)}, " ")
}

type ShowBigkeyScanTaskDetailsResponseStatus struct {
	value string
}

type ShowBigkeyScanTaskDetailsResponseStatusEnum struct {
	WAITING ShowBigkeyScanTaskDetailsResponseStatus
	RUNNING ShowBigkeyScanTaskDetailsResponseStatus
	SUCCESS ShowBigkeyScanTaskDetailsResponseStatus
	FAILED  ShowBigkeyScanTaskDetailsResponseStatus
}

func GetShowBigkeyScanTaskDetailsResponseStatusEnum() ShowBigkeyScanTaskDetailsResponseStatusEnum {
	return ShowBigkeyScanTaskDetailsResponseStatusEnum{
		WAITING: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "waiting",
		},
		RUNNING: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "running",
		},
		SUCCESS: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "success",
		},
		FAILED: ShowBigkeyScanTaskDetailsResponseStatus{
			value: "failed",
		},
	}
}

func (c ShowBigkeyScanTaskDetailsResponseStatus) Value() string {
	return c.value
}

func (c ShowBigkeyScanTaskDetailsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBigkeyScanTaskDetailsResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowBigkeyScanTaskDetailsResponseScanType struct {
	value string
}

type ShowBigkeyScanTaskDetailsResponseScanTypeEnum struct {
	MANUAL ShowBigkeyScanTaskDetailsResponseScanType
	AUTO   ShowBigkeyScanTaskDetailsResponseScanType
}

func GetShowBigkeyScanTaskDetailsResponseScanTypeEnum() ShowBigkeyScanTaskDetailsResponseScanTypeEnum {
	return ShowBigkeyScanTaskDetailsResponseScanTypeEnum{
		MANUAL: ShowBigkeyScanTaskDetailsResponseScanType{
			value: "manual",
		},
		AUTO: ShowBigkeyScanTaskDetailsResponseScanType{
			value: "auto",
		},
	}
}

func (c ShowBigkeyScanTaskDetailsResponseScanType) Value() string {
	return c.value
}

func (c ShowBigkeyScanTaskDetailsResponseScanType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBigkeyScanTaskDetailsResponseScanType) UnmarshalJSON(b []byte) error {
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
