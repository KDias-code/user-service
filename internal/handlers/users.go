package handlers

import (
	"bytes"
	"diplom/user-service/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"io"
	"mime/multipart"
	"net/textproto"
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

func (h *Handler) UpdateUser(c fiber.Ctx) error {
	req := new(models.SaveUserRequest)

	name := c.FormValue("name")
	if name != "" {
		req.Name = name
	}
	studId := c.FormValue("student_id")
	if studId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "student_id is empty",
		})
	} else {
		req.StudentId = studId
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		h.logger.Warn("Failed to get avatar", "error", err)
	} else if avatar != nil {
		src, err := avatar.Open()
		if err != nil {
			h.logger.Error("Failed to get avatar", "error", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid file",
			})
		}
		defer src.Close()

		imageBytes, err := io.ReadAll(src)
		if err != nil {
			h.logger.Error("Failed to read avatar", "error", err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Invalid file",
			})
		}

		req.Avatar = imageBytes
	}

	err = h.services.UpdateUser(*req)
	if err != nil {
		h.logger.Error("Error while saving user", "error", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(200).JSON("OK")
}

func (h *Handler) GetUser(c fiber.Ctx) error {
	id := c.Params("student_id")
	if id == "" {
		h.logger.Error("Failed to get user", "error", "student_id is empty")
		return c.Status(400).JSON(fiber.Map{
			"error": "student_id is empty",
		})
	}

	user, err := h.services.GetUser(id)
	if err != nil {
		h.logger.Error("Error while getting user", "error", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get user",
		})
	}

	boundary := "MyBoundary"
	c.Set("Content-Type", "multipart/mixed; boundary="+boundary)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.SetBoundary(boundary)

	jsonPart, _ := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"application/json"},
	})
	jsonData := fmt.Sprintf(`{"student_id": %s, "name": "%s"}`, user.StudentId, user.Name)
	jsonPart.Write([]byte(jsonData))

	imagePart, _ := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"image/jpeg"},
	})
	imagePart.Write(user.Avatar)

	writer.Close()

	return c.Send(body.Bytes())
}
