package controllers

import (
	domain "TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/internal/server/dto"
	"TestTask-PGPro/lib/byteconv"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type CommandController struct {
	commandRepository domain.ICommandsRepository
	logger            slog.Logger
}

func NewCommandController(logger slog.Logger, commandRepository domain.ICommandsRepository) *CommandController {
	return &CommandController{
		commandRepository: commandRepository,

		logger: logger,
	}
}

func (c *CommandController) CreateCommandHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AddCommandRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := c.commandRepository.AddCommand(r.Context(), req.Text)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.Write(byteconv.Bytes(strconv.Itoa(id)))
}

func (c *CommandController) GetAllCommandsHandler(w http.ResponseWriter, r *http.Request) {
	commands, err := c.commandRepository.GetCommands(r.Context())
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.CommandsToCommandsDTO(commands))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *CommandController) GetCommandHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	command, err := c.commandRepository.GetCommand(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	response, err := json.Marshal(dto.Command(command))
	if err != nil {
		c.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *CommandController) DeleteCommandHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.commandRepository.DeleteCommand(r.Context(), id)
	if err != nil {
		c.logger.Error(err.Error())
		apiErr := dto.NewApiError(err)
		w.WriteHeader(apiErr.StatusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
}
