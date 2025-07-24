package routes

import (
	"budget-tracker/application"
	"budget-tracker/internal/logger"

	"github.com/gofiber/fiber/v2"
)

type ExpenseRoutes struct {
	CreateHandler *application.CreateExpenseHandler
}

func RegisterExpenseRoutes(app *fiber.App, handler *application.CreateExpenseHandler) {
	routes := &ExpenseRoutes{
		CreateHandler: handler,
	}

	expenses := app.Group("/api/expenses")
	expenses.Post("", routes.CreateExpense)
	// expenses.Get("/:id", routes.GetExpense) // ileride eklenebilir
}

func (r *ExpenseRoutes) CreateExpense(c *fiber.Ctx) error {
	var req application.CreateExpenseCommand
	if err := c.BodyParser(&req); err != nil {
		logger.Error("Geçersiz istek", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	resp, err := r.CreateHandler.Handle(&req)
	if err != nil {
		logger.Error("Harcama eklenemedi", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create expense",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": resp.ID,
	})
}
