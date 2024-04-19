package controllers

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/internal/server/dto"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type LaunchController struct {
	launchRepository domain.ILaunchesRepository
	logger           slog.Logger
}

func NewLaunchController(logger slog.Logger, launchRepository domain.ILaunchesRepository) *LaunchController {
	return &LaunchController{
		launchRepository: launchRepository,

		logger: logger,
	}
}

func (c *LaunchController) GetAllLaunches(w http.ResponseWriter, r *http.Request) {
	launches, err := c.launchRepository.GetLaunches(r.Context())
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.LaunchesToLaunchesDTO(launches))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *LaunchController) GetLaunch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	launch, err := c.launchRepository.GetLaunch(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.Launch(launch))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
