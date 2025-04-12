package handlers

import (
	"context"
	"diplom/user-service/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) HealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}

func (h *Handler) SendCode(c fiber.Ctx) error {
	body := c.Body()
	var req models.SendCodeRequest

	err := json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error("Error unmarshalling body", "error", err)
		return c.Status(400).JSON(err.Error())
	}

	err = h.services.SendCode(req.Gmail)
	if err != nil {
		h.logger.Error("Error sending code", "error", err)
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON("OK")
}

func (h *Handler) CheckCode(c fiber.Ctx) error {
	body := c.Body()
	var req models.CheckCodeRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error("Error unmarshalling body", "error", err)
		return c.Status(400).JSON(err.Error())
	}

	ctx := context.Background()
	isCode, err := h.services.CheckCode(ctx, req.Gmail, req.Code)
	if err != nil {
		h.logger.Error("Error checking code", "error", err)
		return c.Status(500).JSON(err.Error())
	}

	if !isCode {
		return c.Status(400).JSON(fiber.Map{"error": "code not exist"})
	}

	return c.Status(200).JSON("OK")
}
