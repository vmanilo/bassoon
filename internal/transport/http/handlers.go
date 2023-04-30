package http

import (
	"net/http"

	"bassoon/internal/app/model"

	"github.com/go-chi/chi/v5"
)

func (s *server) createPort(writer http.ResponseWriter, request *http.Request) {
	var port model.Port
	if err := getRequestData(request, &port); err != nil {
		httpError(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if err := s.service.StorePort(request.Context(), &port); err != nil {
		httpError(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (s *server) getPort(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	port, err := s.service.RetrievePort(request.Context(), id)
	if err != nil {
		httpError(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := sendResponse(writer, port); err != nil {
		httpError(writer, err.Error(), http.StatusInternalServerError)
	}
}
