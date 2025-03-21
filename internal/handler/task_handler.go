package handler

import (
	"skillsrock-todo/internal/model"
	"skillsrock-todo/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(service  service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: service}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[ERROR] Task not found"})
	}

	if err := h.taskService.CreateTask(c.Context(), &task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "[ERROR] Cannot create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	tasks, err := h.taskService.GetTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "[ERROR] Cannot get tasks"})
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[ERROR] Invalid ID"})
	}

	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[ERROR] Invalid request payload"})
	}
	task.ID = int64(id)

	if err := h.taskService.UpdateTask(c.Context(), &task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "[ERROR] Cannot update task"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[ERROR] Invalid ID"})
	}

	if err := h.taskService.DeleteTask(c.Context(), int64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "[ERROR] Cannot delete task"}) 
	}

	return c.SendStatus(fiber.StatusNoContent)
}