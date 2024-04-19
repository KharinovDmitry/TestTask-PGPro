package service

import (
	"TestTask-PGPro/internal/domain/repository"
	"TestTask-PGPro/lib/adapter/executor"
	"TestTask-PGPro/lib/byteconv"
	"context"
	"github.com/pkg/errors"
	"log/slog"
	"sync"
)

type LaunchService struct {
	logger             slog.Logger
	executor           executor.LinuxAdapter
	commandsRepository domain.ICommandsRepository
	launchesRepository domain.ILaunchesRepository

	runs sync.Map
}

func NewLaunchService(logger slog.Logger, executor executor.LinuxAdapter, commandsRepository domain.ICommandsRepository, launchesRepository domain.ILaunchesRepository) *LaunchService {
	return &LaunchService{
		logger:             logger,
		executor:           executor,
		commandsRepository: commandsRepository,
		launchesRepository: launchesRepository,
	}
}

func (l *LaunchService) Launch(ctx context.Context, commandId int) (int, error) {
	command, err := l.commandsRepository.GetCommand(ctx, commandId)
	if err != nil {
		return 0, err
	}

	launchId, err := l.launchesRepository.AddLaunch(ctx, commandId)
	if err != nil {
		return 0, err
	}

	output := make(chan []byte)
	runCtx, stop := context.WithCancel(context.Background())
	l.runs.Store(launchId, stop)
	go l.execute(runCtx, output, command.Text)
	go l.writeOutput(ctx, output, launchId)

	return launchId, nil
}

func (l *LaunchService) execute(ctx context.Context, output chan []byte, command string) {
	if err := l.executor.Run(ctx, command, output); err != nil {
		l.logger.Error(err.Error())
	}
}

func (l *LaunchService) writeOutput(ctx context.Context, output chan []byte, launchId int) {
	for {
		res, ok := <-output
		if ok {
			err := l.launchesRepository.AddOutputToLaunch(ctx, launchId, byteconv.String(res))
			if err != nil {
				l.logger.Error(err.Error())
			}
		} else {
			break
		}
	}
}

func (l *LaunchService) Stop(ctx context.Context, commandId int) error {
	v, ok := l.runs.LoadAndDelete(commandId)
	if !ok {
		return ErrNotFound
	}
	stop, ok := v.(context.CancelFunc)
	if !ok {
		return errors.Wrapf(ErrStop, ": commandID - %d", commandId)
	}
	stop()
	return nil
}
