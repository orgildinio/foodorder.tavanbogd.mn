package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"lambda/lambda/graph/generated"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
	"lambda/lambda/graph/resolvers"
)

func (r *mutationResolver) CreatePersons(ctx context.Context, input model.PersonsInput) (*models.Persons, error) {
	row, err := resolvers.CreatePersons(ctx, input)

	return row, err
}

func (r *mutationResolver) UpdatePersons(ctx context.Context, id string, input model.PersonsInput) (*models.Persons, error) {
	row, err := resolvers.UpdatePersons(ctx, id, input)

	return row, err
}

func (r *mutationResolver) DeletePersons(ctx context.Context, id string) (*models.Persons, error) {
	row, err := resolvers.DeletePersons(ctx, id)

	return row, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
