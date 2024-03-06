package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)

func (a *ApiSetup) TransactionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		no_rekening := parts[0]
		pin := parts[1]

		accountAndPinParam := dao.CheckAccountAndPinReq{
			NoRekening: no_rekening,
			Pin: pin,
		}

		// Check if NoRekening and PIN are correct
		isValid, remark, err := a.Services.CheckAccountAndPin(accountAndPinParam)
		if err != nil || !isValid {
			a.Logger.Error(
				logrus.Fields{"error": err.Error()}, nil, remark,
			)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"resp_data": nil,
				"resp_msg":  remark,
			})
		}


		// Call next middleware or handler
		return c.Next()
	}
}