package services

import (
	"net/http"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceController struct {
	provider *internal.Provider
}

type SwaggerController struct {
	provider *internal.Provider
}

func NewServiceController(p *internal.Provider) (m ServiceController) {
	m = ServiceController{
		provider: p,
	}
	if p == nil {
		p.Logger.Errorln(h.InitError)
	}
	return
}

func NewSwaggerController(p *internal.Provider) (m SwaggerController) {
	m = SwaggerController{
		provider: p,
	}
	if p == nil {
		p.Logger.Errorln(h.InitError)
	}
	return
}

func (o ServiceController) createService(w http.ResponseWriter, r *http.Request) {
	var d domain.ServiceCreateData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	id, err := o.provider.ServiceRepo.CreateService(&d)
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

func (o ServiceController) getServices(w http.ResponseWriter, r *http.Request) {
	query := domain.ServiceQueryData{}

	limit, skip := h.GetLimitSkip(r)
	if limit != 0 {
		query.Limit = limit
	}
	if skip != 0 {
		query.Skip = skip
	}

	if svcs, err := o.provider.ServiceRepo.GetServices(&query); err != nil {
		h.JsonRes(w, http.StatusInternalServerError, h.Json{
			"code":    h.InternalServerError,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, svcs, o.provider.Logger)
		return
	}
}

func (o ServiceController) getServiceById(w http.ResponseWriter, r *http.Request) {
	query := domain.ServiceQueryData{}

	if id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "serviceId")); err != nil {
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		query.Id = id
	}

	if svc, err := o.provider.ServiceRepo.GetOneService(&query); err != nil {
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, svc, o.provider.Logger)
		return
	}
}

func (o SwaggerController) createSwagger(w http.ResponseWriter, r *http.Request) {
	var d domain.SwaggerCreateData
	if err := h.BindJson(r, &d); err != nil {
		h.JsonRes(w, http.StatusBadRequest, h.Json{
			"code":    h.BadRequest,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}

	user, err := h.GetUserIdContext(r)
	if err != nil {
		h.JsonRes(w, http.StatusUnauthorized, h.Json{
			"code":    "ERR_SVC_TODOD",
			"message": err.Error(),
		}, o.provider.Logger)
		return
	}
	d.OwnerRef = user.Id

	id, err := o.provider.SwaggerRepo.CreateSwagger(&d)
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

func (o SwaggerController) getSwaggers(w http.ResponseWriter, r *http.Request) {
	query := domain.SwaggerQueryData{}

	limit, skip := h.GetLimitSkip(r)
	if limit != 0 {
		query.Limit = limit
	}
	if skip != 0 {
		query.Skip = skip
	}

	if swaggers, err := o.provider.SwaggerRepo.GetSwaggers(&query); err != nil {
		h.JsonRes(w, http.StatusInternalServerError, h.Json{
			"code":    h.InternalServerError,
			"message": err.Error(),
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, swaggers, o.provider.Logger)
		return
	}
}

func (o SwaggerController) getSwaggerById(w http.ResponseWriter, r *http.Request) {
	query := domain.SwaggerQueryData{}

	if id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "swaggerId")); err != nil {
		o.provider.Logger.Errorln("ERR_SWAGGER_TODO", err.Error())
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		query.Id = id
	}

	if swagger, err := o.provider.SwaggerRepo.GetOneSwagger(&query); err != nil {
		o.provider.Logger.Errorln("ERR_SWAGGER_TODO", err.Error())
		h.JsonRes(w, http.StatusNotFound, h.Json{
			"code":    h.NotFoundError,
			"message": "",
		}, o.provider.Logger)
		return
	} else {
		h.JsonRes(w, http.StatusOK, swagger, o.provider.Logger)
		return
	}
}
