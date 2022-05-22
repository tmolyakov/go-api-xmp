package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tmolyakov/go-api-xmp/internal/config"
	"github.com/tmolyakov/go-api-xmp/internal/handlers/user"
	"github.com/tmolyakov/go-api-xmp/internal/handlers/user/db"
	"github.com/tmolyakov/go-api-xmp/pkg/client/mongodb"
	"github.com/tmolyakov/go-api-xmp/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(
		context.Background(),
		cfgMongo.Host,
		cfgMongo.Port,
		cfgMongo.Username,
		cfgMongo.Password,
		cfgMongo.Database,
		cfgMongo.AuthDB,
	)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	user1 := user.User{
		ID:           "",
		Email:        "test@test.com",
		Username:     "testUser",
		PasswordHash: "asdfas42345234",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("starting application...")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening at %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
