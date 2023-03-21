package users

import (
	"net/http"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller for user routes
type UserController struct {
	provider *internal.Provider
}

type AuthController struct {
	provider *internal.Provider
}

// Factory method to create user controller
func NewUserController(p *internal.Provider) (m UserController) {
	m = UserController{
		provider: p,
	}

	if p == nil {
		p.Logger.Errorln(h.InitError)
	}
	return
}

func NewAuthController(p *internal.Provider) (m AuthController) {
	m = AuthController{
		provider: p,
	}

	if p == nil {
		p.Logger.Errorln(h.InitError)
	}
	return
}

func (o UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var d domain.UserCreateData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	id, err := o.provider.UserRepo.CreateUser(&d)
	if err != nil {
		o.provider.Logger.Errorln(h.MongoError, err.Error())
		h.JsonRes(w, http.StatusInternalServerError, h.Json{
			"code":    h.InternalServerError,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	h.JsonRes(w, http.StatusOK, h.Json{
		"id": id,
	}, o.provider.Logger)
}

func (o UserController) getUserById(w http.ResponseWriter, r *http.Request) {
	query := domain.UserQueryData{}

	if id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "userId")); err != nil {
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		query.Id = id
	}

	if user, err := o.provider.UserRepo.GetOneUser(&query); err != nil {
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, user, o.provider.Logger)
		return
	}
}

func (o UserController) getUsers(w http.ResponseWriter, r *http.Request) {
	query := domain.UserQueryData{}

	limit, skip := h.GetLimitSkip(r)
	if limit != 0 {
		query.Limit = limit
	}
	if skip != 0 {
		query.Skip = skip
	}

	if users, err := o.provider.UserRepo.GetUsers(&query); err != nil {
		h.JsonRes(w, http.StatusInternalServerError, h.Json{
			"code":    h.InternalServerError,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, users, o.provider.Logger)
		return
	}
}

func (o AuthController) googleLogin(w http.ResponseWriter, r *http.Request) {
	var d domain.GoogleLoginPostData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	user, err := o.provider.AuthRepo.GoogleLogin(&d)
	if err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    "ERR_AUTH_TODO",
			"message": "Invalid credential",
		}, o.provider.Logger)
		return
	}

	h.JsonRes(w, http.StatusOK, user, o.provider.Logger)
	return
}

func (o AuthController) logout(w http.ResponseWriter, r *http.Request) {
	h.JsonRes(w, http.StatusOK, nil, o.provider.Logger)
}

func (o UserController) getSelf(w http.ResponseWriter, r *http.Request) {
	user, err := h.GetUserIdContext(r)
	if err != nil {
		o.provider.Logger.Infoln("ERR_AUTH")
		h.JsonRes(w, http.StatusUnauthorized, h.Json{
			"code": "ERR_AUTH",
		}, o.provider.Logger)
		return
	}

	o.provider.Logger.Infoln("OK_USER_GET-SELF")
	h.JsonRes(w, http.StatusOK, user, o.provider.Logger)
	return
}
