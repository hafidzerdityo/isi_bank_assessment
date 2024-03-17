package api

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/pkg/utils"
)

func (a *ApiSetup) TransactionMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            err := fmt.Errorf("auth header format error")
            remark := "Authorization header is not in the expected format"
            a.Logger.Error(
                logrus.Fields{"error": err.Error()}, nil, remark,
            )
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "resp_data": nil,
                "resp_msg":  remark,
            })
        }
        token := parts[1]

        // Check JWT Token
        isValid, remark, decodedJWT, err := utils.ValidateJWTToken(token)
        if err != nil || !isValid {
            a.Logger.Error(
                logrus.Fields{"error": err.Error()}, nil, remark,
            )
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "resp_data": nil,
                "resp_msg":  remark,
            })
        }

        // Store decodedJWT in context locals for later use
        c.Locals("decodedJWT", decodedJWT)

        // Call next middleware or handler
        return c.Next()
    }
}
