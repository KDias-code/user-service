package handlers

import (
	"diplom/user-service/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) SaveUser(c fiber.Ctx) error {
	body := c.Body()
	var user models.SaveUserRequest
	err := json.Unmarshal(body, &user)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	if user.Name == "" || user.StudentId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name or student_id is empty",
		})
	}

	err = h.services.AddUser(user)
	if err != nil {
		h.logger.Error("Error while saving user", "error", err)
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON("OK")
}
