package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	secret string
}

func NewAuthHandler(secret string) *AuthHandler {
	return &AuthHandler{secret: secret}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": req.Username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	t, _ := token.SignedString([]byte(h.secret))

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return c.JSON(fiber.Map{"message": "logged in"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{Name: "jwt", Value: "", Expires: time.Now().Add(-time.Hour)})
	return c.JSON(fiber.Map{"message": "logged out"})
}
