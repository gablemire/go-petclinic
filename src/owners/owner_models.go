package owners

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	Dog = "dog"
	Cat = "cat"
)

type Pet struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	BirthDate string             `bson:"birth_date"`
	PetType   string             `bson:"pet_type"`
}

type InsertPet struct {
	Name      string `bson:"name"`
	BirthDate string `bson:"birth_date"`
	PetType   string `bson:"pet_type"`
}

type Owner struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Address   string             `bson:"address"`
	City      string             `bson:"city"`
	Telephone string             `bson:"telephone"`
	Pets      []Pet              `bson:"pets"`
}

type InsertOwner struct {
	FirstName string      `bson:"first_name"`
	LastName  string      `bson:"last_name"`
	Address   string      `bson:"address"`
	City      string      `bson:"city"`
	Telephone string      `bson:"telephone"`
	Pets      []InsertPet `bson:"pets"`
}
