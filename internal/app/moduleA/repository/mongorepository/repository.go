package mongorepository

import (
	"context"
	"golang-integration-test/internal/app/ent"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModuleAMongoRepositoryOptions struct {
	DB *mongo.Database
}

type ModuleAMongoRepository struct {
	options ModuleAMongoRepositoryOptions
}

func NewModuleAMongoRepository(options ModuleAMongoRepositoryOptions) *ModuleAMongoRepository {
	return &ModuleAMongoRepository{options: options}
}

func (r *ModuleAMongoRepository) collection() *mongo.Collection {
	return r.options.DB.Collection("ModuleA")
}

func (r *ModuleAMongoRepository) Insert(ctx context.Context, entity *ent.EntityA) error {
	if _, err := r.collection().InsertOne(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (r *ModuleAMongoRepository) Get(ctx context.Context, a int) (*ent.EntityA, error) {
	var entity ent.EntityA

	filter := bson.M{"a": a}

	if err := r.collection().FindOne(ctx, filter).Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}
