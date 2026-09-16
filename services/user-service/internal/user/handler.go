package user

import (
	"errors"
	"net/http"
	"services/user-service/internal/dto"
	"services/user-service/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
TODOs
1. Validate service and repository code (superficial code review)
2. implement handler
3. Implement database configuration
4. Implement router + call to relevant configs and modules in main.go
6. test the workflow for creation and search by username
7. upon success implement the remaining endpoints from repository to handler
*/

type UserHandler struct {
	userService IUserService
}

func NewUserHandler(service IUserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) Register(r *gin.Engine) {
	g := r.Group("/api/v1/users")
	{
		g.GET("/:username", h.SearchByUsername)
		g.POST("", h.Create)
		g.PATCH("/:id", h.Update)
		g.DELETE("/:id", h.Delete)
	}
}

func (h *UserHandler) SearchByUsername(c *gin.Context) {
	var userNotFound *UserNotFoundError

	username, ok := parseStringParams(c, "username")

	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid request",
		})
		return
	}

	foundUser, err := h.userService.SearchByUsername(c.Request.Context(), username)

	if err != nil {
		switch {
		case errors.As(err, &userNotFound):
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Success: false,
				Error:   userNotFound.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Data:    toUserResponse(foundUser),
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUpdateUserRequest
	var usernameTaken *UsernameTakenError
	var requiredParameterNotFound *RequiredParameterNotFoundError

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	user, err := h.userService.Create(c.Request.Context(), createUpdateReqToUser(&req))

	if err != nil {
		switch {
		case errors.As(err, &requiredParameterNotFound):
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Success: false,
				Error:   requiredParameterNotFound.Error(),
			})
			return
		case errors.As(err, &usernameTaken):
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Success: false,
				Error:   usernameTaken.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Data:    toUserResponse(user),
		Message: "User created successfully",
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	id, ok := parseUintParams(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid ID Parameter",
		})
		return
	}

	var req dto.CreateUpdateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	err := h.userService.Update(c.Request.Context(), id, createUpdateReqToUser(&req))

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User information updated successfully",
	})
}

// TODO - Test deletion implementation
func (h *UserHandler) Delete(c *gin.Context) {
	var userNotFound *UserNotFoundError

	id, ok := parseUintParams(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid Parameter",
		})
		return
	}

	if err := h.userService.Delete(c.Request.Context(), id); err != nil {
		switch {
		case errors.As(err, &userNotFound):
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}

// Helper methods
func parseStringParams(c *gin.Context, paramName string) (string, bool) {
	parameter := c.Param(paramName)

	if parameter == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Missing parameter " + paramName,
		})

		return "", false
	}

	return parameter, true
}

func parseUintParams(c *gin.Context, paramName string) (uint, bool) {
	parameter := c.Param(paramName)

	if parameter == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Missing parameter " + paramName,
		})
		return 0, false
	}

	id, err := strconv.ParseInt(parameter, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid parameter " + paramName,
		})
		return 0, false
	}

	return uint(id), true
}

// Mappers
func toUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Country:   user.Country,
		// Languages: user.Languages,
		// Dialects: user.Dialects,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		// DeletedAt: user.DeletedAt.Time,
	}
}

func createUpdateReqToUser(req *dto.CreateUpdateUserRequest) *model.User {
	return &model.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Country:   req.Country,
	}
}

/*func updateReqToUser(req *dto.UpdateUserRequest) *model.User {
	return &model.User{
		ID:        req.ID,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Country:   req.Country,
	}
}*/
