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

func (r *queryResolver) LutDocumentSorting(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.LutDocumentSorting, error) {
	return resolvers.LutDocumentSorting(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) ViewEzBodlogiinBb(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.ViewEzBodlogiinBb, error) {
	return resolvers.ViewEzBodlogiinBb(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) HamtragchBaiguullaga(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.HamtragchBaiguullaga, error) {
	return resolvers.HamtragchBaiguullaga(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubOronNutagGamshigErsdel(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubOronNutagGamshigErsdel, error) {
	return resolvers.SubOronNutagGamshigErsdel(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubOronNutagTogtool(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubOronNutagTogtool, error) {
	return resolvers.SubOronNutagTogtool(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubOronNutagUaHeregjilt(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubOronNutagUaHeregjilt, error) {
	return resolvers.SubOronNutagUaHeregjilt(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubOronNutagUaTolvolgoo(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubOronNutagUaTolvolgoo, error) {
	return resolvers.SubOronNutagUaTolvolgoo(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) ViewOronNutagZovlol(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewOronNutagZovlol, error) {
	return resolvers.ViewOronNutagZovlol(ctx, sorts, groupFilters, filters, subSorts, subFilters, limit, offset)
}

func (r *queryResolver) SubUndesZovlolHuraldaan(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubUndesZovlolHuraldaan, error) {
	return resolvers.SubUndesZovlolHuraldaan(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubUndesZovlolGamshigErsdelBuuruulah(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubUndesZovlolGamshigErsdelBuuruulah, error) {
	return resolvers.SubUndesZovlolGamshigErsdelBuuruulah(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) ViewUndesniiZovlol(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewUndesniiZovlol, error) {
	return resolvers.ViewUndesniiZovlol(ctx, sorts, groupFilters, filters, subSorts, subFilters, limit, offset)
}

func (r *queryResolver) LutAngi(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.LutAngi, error) {
	return resolvers.LutAngi(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) SubNewsSocialTypies(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubNewsSocialTypies, error) {
	return resolvers.SubNewsSocialTypies(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) ViewNews(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewNews, error) {
	return resolvers.ViewNews(ctx, sorts, groupFilters, filters, subSorts, subFilters, limit, offset)
}

func (r *queryResolver) LutNewsType(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.LutNewsType, error) {
	return resolvers.LutNewsType(ctx, sorts, groupFilters, filters, limit, offset)
}

func (r *queryResolver) Paginate(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, page int, size int) (*model.Paginate, error) {
	return resolvers.Paginate(ctx, sorts, groupFilters, filters, subSorts, subFilters, page, size)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
