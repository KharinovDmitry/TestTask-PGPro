package controllers

import (
	"TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/dto"
	"TestTask-PGPro/lib/byteconv"
	"log/slog"
	"net/http"
	"strconv"
)

type RunController struct {
	logger        slog.Logger
	launchService domain.ILaunchService
}

func NewRunController(logger slog.Logger, launchService domain.ILaunchService) *RunController {
	return &RunController{
		launchService: launchService,
		logger:        logger,
	}
}

func (c *RunController) Run(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	launchId, err := c.launchService.Launch(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.Write(byteconv.Bytes(strconv.Itoa(launchId)))
}
