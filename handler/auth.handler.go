package handler

import (
	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/epidemIT/epidemIT-Backend/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserLogin(ctx *fiber.Ctx) error {
	loginRequest := new(dto.LoginRequestDTO)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing login request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating login request",
			"error":   errValidate.Error(),
		})
	}

	var existingUser entity.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&existingUser).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	isValid := utils.CheckPassword(existingUser.Password, loginRequest.Password)
	if !isValid {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	claims := jwt.MapClaims{}
	claims["email"] = existingUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := utils.GenerateToken(&claims)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error generating token",
			"error":   err.Error(),
		})
	}

	responseDTO := dto.LoginResponseDTO{
		Message: "Login successful",
		Token:   token,
	}

	return ctx.Status(200).JSON(responseDTO)
}

func GetCurrentUser(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	var user entity.User
	err = database.DB.Where("email = ?", claims["email"]).First(&user).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error getting user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(200).JSON(user)
}