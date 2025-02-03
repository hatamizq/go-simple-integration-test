package usecaseimpl

import (
	"context"
	"golang-integration-test/internal/app/ent"
	"golang-integration-test/internal/app/moduleA/repository"
	"golang-integration-test/internal/app/moduleA/usecase"
	modulesB "golang-integration-test/internal/app/moduleB/usecase"
)

type ModuleACreatorOptions struct {
	Repository repository.ModuleARepository
	ModuleBUC  modulesB.ModuleBUsecase
}

type ModuleACreator struct {
	options ModuleACreatorOptions
}

func NewModuleACreator(options ModuleACreatorOptions) *ModuleACreator {
	return &ModuleACreator{options: options}
}

func (s *ModuleACreator) Create(ctx context.Context, cmd usecase.ModuleACreatorCommand) error {
	BDto, err := s.options.ModuleBUC.Do(modulesB.ModuleBUsecaseCommand{B: cmd.A})
	if err != nil {
		return err
	}

	return s.options.Repository.Insert(ctx, &ent.EntityA{
		A: cmd.A.Unix(),
		B: BDto.B,
	})
}
