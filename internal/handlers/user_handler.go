package handlers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"user-management-api/internal/models"
	"user-management-api/internal/service"
	"user-management-api/pkg/logger"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a user with name and date of birth
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user, err := h.service.CreateUser(c.Context(), &req)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Error:   err.Error(),
		})
	}

	age := h.service.CalculateAgeForUser(&user)

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}

	requestID := c.Locals("request_id").(string)
	logger.Log.Info("User created successfully",
		zap.Int32("id", user.ID),
		zap.String("request_id", requestID))

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user details by ID with calculated age
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
	}
	id := int32(id64)
	user, err := h.service.GetUserByID(c.Context(), id)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Success: false,
				Message: "User not found",
			})
		}
		logger.Log.Error("Failed to get user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	age := h.service.CalculateAgeForUser(&user)

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user name and/or date of birth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
	}
	id := int32(id64)

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user, err := h.service.UpdateUser(c.Context(), id, &req)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Success: false,
				Message: "User not found",
			})
		}
		logger.Log.Error("Failed to update user", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Error:   err.Error(),
		})
	}

	age := h.service.CalculateAgeForUser(&user)

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  age,
	}

	requestID := c.Locals("request_id").(string)
	logger.Log.Info("User updated successfully",
		zap.Int32("id", user.ID),
		zap.String("request_id", requestID))

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id64, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
		})
	}
	id := int32(id64)
	err = h.service.DeleteUser(c.Context(), id)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Success: false,
				Message: "User not found",
			})
		}
		logger.Log.Error("Failed to delete user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	requestID := c.Locals("request_id").(string)
	logger.Log.Info("User deleted successfully",
		zap.Int32("id", id),
		zap.String("request_id", requestID))

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// ListUsers godoc
// @Summary List all users
// @Description Get paginated list of users with calculated ages
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(20)
// @Success 200 {object} models.ListUsersResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	users, total, err := h.service.ListUsers(c.Context(), page, pageSize)
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	var responses []models.UserResponse
	for _, user := range users {
		age := h.service.CalculateAgeForUser(&user)
		responses = append(responses, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  age,
		})
	}

	response := models.ListUsersResponse{
		Users:      responses,
		TotalCount: int(total),
		Page:       page,
		PageSize:   pageSize,
	}

	requestID := c.Locals("request_id").(string)
	logger.Log.Info("Users listed successfully",
		zap.Int("count", len(responses)),
		zap.Int64("total", total),
		zap.String("request_id", requestID))

	return c.Status(fiber.StatusOK).JSON(response)
}