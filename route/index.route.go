package route

import (
	"epidemIT-Backend/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// users := v1.Group("/users")
	// users.Get("/current", middleware.Auth, handler.GetCurrentUser)
	// users.Post("/login", handler.UserLogin)
	// users.Post("/talent/register", handler.UserTalentRegister)
	// users.Post("/partner/register", handler.UserPartnerRegister)
	// users.Post("/mentor/register", middleware.Auth, handler.UserMentorRegister)

	// partners := v1.Group("/partners")
	// partners.Get("/", middleware.Auth, handler.PartnerHandlerGetAll)

	mentors := v1.Group("/mentors")
	mentors.Get("/", handler.MentorHandlerGetAll)
	mentors.Post("/:id/assign/:idProject", handler.MentorAssign)

	// projects := v1.Group("/projects")
	// projects.Get("/", middleware.Auth, handler.ProjectHandlerGetAll)
	// projects.Get("/available", middleware.Auth, handler.ProjectHandlerGetAllAvailable)
	// projects.Post("/", middleware.Auth, handler.ProjectRegister)
	// projects.Post("/:id/apply/register", middleware.Auth, handler.ProjectApplyRegister)

}
