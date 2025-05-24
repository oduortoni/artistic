package main

import (
	"os"
	"strconv"

	_ "artistic/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// If PORT environment variable is set, override the httpport config
	if port := os.Getenv("PORT"); port != "" {
		if p := parsePort(port); p > 0 {
			beego.BConfig.Listen.HTTPPort = p
		}
	}
}

func parsePort(port string) int {
	// Default port if parsing fails
	defaultPort := 8000
	if p, err := beego.AppConfig.Int("httpport"); err == nil {
		defaultPort = p
	}

	// Try to parse the port
	if p, err := strconv.Atoi(port); err == nil && p > 0 {
		return p
	}

	return defaultPort
}

func main() {
	beego.Run()
}
