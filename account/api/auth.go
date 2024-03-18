package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/pkg/utils"
)


func (a *ApiSetup) AccountLogin(c *fiber.Ctx) error {
	var reqPayload dao.AccountLoginReq

    if err := c.BodyParser(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, err.Error(),
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : err.Error(),
            "resp_data" : dao.CreateCustRes{
                NoRekening: nil,
            },
        })
    }

	reqPayloadForLog := reqPayload
	reqPayloadForLog.Pin = "*REDACTED*"
	reqPayloadForLog.Password = "*REDACTED*"
    a.Logger.Info(
        logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayloadForLog)}, nil, "START: AccountLogin API",
    )

    // Validate request payload
    if errMsg, err := utils.ValidateStruct(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": errMsg}, nil, errMsg,
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : errMsg,
            "resp_data" : dao.CreateCustRes{
                NoRekening: nil,
            },
        })
    }

    data, remark, err := a.Services.AccountLogin(reqPayload)
    if err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : remark,
            "resp_data" : data,
        })
    }

	response := map[string]interface{}{
        "resp_msg" : "Login Succeed",
        "resp_data" : data,
	}

	remark = "END: AccountLogin API"
    a.Logger.Info(
        logrus.Fields{"response": fmt.Sprintf("%+v", response)}, nil, remark,
    )
    return c.JSON(response)
}
