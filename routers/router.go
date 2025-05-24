package routers

import (
	"os"

	"artistic/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// API routes
	beego.Router("/api/users", &controllers.UsersController{}, "get:GetUsers")

	// Serve static files from the frontend/dist directory
	frontendDir := "frontend/dist"
	if _, err := os.Stat(frontendDir); err == nil {
		beego.SetStaticPath("/", frontendDir)
	}

	// Handle all other routes by serving index.html (for client-side routing)
	beego.Router("/*", &controllers.MainController{}, "get:Index")
}
