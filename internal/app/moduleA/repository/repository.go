package repository

import (
	"context"
	"golang-integration-test/internal/app/ent"
)

type ModuleARepository interface {
	Insert(ctx context.Context, entity *ent.EntityA) error
	Get(ctx context.Context, a int) (*ent.EntityA, error)
}
