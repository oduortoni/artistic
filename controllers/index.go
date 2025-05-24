package controllers

import (
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	// Try to serve index.html from the frontend/dist directory
	indexPath := filepath.Join("frontend", "dist", "index.html")
	if _, err := os.Stat(indexPath); err == nil {
		c.Ctx.Output.Header("Content-Type", "text/html")
		c.Ctx.Output.Header("Cache-Control", "no-cache")
		c.Ctx.Output.Download(indexPath, "index.html")
		return
	}

	// Fallback response if index.html is not found
	c.Data["json"] = map[string]string{"error": "Frontend not built"}
	c.ServeJSON()
}
