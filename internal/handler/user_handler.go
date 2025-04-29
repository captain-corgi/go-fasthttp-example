package handler

import (
	"encoding/json"

	"github.com/captain-corgi/go-fasthttp-example/internal/domain/model"
	"github.com/captain-corgi/go-fasthttp-example/internal/service"
	"github.com/valyala/fasthttp"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// HandleGetUser handles GET requests for user retrieval
func (h *UserHandler) HandleGetUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	user, err := h.userService.GetUser(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	if user == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// HandleCreateUser handles POST requests for user creation
func (h *UserHandler) HandleCreateUser(ctx *fasthttp.RequestCtx) {
	var user model.User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := h.userService.CreateUser(&user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// HandleUpdateUser handles PUT requests for user updates
func (h *UserHandler) HandleUpdateUser(ctx *fasthttp.RequestCtx) {
	var user model.User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	user.ID = ctx.UserValue("id").(string)
	if err := h.userService.UpdateUser(&user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// HandleDeleteUser handles DELETE requests for user deletion
func (h *UserHandler) HandleDeleteUser(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	if err := h.userService.DeleteUser(id); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

// HandleGetAllUsers handles GET requests for retrieving all users
func (h *UserHandler) HandleGetAllUsers(ctx *fasthttp.RequestCtx) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(users)
}
