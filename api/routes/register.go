package routes

import (
	"connectrpc.com/vanguard"
	"github.com/betterde/focusly/docs"
	"github.com/betterde/focusly/internal/gen/svc/v1/svcv1connect"
	"github.com/betterde/focusly/internal/grpc/services"
	"github.com/betterde/focusly/internal/journal"
	"github.com/betterde/focusly/internal/response"
	"github.com/betterde/focusly/internal/utils"
	"github.com/betterde/focusly/spa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response.Success("Success", nil))
	}).Name("Health check")

	// Swagger API specification file router
	app.Get("/swagger/*", filesystem.New(filesystem.Config{
		Root:               docs.Serve(),
		Index:              "openapi.yaml",
		NotFoundFile:       "openapi.yaml",
		ContentTypeCharset: "UTF-8",
	})).Name("Swagger JSON Schema")

	// Swagger UI router
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL:          "/swagger/openapi.yaml",
		DeepLinking:  false,
		DocExpansion: "none",
	})).Name("Swagger UI")

	userService := &services.UserService{}

	path, handler := svcv1connect.NewUserServiceHandler(userService)
	transcoder, err := vanguard.NewTranscoder([]*vanguard.Service{vanguard.NewService(path, handler)})
	if err != nil {
		journal.Logger.Panic(err)
	}
	app.Post("/v1/*", adaptor.HTTPHandler(utils.CamelToSnake(transcoder))).Name("User service route") // Embed SPA static resource

	app.Get("*", filesystem.New(filesystem.Config{
		Root:               spa.Serve(),
		Index:              "index.html",
		NotFoundFile:       "index.html",
		ContentTypeCharset: "UTF-8",
	})).Name("SPA static resource")
}
