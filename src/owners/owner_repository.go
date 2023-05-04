package owners

import (
	"GoPetClinic/src/persistence"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "owners"

type OwnerRepository interface {
	InsertOne(ctx context.Context, owner *InsertOwner) (*Owner, error)
	List(ctx context.Context, pagination persistence.PaginatedRequest) (persistence.PaginatedResult[*Owner], error)
	FindOneById(ctx context.Context, id primitive.ObjectID) (*Owner, error)
}

type repositoryImpl struct {
	GetDatabase persistence.GetDB
}

func NewOwnerRepository(getDatabase persistence.GetDB) OwnerRepository {
	return &repositoryImpl{
		GetDatabase: getDatabase,
	}
}

func (repo *repositoryImpl) GetOwnerCollection() *mongo.Collection {
	return repo.GetDatabase().Collection(CollectionName)
}

func (repo *repositoryImpl) InsertOne(ctx context.Context, owner *InsertOwner) (*Owner, error) {
	inserted, err := repo.GetOwnerCollection().InsertOne(ctx, owner)

	if err != nil {
		return nil, err
	}

	return &Owner{
		Id: inserted.InsertedID.(primitive.ObjectID),
	}, nil
}

func (repo *repositoryImpl) List(pagination persistence.PaginatedRequest) (persistence.PaginatedResult[*Owner], error) {
	//TODO implement me
	panic("implement me")
}

func (repo *repositoryImpl) FindOneById(id primitive.ObjectID) (*Owner, error) {
	//TODO implement me
	panic("implement me")
}
