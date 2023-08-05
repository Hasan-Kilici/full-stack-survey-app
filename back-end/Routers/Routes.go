package Routers

import(
	"survey/Middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"survey/Handlers"
)

func Initalize(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", Handlers.HelloWorld)
	app.Get("/find/user/:Token", Handlers.FindUser)
	app.Get("/find/survey/:Token", Handlers.FindSurvey)
	app.Get("/find/options/:SurveyToken", Handlers.FindOptions)
	app.Get("/surveys", Handlers.ListSurveys)

	app.Post("/api/login", Handlers.Login)
	app.Post("/api/register", Handlers.Register)
	app.Post("/api/create/survey", Handlers.CreateSurvey)
	app.Post("/give/vote/:SurveyToken/:Token", Handlers.VoteOption)
	app.Post("/delete/survey/:Token", Handlers.DeleteSurvey)

	app.Use(Middleware.Security)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}