package api

import (
	"flag"
	"net/url"
	"os"

	e "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var flagHTTPListenAddr = flag.String("http-listen-port", ":8080", "HTTP listen address, if blacnk will default to ENV PORT")

type ApiServer struct {
	Zlog *zap.Logger
	Echo *e.Echo
}

// NewApiServer creates a new ApiServer with logger
func NewApiServer(l *zap.Logger, e *e.Echo) *ApiServer {
	return &ApiServer{

		Zlog: l,
		Echo: e,
	}

}

//Start http server on port 8080 if no port is specified in the env variable PORT
func (s *ApiServer) Start() {
	s.Echo.Logger.Info("Starting API Server")

	// Bind the route to the handler
	setRoutes(s)

	// User recover middleware to handle panics
	s.Echo.Use(middleware.Recover())

	// Start the server
	if os.Getenv("PORT") == "" {
		s.Echo.Logger.Info("on port 8080")
		s.Echo.Logger.Fatal(s.Echo.Start(*flagHTTPListenAddr))
	} else {
		//Check if port is valid by using url.Parse a fake url
		tmpUrl := "http://localhost:" + os.Getenv("PORT") + "/tmpurl"
		parsedUrl, err := url.Parse(tmpUrl)
		if err != nil {
			s.Echo.Logger.Fatal("Invalid port:", err)
		}
		s.Echo.Logger.Info("on port", parsedUrl.Port())
		s.Echo.Logger.Fatal(s.Echo.Start(":" + parsedUrl.Port()))
	}

}

//SetLogger sets the logger for the server
func (s *ApiServer) SetLogger(l *zap.Logger) {
	s.Zlog = l
}

//SetEcho sets the echo server for the server
func (s *ApiServer) SetEcho(e *e.Echo) {
	s.Echo = e
}
