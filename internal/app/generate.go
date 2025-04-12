package app

import (
	"diplom/user-service/internal/handlers"
	"diplom/user-service/internal/service"
	"diplom/user-service/internal/store"
	"diplom/user-service/pkg/configs"
	"diplom/user-service/pkg/db"
	"diplom/user-service/pkg/gmail"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
	"strconv"
)

type server struct {
	app      *fiber.App
	logger   hclog.Logger
	handlers handlers.IHandler
}

func (s *server) generate(conf *configs.Configs) {
	s.app = fiber.New()

	s.logger = hclog.New(&hclog.LoggerOptions{
		Level:      2,
		JSONFormat: true,
	})

	dbConnection, err := db.ConnectPostgres(conf.Db)
	if err != nil {
		s.logger.Error("Failed to connect to database", "error", err)
		return
	}

	redisDb, err := strconv.Atoi(conf.RedisDB)
	if err != nil {
		s.logger.Error("Failed to convert database to int", "error", err)
		return
	}

	redisConnection := db.ConnectRedis(conf.RedisHost, conf.RedisPort, conf.RedisPass, redisDb)

	gmails := gmail.NewGmail(conf.GmailLogin, conf.GmailPass, conf.GmailHost, conf.GmailPort)

	stores := store.NewStore(redisConnection, dbConnection)
	services := service.NewService(gmails, stores)
	s.handlers = handlers.NewHandler(s.logger, services)

	s.router()
}
