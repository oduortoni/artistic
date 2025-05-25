package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	_ "rtstk/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// If PORT environment variable is set, override the httpport config
	if port := os.Getenv("PORT"); port != "" {
		if p := parsePort(port); p > 0 {
			beego.BConfig.Listen.HTTPPort = p
		}
	}

	// If default port is taken, try to find an available port
	if !isPortAvailable(beego.BConfig.Listen.HTTPPort) {
		for port := 8001; port < 8100; port++ {
			if isPortAvailable(port) {
				beego.BConfig.Listen.HTTPPort = port
				fmt.Printf("Port %d is in use, switching to port %d\n", 8000, port)
				break
			}
		}
	}
}

func isPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	ln.Close()
	return true
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
