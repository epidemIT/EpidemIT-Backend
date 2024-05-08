package handler

import (
	"strconv"
	"strings"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/epidemIT/epidemIT-Backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ProjectHandlerGetAll(c *fiber.Ctx) error {
	// Parse query parameters
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	length, err := strconv.Atoi(c.Query("length", "10")) // default page length is 10
	if err != nil || length < 1 {
		length = 10
	}

	// Calculate offset
	offset := (page - 1) * length

	// Fetch projects with pagination
	var projects []entity.Project
	results := database.DB.Offset(offset).Limit(length).Find(&projects)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	// Prepare response DTO
	responseDTO := make([]dto.ProjectGetResponseDTO, len(projects))

	for i, project := range projects {
		responseDTO[i] = dto.ProjectGetResponseDTO{
			ID:                 project.ID,
			Name:               project.Name,
			ProjectDescription: project.ProjectDescription,
			Deadline:           project.Deadline,
			ImageURL:           project.ImageURL,
			PartnerName:        project.PartnerName,
			PartnerDescription: project.PartnerDescription,
			Users:              project.Users,
			Skills:             project.Skills,
			CreatedAt:          project.CreatedAt,
		}
	}

	return c.Status(200).JSON(responseDTO)
}

func ProjectHandlerGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var project entity.Project
	results := database.DB.Where("id = ?", id).First(&project)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := dto.ProjectGetResponseDTO{
		ID:                 project.ID,
		Name:               project.Name,
		ProjectDescription: project.ProjectDescription,
		Deadline:           project.Deadline,
		ImageURL:           project.ImageURL,
		PartnerName:        project.PartnerName,
		PartnerDescription: project.PartnerDescription,
		Users:              project.Users,
		Skills:             project.Skills,
		CreatedAt:          project.CreatedAt,
	}

	return c.Status(200).JSON(responseDTO)
}

func ProjectHandlerCreate(c *fiber.Ctx) error {
	var requestDTO dto.ProjectRegisterRequestDTO
	err := c.BodyParser(&requestDTO)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	project := entity.Project{
		Name:               requestDTO.Name,
		ProjectDescription: requestDTO.Description,
		Deadline:           requestDTO.Deadline,
		ImageURL:           requestDTO.ImageURL,
		PartnerName:        requestDTO.PartnerName,
		PartnerDescription: requestDTO.PartnerDesc,
	}

	results := database.DB.Create(&project)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to create project",
			"error":   results.Error,
		})
	}

	responseDTO := dto.ProjectRegisterResponseDTO{
		ID:          project.ID,
		Message:     "Project created successfully",
		Name:        project.Name,
		Description: project.ProjectDescription,
		Deadline:    project.Deadline,
		ImageURL:    project.ImageURL,
	}

	return c.Status(200).JSON(responseDTO)
}

func ProjectApplyRegister(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error verifying token",
			"error":   err.Error(),
		})
	}

	body := new(dto.ProjectApplyRequestDTO)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing new ProjectApply",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(body)

	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error validating new ProjectApply",
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

	var project entity.Project
	err = database.DB.Where("id = ?", body.ProjectID).First(&project).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error getting project",
			"error":   err.Error(),
		})
	}

	project.Users = append(project.Users, user)

	projectApply := entity.ProjectApply{
		ProjectID: project.ID,
		UserID:    user.ID,
		Progress:  body.Progress,
	}

	results := database.DB.Save(&project)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error applying to project",
			"error":   results.Error,
		})
	}

	results = database.DB.Create(&projectApply)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error applying to project",
			"error":   results.Error,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":    "Applied to project successfully",
		"id":         projectApply.ID,
		"project_id": projectApply.ProjectID,
		"user_id":    projectApply.UserID,
		"progress":   projectApply.Progress,
	})
}

func GetAllProjectAppliedByUserID(c *fiber.Ctx) error {
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

	var projectApplies []entity.ProjectApply
	results := database.DB.Where("user_id = ?", user.ID).Find(&projectApplies)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := make([]dto.ProjectApplyGetByUserIDResponseDTO, len(projectApplies))

	for i, projectApply := range projectApplies {
		var project entity.Project
		results := database.DB.Where("id = ?", projectApply.ProjectID).First(&project)

		if results.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": results.Error,
			})
		}

		responseDTO[i] = dto.ProjectApplyGetByUserIDResponseDTO{
			ID: projectApply.ID,
			UserID: user.ID,
			Project: dto.ProjectGetResponseDTO{
				ID:                 project.ID,
				Name:               project.Name,
				ProjectDescription: project.ProjectDescription,
				Deadline:           project.Deadline,
				ImageURL:           project.ImageURL,
				PartnerName:        project.PartnerName,
				PartnerDescription: project.PartnerDescription,
				Users:              project.Users,
				Skills:             project.Skills,
				CreatedAt:          project.CreatedAt,
			},
			Progress: projectApply.Progress,
		}
	}

	return c.Status(200).JSON(responseDTO)
}
