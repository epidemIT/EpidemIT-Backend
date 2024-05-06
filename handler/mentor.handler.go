package handler

import (
	"errors"
	"strconv"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MentorHandlerGetAll(c *fiber.Ctx) error {
	// Parse query parameters
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	length, err := strconv.Atoi(c.Query("length", "10")) // default page length is 10
	if err != nil || length < 1 {
		length = 10
	}

	// Calculate offset and limit
	offset := (page - 1) * length

	var mentors []entity.Mentor
	results := database.DB.Offset(offset).Limit(length).Find(&mentors)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.MentorGetResponseDTO, len(mentors))

	for i, mentor := range mentors {
		responseDTO[i] = dto.MentorGetResponseDTO{
			ID:        mentor.ID,
			FullName:  mentor.FullName,
			Email:     mentor.Email,
			Company:   mentor.Company,
			Specialty: mentor.Specialty,
			Bio:       mentor.Bio,
			Photo:     mentor.Photo,
		}
	}

	return c.Status(200).JSON(responseDTO)
}

func MentorHandlerGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var mentor entity.Mentor
	results := database.DB.Where("id = ?", id).First(&mentor)

	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Mentor not found",
			})
		}
		return handleError(c, "Failed to get mentor", results.Error)
	}

	responseDTO := dto.MentorGetResponseDTO{
		ID:        mentor.ID,
		FullName:  mentor.FullName,
		Email:     mentor.Email,
		Company:   mentor.Company,
		Specialty: mentor.Specialty,
		Bio:       mentor.Bio,
		Photo:     mentor.Photo,
	}

	return c.Status(200).JSON(responseDTO)
}

//post method

func MentorHandlerCreate(c *fiber.Ctx) error {
	mentor := new(dto.MentorCreateRequestDTO)

	if err := c.BodyParser(mentor); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new mentor",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(mentor)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new mentor",
			"error":   errValidate.Error(),
		})
	}

	newMentor := entity.Mentor{
		FullName:  mentor.FullName,
		Email:     mentor.Email,
		Company:   mentor.Company,
		Specialty: mentor.Specialty,
		Bio:       mentor.Bio,
		Photo:     mentor.Photo,
	}

	var existingMentor entity.Mentor
	res := database.DB.Where("email = ?", mentor.Email).First(&existingMentor)
	if res.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	newMentorRes := database.DB.Create(&newMentor)

	if newMentorRes.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error creating new mentor",
			"error":   newMentorRes.Error.Error(),
		})
	}

	responseDTO := dto.MentorCreateResponseDTO{
		Message:  "Mentor created successfully",
		ID:       newMentor.ID,
		FullName: newMentor.FullName,
	}

	return c.Status(201).JSON(responseDTO)
}

func handleError(c *fiber.Ctx, message string, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"message": message,
		"error":   err.Error(),
	})
}
