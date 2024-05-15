package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/epidemIT/epidemIT-Backend/database"
	"github.com/epidemIT/epidemIT-Backend/model/dto"
	"github.com/epidemIT/epidemIT-Backend/model/entity"
	"github.com/epidemIT/epidemIT-Backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const (
	emailTemplate = `
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f7f7f7;
						margin: 0;
						padding: 20px;
					}
					.email-container {
						background-color: #ffffff;
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						border-radius: 5px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					.details-header {
						font-weight: bold;
					}
				</style>
					</head>
					<body>
						<div class="email-container">
						<p>Dear Applicants, </p>
						<p>We are delighted to inform you that your registration for the Financial Aid has been successfully received and confirmed.</p>
						<p>Our team will review your application and get back to you as soon as possible. Please stay tuned for further updates.</p>
						<p>Please keep this email for your records. In case you have any questions or need to make any changes to your registration details, please do not hesitate to contact us at <a href="mailto:bistleague@std.stei.itb.ac.id">gibran@epidemit.id</a> or +62 81290908333.</p>

						<p>Once again, congratulations on your successful registration. We look forward to seeing you at the project. Best of luck with your preparations!</p>

						<br />
						<p>Sincerely,</p>
						<p>EpidemIT Operations Team</p>
						</div>
					</body>
					</html>
			`
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

	// get skills from project_skills table
	for i, project := range projects {
		var skills []entity.Skill
		database.DB.Model(&project).Association("Skills").Find(&skills)

		projects[i].Skills = skills
	}

	// Prepare response DTO
	responseDTO := make([]dto.ProjectGetResponseDTO, len(projects))

	for i, project := range projects {
		responseDTO[i] = dto.ProjectGetResponseDTO{
			ID:                 project.ID,
			Name:               project.Name,
			ProjectDescription: project.ProjectDescription,
			ShortDescription:   project.ShortDescription,
			MetodeBelajar:      project.MetodeBelajar,
			PeralatanBelajar:   project.PeralatanBelajar,
			Silabus:            project.Silabus,
			TotalHours:         project.TotalHours,
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

	// get skills from project_skills table
	var skills []entity.Skill
	database.DB.Model(&project).Association("Skills").Find(&skills)

	project.Skills = skills

	// get users from user_project table
	var users []entity.User
	database.DB.Model(&project).Association("Users").Find(&users)

	project.Users = users

	responseDTO := dto.ProjectGetResponseDTO{
		ID:                 project.ID,
		Name:               project.Name,
		ProjectDescription: project.ProjectDescription,
		ShortDescription:   project.ShortDescription,
		MetodeBelajar:      project.MetodeBelajar,
		PeralatanBelajar:   project.PeralatanBelajar,
		Silabus:            project.Silabus,
		TotalHours:         project.TotalHours,
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
		ShortDescription:   requestDTO.ShortDescription,
		MetodeBelajar:      requestDTO.MetodeBelajar,
		PeralatanBelajar:   requestDTO.PeralatanBelajar,
		Silabus:            requestDTO.Silabus,
		TotalHours:         requestDTO.TotalHours,
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
		ID:               project.ID,
		Message:          "Project created successfully",
		Name:             project.Name,
		Description:      project.ProjectDescription,
		ShortDescription: project.ShortDescription,
		MetodeBelajar:    project.MetodeBelajar,
		PeralatanBelajar: project.PeralatanBelajar,
		Silabus:          project.Silabus,
		TotalHours:       project.TotalHours,
		Deadline:         project.Deadline,
		ImageURL:         project.ImageURL,
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
			ID:     projectApply.ID,
			UserID: user.ID,
			Project: dto.ProjectGetResponseDTO{
				ID:                 project.ID,
				Name:               project.Name,
				ProjectDescription: project.ProjectDescription,
				ShortDescription:   project.ShortDescription,
				MetodeBelajar:      project.MetodeBelajar,
				PeralatanBelajar:   project.PeralatanBelajar,
				Silabus:            project.Silabus,
				TotalHours:         project.TotalHours,
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

func GetProjectApplyByProjectIDAndUserID(c *fiber.Ctx) error {
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

	projectID := c.Params("project_id")

	var projectApply entity.ProjectApply
	results := database.DB.Where("project_id = ? AND user_id = ?", projectID, user.ID).First(&projectApply)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	var project entity.Project
	results = database.DB.Where("id = ?", projectApply.ProjectID).First(&project)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	responseDTO := dto.ProjectApplyGetByUserIDResponseDTO{
		ID:     projectApply.ID,
		UserID: user.ID,
		Project: dto.ProjectGetResponseDTO{
			ID:                 project.ID,
			Name:               project.Name,
			ProjectDescription: project.ProjectDescription,
			ShortDescription:   project.ShortDescription,
			MetodeBelajar:      project.MetodeBelajar,
			PeralatanBelajar:   project.PeralatanBelajar,
			Silabus:            project.Silabus,
			TotalHours:         project.TotalHours,
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

	return c.Status(200).JSON(responseDTO)
}

func SendEmailHTML(to []string, subject, templateHTML string, data interface{}) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))

	t, err := template.New("email").Parse(templateHTML)
	if err != nil {
		return err
	}

	err = t.Execute(&body, data)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func SendFinancialAidEmail(c *fiber.Ctx) error {
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
  
	type EmailData struct {
		Username string
		Project  string
	}

	data := EmailData{
		Username: user.FullName,
	}

	err = SendEmailHTML([]string{user.Email}, "Financial Aid Application - EpidemIT", emailTemplate, data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error sending email",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}
  
 func ProjectHandlerGetAllAvailable(c *fiber.Ctx) error {
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
   
	var projects []entity.Project
	results := database.DB.Find(&projects)

	if results.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": results.Error,
		})
	}

	// Filter project yang belum pernah diambil oleh user
	var availableProjects []entity.Project
	for _, project := range projects {
		var projectApplies []entity.ProjectApply
		results := database.DB.Where("project_id = ? AND user_id = ?", project.ID, user.ID).Find(&projectApplies)

		if results.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": results.Error,
			})
		}

		if len(projectApplies) == 0 {
			availableProjects = append(availableProjects, project)
		}
	}

	// get skills from project_skills table
	for i, project := range availableProjects {
		var skills []entity.Skill
		database.DB.Model(&project).Association("Skills").Find(&skills)

		availableProjects[i].Skills = skills
	}

	responseDTO := make([]dto.ProjectGetResponseDTO, len(availableProjects))

	for i, project := range availableProjects {
		responseDTO[i] = dto.ProjectGetResponseDTO{
			ID:                 project.ID,
			Name:               project.Name,
			ProjectDescription: project.ProjectDescription,
			ShortDescription:   project.ShortDescription,
			MetodeBelajar:      project.MetodeBelajar,
			PeralatanBelajar:   project.PeralatanBelajar,
			Silabus:            project.Silabus,
			TotalHours:         project.TotalHours,
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
