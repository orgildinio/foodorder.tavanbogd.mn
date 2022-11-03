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

func (r *queryResolver) Persons(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.Persons, error) {
	return resolvers.Persons(ctx, sorts, groupFilters, filters, subSorts, subFilters, limit, offset)
}

func (r *queryResolver) PersonItems(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.PersonItems, error) {
	return resolvers.PersonItems(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) Paginate(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, page int, size int) (*model.Paginate, error) {
	return resolvers.Paginate(ctx, sorts, groupFilters, filters, subSorts, subFilters, page, size)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
