package http

import (
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	svc ports.BookService
}

func NewBookHandler(svc ports.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	var body struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	b, err := h.svc.Create(body.Title, body.Author)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(b)
}

func (h *BookHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	b, err := h.svc.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(b)
}

func (h *BookHandler) List(c *fiber.Ctx) error {
	items, err := h.svc.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *BookHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	b, err := h.svc.Update(id, body.Title, body.Author)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(b)
}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
