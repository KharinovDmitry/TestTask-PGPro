package controllers

import (
	domain "TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/dto"
	"log/slog"
	"net/http"
	"strconv"
)

type StopController struct {
	logger        slog.Logger
	launchService domain.ILaunchService
}

func NewStopController(logger slog.Logger, launchService domain.ILaunchService) *StopController {
	return &StopController{
		launchService: launchService,
		logger:        logger,
	}
}

func (c *StopController) Stop(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.launchService.Stop(r.Context(), id)
	if err != nil {
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}
