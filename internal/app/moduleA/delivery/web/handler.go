package web

import (
	"context"
	"encoding/json"
	"net/http"

	"golang-integration-test/internal/app/moduleA/usecase"
	customerrors "golang-integration-test/internal/pkg/errors"

	"github.com/go-chi/render"
)

func MakeModuleACreatorWebHandler(creator usecase.ModuleACreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cmd := usecase.ModuleACreatorCommand{}

		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			render.JSON(w, r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, http.StatusBadRequest)), err)
			return
		}

		if err := creator.Create(ctx, cmd); err != nil {
			switch err {
			case customerrors.ErrInvalidRequestParameter:
				render.JSON(w, r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, http.StatusBadRequest)), err)
			default:
				render.JSON(w, r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, http.StatusInternalServerError)), err)
			}
			return
		}

		render.Status(r, http.StatusCreated)
	}
}
