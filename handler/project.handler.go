package handler

import (
	"strconv"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
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
			PartnerName:        project.PartnerName,
			PartnerDescription: project.PartnerDescription,
			Users:              project.Users,
			Skills:             project.Skills,
			FirstMaterial:      project.FirstMaterial,
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
		PartnerName:        project.PartnerName,
		PartnerDescription: project.PartnerDescription,
		Users:              project.Users,
		Skills:             project.Skills,
		FirstMaterial:      project.FirstMaterial,
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
		})
	}

	project := entity.Project{
		Name:               requestDTO.Name,
		ProjectDescription: requestDTO.Description,
		Deadline:           requestDTO.Deadline,
		PartnerName:        requestDTO.PartnerName,
		PartnerDescription: requestDTO.PartnerDesc,
		FirstMaterial:      requestDTO.FirstMaterial,
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
	}

	return c.Status(200).JSON(responseDTO)
}
