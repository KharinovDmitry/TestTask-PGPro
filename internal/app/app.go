package app

import (
	"TestTask-PGPro/cmd/migrator"
	"TestTask-PGPro/internal/config"
	"TestTask-PGPro/internal/logger"
	"TestTask-PGPro/internal/server"
	"TestTask-PGPro/internal/service"
	"TestTask-PGPro/internal/storage"
	adapter "TestTask-PGPro/lib/adapter/db"
	"TestTask-PGPro/lib/adapter/executor"
	"context"
)

func MustRun(cfg *config.Config) {
	logger, err := logger.SetupLogger(cfg.Env)
	if err != nil {
		panic(err.Error())
	}

	migrator.MustRun(cfg.DriverName, cfg.ConnStr, cfg.MigrationsDir)
	db := adapter.NewPostgresAdapter(cfg.TimeoutDB)
	db.Connect(context.Background(), cfg.ConnStr)
	defer db.Close()

	store := storage.NewStorage(db)
	store.InitStorage()

	executor := executor.NewLinuxAdapter()

	launchService := service.NewLaunchService(*logger, executor, store.CommandsRepository, store.LaunchesRepository)

	server.MustRun(*logger, launchService, *store, cfg.Address)
}
