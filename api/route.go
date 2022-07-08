package api

import (
	"github.com/labstack/echo/v4"
)

func setRoutes(s *ApiServer) {
	e := s.Echo

	// Hello world ... :)
	e.GET("/", helloworld)

	//Ping responds with pong for health checking
	e.GET("/ping", ping)

	//GetTokens responds with a list of tokens
	e.GET("/tokens", GetTokens)
}

func helloworld(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

func ping(c echo.Context) error {
	return c.String(200, "pong")
}
