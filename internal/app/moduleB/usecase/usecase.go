package usecase

import "time"

type ModuleBUsecaseCommand struct {
	B time.Time
}

type ModuleBUsecase interface {
	Do(cmd ModuleBUsecaseCommand) (ModuleBDto, error)
}
