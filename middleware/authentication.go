package middleware

import (
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/google/uuid"
)

func Authorization() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod:  "HS256",
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(config.Server().Secret_Key),
	})
}

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized)
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	token, err := verifyToken(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var userInfo utils.UserInfo
		utils.JsonDeserialize(fmt.Sprint(claims[enum.USER_INFO]), &userInfo)
		userInfo.TransactionId = uuid.New().String()
		c.Locals(enum.USER_INFO, userInfo)

		if err := c.Next(); err != nil {
			panic(err)
		}
		return nil
	}

	return fiber.ErrUnauthorized
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Server().Secret_Key), nil
}
