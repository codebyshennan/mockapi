package mockservers

import (
	"net/http"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockServerController struct {
	provider *internal.Provider
}

func NewMockServerController(p *internal.Provider) (m *MockServerController) {
	m = &MockServerController{}
	if p == nil {
		return
	}
	m.provider = p
	return
}

func (o MockServerController) createMockServer(w http.ResponseWriter, r *http.Request) {
	user, err := h.GetUserIdContext(r)
	if err != nil {
		h.JsonRes(w, http.StatusUnauthorized, h.Json{
			"code": "ERR_AUTH_TODO",
		}, o.provider.Logger)
		return
	}

	var d domain.MockServerCreateData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	d.OwnerRef = user.Id
	id, err := o.provider.MockServerRepo.CreateMockServer(&d)
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

func (o MockServerController) getMockServerById(w http.ResponseWriter, r *http.Request) {
	query := domain.MockServerQueryData{}

	if id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "mockServerId")); err != nil {
		o.provider.Logger.Errorln("ERR_MOCK-SERVER_TODO", err.Error())
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		query.Id = id
	}

	if mockServer, err := o.provider.MockServerRepo.GetOneMockServer(&query); err != nil {
		o.provider.Logger.Errorln("ERR_MOCK-SERVER_TODO", err.Error())
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, mockServer, o.provider.Logger)
		return
	}
}

func (o MockServerController) getMockServers(w http.ResponseWriter, r *http.Request) {
	query := domain.MockServerQueryData{}

	limit, skip := h.GetLimitSkip(r)
	if limit != 0 {
		query.Limit = limit
	}
	if skip != 0 {
		query.Skip = skip
	}

	if mockServers, err := o.provider.MockServerRepo.GetMockServers(&query); err != nil {
		h.JsonRes(w, http.StatusInternalServerError, h.Json{
			"code":    h.InternalServerError,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, mockServers, o.provider.Logger)
		return
	}
}

func (o MockServerController) updateMockServer(w http.ResponseWriter, r *http.Request) {
	var d domain.MockServerUpdateData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "mockServerId"))
	if err != nil {
		o.provider.Logger.Errorln("ERR_MOCK-SERVER_TODO", err.Error())
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	}

	if err = o.provider.MockServerRepo.UpdateMockServer(id, &d); err != nil {
		o.provider.Logger.Errorln("ERR_MOCK-SERVER_TODO", err.Error())
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, nil, o.provider.Logger)
		return
	}
}
