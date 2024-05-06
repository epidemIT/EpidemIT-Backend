package handler

import (
	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *fiber.Ctx) error {
	user := new(dto.UserRegisterRequestDTO)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new user",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new user",
			"error":   errValidate.Error(),
		})
	}

	birthDate, err := time.Parse("2006-01-02", user.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing BirthDate",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error hashing password",
			"error":   err.Error(),
		})
	}

	newUser := entity.User{
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  string(hashedPassword),
		BirthDate: birthDate,
	}
	var existingUser entity.User
	res := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newUserRes := database.DB.Create(&newUser)

	if newUserRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new user",
			"error":   newUserRes.Error.Error(),
		})
	}

	responseDTO := dto.UserRegisterResponseDTO{
		Message:   "New user created successfully",
		ID:        newUser.ID,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		BirthDate: newUser.BirthDate.Format("2006-01-02"),
	}

	return c.Status(201).JSON(responseDTO)
}