package handlers

import (
	"diplom/user-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
)

type IHandler interface {
	HealthCheck(c fiber.Ctx) error
	SendCode(c fiber.Ctx) error
	CheckCode(c fiber.Ctx) error
}
type Handler struct {
	logger   hclog.Logger
	services service.IService
}

func NewHandler(logger hclog.Logger, service service.IService) *Handler {
	return &Handler{
		logger:   logger,
		services: service,
	}
}
