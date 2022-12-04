package middleware

import (
	"encoding/json"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ===============================================================//
// =========== ใช้สำหรับ Track Logs => request and respose =========//
// ==============================================================//
type transactionLogs struct {
	TransactionId string                 `json:"TransactionId,omitempty"`
	UserId        string                 `json:"UserId,omitempty"`
	Level         string                 `json:"Level"`
	Method        string                 `json:"Method,omitempty"`
	HttpStatus    int                    `json:"HttpStatus,omitempty"`
	RequestURI    string                 `json:"RequestURI,omitempty"`
	Timestamp     time.Time              `json:"@timestamp"`
	Code          string                 `json:"Code,omitempty"`
	Request       map[string]interface{} `json:"Request"`
	Response      interface{}            `json:"Response"`
}

func Logger(c *fiber.Ctx) error {
	var tLog transactionLogs
	var request map[string]interface{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	//set log
	tLog.Method = c.Method()
	tLog.RequestURI = c.OriginalURL()
	tLog.Request = request

	if err := c.Next(); err != nil {
		return err
	}

	var response interface{}
	responseBody := c.Response().Body()
	if err := json.Unmarshal([]byte(string(string(responseBody))), &response); err != nil {
		return err
	}
	statusCode := c.Response().Header.StatusCode()

	//set log
	tLog.Response = response
	tLog.HttpStatus = statusCode
	tLog.Timestamp = utils.GetTimeStampLog()
	userInfo := utils.GetUserInfo(c)
	tLog.TransactionId = userInfo.TransactionId
	tLog.UserId = userInfo.UserId

	switch statusCode {
	case fiber.StatusOK:
		tLog.Level = enum.LogLv_Info
	case fiber.StatusBadRequest,
		fiber.StatusUnauthorized:
		tLog.Level = enum.LogLv_Warn
	default:
		tLog.Level = enum.LogLv_Err
	}

	b, err := json.Marshal(tLog)
	fmt.Println(string(b))
	return err
}
