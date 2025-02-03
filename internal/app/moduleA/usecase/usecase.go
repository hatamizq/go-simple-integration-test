package usecase

import (
	"context"
	"time"
)

type ModuleACreatorCommand struct {
	A time.Time `json:"a"`
}

type ModuleACreator interface {
	Create(ctx context.Context, cmd ModuleACreatorCommand) error
}
