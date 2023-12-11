package dash

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/lofustudio/VEGA/dash/template"
)

func Start() func() error {
	dash := fiber.New(fiber.Config{
		AppName:               "Dash",
		DisableStartupMessage: true,
	})

	dash.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	dash.All("/*", filesystem.New(filesystem.Config{
		Root:         template.Dist(),
		NotFoundFile: "404.html",
		Index:        "index.html",
	}))

	if err := dash.Listen(":3000"); err != nil {
		panic(err)
	}

	return dash.Shutdown
}
