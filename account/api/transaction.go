package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/pkg/utils"
)

func (a *ApiSetup) CreateTabung(c *fiber.Ctx) error {
	var reqPayload dao.CreateTabungTarikReq

    if err := c.BodyParser(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, err.Error(),
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : err.Error(),
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    a.Logger.Info(
        logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTabung API",
    )

    // Validate request payload
    if errMsg, err := utils.ValidateStruct(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": errMsg}, nil, errMsg,
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : errMsg,
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    data, remark, err := a.Services.CreateTabung(reqPayload)
    if err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "resp_msg" : remark,
            "resp_data" : data,
        })
    }

	response := map[string]interface{}{
        "resp_msg" : "Tabung Succeed",
        "resp_data" : data,
	}

	remark = "END: CreateTabung API"
    a.Logger.Info(
        logrus.Fields{"response": fmt.Sprintf("%+v", response)}, nil, remark,
    )
    return c.JSON(response)
}

func (a *ApiSetup) CreateTarik(c *fiber.Ctx) error {
	var reqPayload dao.CreateTabungTarikReq

    if err := c.BodyParser(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, err.Error(),
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : err.Error(),
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    a.Logger.Info(
        logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTarik API",
    )

    // Validate request payload
    if errMsg, err := utils.ValidateStruct(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": errMsg}, nil, errMsg,
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : errMsg,
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    data, remark, err := a.Services.CreateTarik(reqPayload)
    if err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "resp_msg" : remark,
            "resp_data" : data,
        })
    }

	response := map[string]interface{}{
        "resp_msg" : "Tarik Succeed",
        "resp_data" : data,
	}

	remark = "END: CreateTarik API"
    a.Logger.Info(
        logrus.Fields{"response": fmt.Sprintf("%+v", response)}, nil, remark,
    )
    return c.JSON(response)
}

func (a *ApiSetup) CreateTransfer(c *fiber.Ctx) error {
	var reqPayload dao.CreateTransferReq

    if err := c.BodyParser(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, err.Error(),
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : err.Error(),
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    a.Logger.Info(
        logrus.Fields{"req_payload": fmt.Sprintf("%+v", reqPayload)}, nil, "START: CreateTransfer API",
    )

    // Validate request payload
    if errMsg, err := utils.ValidateStruct(&reqPayload); err != nil {
        a.Logger.Error(
            logrus.Fields{"error": errMsg}, nil, errMsg,
        )
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "resp_msg" : errMsg,
            "resp_data" : dao.SaldoRes{
                Saldo: nil,
            },
        })
    }

    data, remark, err := a.Services.CreateTransfer(reqPayload)
    if err != nil {
        a.Logger.Error(
            logrus.Fields{"error": err.Error()}, nil, remark,
        )
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "resp_msg" : remark,
            "resp_data" : data,
        })
    }

	response := map[string]interface{}{
        "resp_msg" : "Transfer Succeed",
        "resp_data" : data,
	}

	remark = "END: CreateTransfer API"
    a.Logger.Info(
        logrus.Fields{"response": fmt.Sprintf("%+v", response)}, nil, remark,
    )
    return c.JSON(response)
}
