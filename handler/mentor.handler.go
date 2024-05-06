package handler

import (
	"errors"
	"strconv"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"

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
			Mentees:   mentor.Mentees,
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
		Mentees:   mentor.Mentees,
		Photo:     mentor.Photo,
	}

	return c.Status(200).JSON(responseDTO)
}

//post method

func MentorHandlerCreate(c *fiber.Ctx) error {
	var requestDTO dto.MentorCreateRequestDTO
	err := c.BodyParser(&requestDTO)

	if err != nil {
		return handleError(c, "Failed to parse request body", err)
	}

	mentor := entity.Mentor{
		FullName:  requestDTO.FullName,
		Email:     requestDTO.Email,
		Company:   requestDTO.Company,
		Specialty: requestDTO.Specialty,
		Bio:       requestDTO.Bio,
		Photo:     requestDTO.Photo,
	}

	results := database.DB.Create(&mentor)

	if results.Error != nil {
		return handleError(c, "Failed to create mentor", results.Error)
	}

	responseDTO := dto.MentorCreateResponseDTO{
		Message:  "Mentor created successfully",
		ID:       mentor.ID,
		FullName: mentor.FullName,
	}

	return c.Status(201).JSON(responseDTO)
}

func handleError(c *fiber.Ctx, message string, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"message": message,
		"error":   err.Error(),
	})
}
