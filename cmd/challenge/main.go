package main

import (
	"backend-challenge-go/api"
	"backend-challenge-go/eth"
	"backend-challenge-go/eth/rpc"
	"backend-challenge-go/logging"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	zlog = logging.MustCreateLoggerWithServiceName("challenge")
	rpc.SetLogger(zlog)
	eth.SetLogger(zlog)
}

func main() {
	// cron of loading token info
	StartCronJob()

	a := api.NewApiServer(zlog, echo.New())
	a.Start()
}
