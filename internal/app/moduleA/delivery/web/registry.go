package web

import (
	"golang-integration-test/internal/app/moduleA/usecase"

	"github.com/go-chi/chi"
)

type ModuleAWebRegistryOptions struct {
	Creator usecase.ModuleACreator
}

type ModuleAWebRegistry struct {
	options ModuleAWebRegistryOptions
}

func NewModuleAWebRegistry(options ModuleAWebRegistryOptions) *ModuleAWebRegistry {
	return &ModuleAWebRegistry{options: options}
}

func (w *ModuleAWebRegistry) RegisterRoutesTo(router chi.Router) error {
	router.Route("/a", func(r chi.Router) {
		r.Post("/create/v1", MakeModuleACreatorWebHandler(w.options.Creator))
	})
	return nil
}
