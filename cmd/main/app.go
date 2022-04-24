package main

import (
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tmolyakov/go-api-xmp/internal/handlers/user"
	"github.com/tmolyakov/go-api-xmp/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("starting application...")
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening at 0.0.0.0:1234")

	logger.Fatal(server.Serve(listener))
}
