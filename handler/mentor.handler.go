package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/epidemIT/epidemIT-Backend/utils"
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
			Reviews:   mentor.Reviews,
			Sessions:  mentor.Sessions,
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
		Reviews:   mentor.Reviews,
		Sessions:  mentor.Sessions,
	}

	return c.Status(200).JSON(responseDTO)
}

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

func MentorApplyRegister(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	body := new(dto.MentorApplyRequestDTO)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(body)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating request",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err = database.DB.Where("email = ?", claims["email"]).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting user",
			"error":   err.Error(),
		})
	}

	var mentor entity.Mentor
	err = database.DB.Where("id = ?", body.MentorID).First(&mentor).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting project",
			"error":   err.Error(),
		})
	}

	mentor.Mentees = append(mentor.Mentees, user)
	mentorApply := entity.MentorApply{
		MentorID: mentor.ID,
		UserID:   user.ID,
		Date:     body.Date,
	}

	results := database.DB.Save(&mentor)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error applying to mentor",
			"error":   results.Error,
		})
	}

	results = database.DB.Create(&mentorApply)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error applying to project",
			"error":   results.Error,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Applied to mentor successfully",
		"id":      mentorApply.ID,
		"mentor":  mentorApply.MentorID,
		"user_id": mentorApply.UserID,
		"date":    mentorApply.Date,
	})
}

func GetAllMentorAppliedByUserID(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	var user entity.User
	err = database.DB.Where("email = ?", claims["email"]).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting user",
			"error":   err.Error(),
		})
	}

	var mentorApplies []entity.MentorApply
	results := database.DB.Where("user_id = ?", user.ID).Find(&mentorApplies)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.MentorApplyGetByUserIDResponseDTO, len(mentorApplies))

	for i, mentorApply := range mentorApplies {
		var mentor entity.Mentor
		results := database.DB.Where("id = ?", mentorApply.MentorID).First(&mentor)

		if results.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": results.Error,
			})
		}

		responseDTO[i] = dto.MentorApplyGetByUserIDResponseDTO{
			ID:     mentorApply.ID,
			UserID: mentorApply.UserID,
			Mentor: dto.MentorGetResponseDTO{
				ID:       mentor.ID,
				FullName: mentor.FullName,
				Photo:    mentor.Photo,
			},
		}
	}

	return c.Status(200).JSON(responseDTO)
}
