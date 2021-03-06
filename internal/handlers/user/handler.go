package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tmolyakov/go-api-xmp/internal/handlers"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersUrl, h.GetList)
	router.GET(userUrl, h.GetByUuid)
	router.POST(userUrl, h.Create)
	router.PUT(userUrl, h.Update)
	router.PATCH(userUrl, h.PartiallyUpdate)
	router.DELETE(userUrl, h.Delete)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetByUuid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is requested user"))
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user created"))
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user udpated"))
}

func (h *handler) PartiallyUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user partially updated"))
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user deleted"))
}
