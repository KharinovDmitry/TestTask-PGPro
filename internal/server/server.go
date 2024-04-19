package server

import (
	domain "TestTask-PGPro/internal/domain/service"
	"TestTask-PGPro/internal/server/controllers"
	"TestTask-PGPro/internal/storage"
	"log/slog"
	"net/http"
)

func MustRun(logger slog.Logger, launchService domain.ILaunchService, storage storage.Storage, address string) {
	commandController := controllers.NewCommandController(logger, storage.CommandsRepository)
	launchController := controllers.NewLaunchController(logger, storage.LaunchesRepository)

	runController := controllers.NewRunController(logger, launchService)
	stopController := controllers.NewStopController(logger, launchService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /commands", commandController.GetAllCommandsHandler)
	mux.HandleFunc("GET /command/{id}", commandController.GetCommandHandler)
	mux.HandleFunc("POST /command", commandController.CreateCommandHandler)
	mux.HandleFunc("DELETE /command/{id}", commandController.DeleteCommandHandler)

	mux.HandleFunc("GET /launches", launchController.GetAllLaunches)
	mux.HandleFunc("GET /launch/{id}", launchController.GetLaunch)

	mux.HandleFunc("POST /run/{id}", runController.Run)
	mux.HandleFunc("POST /stop/{id}", stopController.Stop)

	if err := http.ListenAndServe(address, mux); err != nil {
		panic(err.Error())
	}

}
