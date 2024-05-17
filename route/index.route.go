package route

import (
	"github.com/epidemIT/epidemIT-Backend/handler"
	"github.com/epidemIT/epidemIT-Backend/handler/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/current", middleware.Auth, handler.GetCurrentUser)
	users.Post("/login", handler.UserLogin)
	users.Post("/talent/register", handler.UserRegister)

	mentors := v1.Group("/mentors")
	mentors.Post("/", handler.MentorHandlerCreate)
	mentors.Get("/", handler.MentorHandlerGetAll)
	mentors.Get("/", handler.MentorHandlerGetAll)
	mentors.Get("/:id", handler.MentorHandlerGetByID)
	mentors.Post("/apply/register", middleware.Auth, handler.MentorApplyRegister)

	projects := v1.Group("/projects")
	projects.Post("/", middleware.Auth, handler.ProjectHandlerCreate)
	projects.Get("/", handler.ProjectHandlerGetAll)
	projects.Post("/apply/register", middleware.Auth, handler.ProjectApplyRegister)
	projects.Get("/:id", handler.ProjectHandlerGetByID)
	projects.Post("/apply/financialaid", middleware.Auth, handler.SendFinancialAidEmail)
	projects.Get("/user/available", middleware.Auth, handler.ProjectHandlerGetAllAvailable)

	projectApply := v1.Group("/project-apply")
	projectApply.Get("/user", middleware.Auth, handler.GetAllProjectAppliedByUserID)
	projectApply.Get("/:project_id", middleware.Auth, handler.GetProjectApplyByProjectIDAndUserID)

	mentorApply := v1.Group("/mentor-apply")
	mentorApply.Get("/user", middleware.Auth, handler.GetAllMentorAppliedByUserID)
}
