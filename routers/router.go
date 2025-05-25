package routers

import (
	"os"

	"rtstk/controllers"

	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
)

func init() {
	// API routes
	beego.Router("/api/users", &controllers.UsersController{}, "get:GetUsers")

	// Enable CORS for development
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.SetStatus(200)
			ctx.Output.JSON(map[string]interface{}{}, false, false)
		}
	})

	// Serve static files from the frontend/dist directory
	frontendDir := "frontend/dist"
	if _, err := os.Stat(frontendDir); err == nil {
		beego.SetStaticPath("/", frontendDir)
	}

	// Handle all other routes by serving index.html (for client-side routing)
	beego.Router("/*", &controllers.MainController{}, "get:Index")
}
