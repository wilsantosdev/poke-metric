package repository

import (
	"context"
	"trainer/internal/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const (
	mongoDBURI     = "mongodb://root:example@mongo:27017/"
	databaseName   = "pokemon-prometheus"
	collectionName = "trainers"
)

type trainerMongoRepository struct {
	client *mongo.Client
}

func NewTrainerMongoRepository() *trainerMongoRepository {
	clientOptions := options.Client().ApplyURI(mongoDBURI)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(err)
	}
	return &trainerMongoRepository{
		client: client,
	}
}

func (t *trainerMongoRepository) CreateTrainer(ctx context.Context, name string, favoritePokemonType domain.PokemonType) (*domain.Trainer, error) {
	tr := otel.Tracer("trainer.repository")
	_, span := tr.Start(ctx, "MongoDB InsertOne")
	span.SetAttributes(
		attribute.String("db.system", "mongodb"),
		attribute.String("db.operation", "insert"),
		attribute.String("db.name", databaseName),
		attribute.String("db.collection", collectionName),
	)
	defer span.End()

	trainerId := uuid.New().String()
	collection := t.client.Database(databaseName).Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), bson.M{
		"id":                  trainerId,
		"name":                name,
		"favotitePokemonType": favoritePokemonType,
		"pokemons":            []domain.Pokemon{},
	})
	if err != nil {
		return nil, err
	}

	return domain.NewTrainer(trainerId, name, favoritePokemonType, []domain.Pokemon{}), nil
}

func (t *trainerMongoRepository) GetTrainer(ctx context.Context, id string) (*domain.Trainer, error) {
	tr := otel.Tracer("trainer.repository")
	_, span := tr.Start(ctx, "MongoDB FindOne")
	span.SetAttributes(
		attribute.String("db.system", "mongodb"),
		attribute.String("db.operation", "find"),
		attribute.String("db.name", databaseName),
		attribute.String("db.collection", collectionName),
	)
	defer span.End()

	collection := t.client.Database(databaseName).Collection(collectionName)

	var result struct {
		ID                  string `bson:"id"`
		Name                string `bson:"name"`
		FavoritePokemonType string `bson:"favotitePokemonType"`
		Pokemons            []struct {
			ID           int32    `bson:"id"`
			Name         string   `bson:"name"`
			PokemonTypes []string `bson:"pokemonTypes"`
		} `bson:"pokemons"`
	}
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.ID == "" {
		return nil, mongo.ErrNoDocuments
	}

	var pokemons []domain.Pokemon
	for i := range result.Pokemons {
		pokemons = append(pokemons, *domain.NewPokemon(
			result.Pokemons[i].ID,
			result.Pokemons[i].Name,
			func() []domain.PokemonType {
				var types []domain.PokemonType
				for _, t := range result.Pokemons[i].PokemonTypes {
					types = append(types, domain.PokemonType(t))
				}
				return types
			}(),
		))
	}

	trainer := domain.NewTrainer(
		result.ID,
		result.Name,
		domain.PokemonType(result.FavoritePokemonType),
		pokemons,
	)

	return trainer, nil
}

func (t *trainerMongoRepository) AddPokemon(ctx context.Context, id string, pokemon domain.Pokemon) (*domain.Trainer, error) {
	tr := otel.Tracer("trainer.repository")
	_, span := tr.Start(ctx, "MongoDB UpdateOne")
	span.SetAttributes(
		attribute.String("db.system", "mongodb"),
		attribute.String("db.operation", "update"),
		attribute.String("db.name", databaseName),
		attribute.String("db.collection", collectionName),
	)
	defer span.End()

	collection := t.client.Database(databaseName).Collection(collectionName)
	filter := bson.M{"id": id}
	update := bson.M{"$push": bson.M{"pokemons": bson.M{
		"id":           pokemon.ID(),
		"name":         pokemon.Name(),
		"pokemonTypes": pokemon.PokemonTypes(),
	},
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return t.GetTrainer(ctx, id)
}
