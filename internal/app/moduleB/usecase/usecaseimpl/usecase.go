package usecaseimpl

import (
	"fmt"
	"golang-integration-test/internal/app/moduleB/usecase"
	customerrors "golang-integration-test/internal/pkg/errors"

	"github.com/google/uuid"
)

type ModuleBOptions struct{}

type ModuleB struct {
	options ModuleBOptions
}

func NewModuleB(options ModuleBOptions) *ModuleB {
	return &ModuleB{options: options}
}

func (b *ModuleB) Do(cmd usecase.ModuleBUsecaseCommand) (usecase.ModuleBDto, error) {
	if cmd.B.IsZero() {
		return usecase.ModuleBDto{}, customerrors.ErrInvalidRequestParameter
	}

	return usecase.ModuleBDto{
		B: fmt.Sprintf("%s-%d", uuid.New().String(), cmd.B.Unix()),
	}, nil
}
