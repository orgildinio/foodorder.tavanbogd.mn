package resolvers

import (
	"context"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func SubOronNutagGamshigErsdel(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.SubOronNutagGamshigErsdel, error) {

	result := []*models.SubOronNutagGamshigErsdel{}
	requestColumns, _ := gql.GetColumns(ctx, "sub_oron_nutag_gamshig_ersdel")
	requestColumns = append(requestColumns, "id")
	requestColumns = append(requestColumns, []string{"o_zovlol_id"}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := SubOronNutagGamshigErsdelColumns()
	query, errorFilter := gql.Filter(filters, query, columns)
	query, errorGroupFilter := gql.GroupFilter(groupFilters, query, columns)
	if errorFilter != nil {
		return result, errorFilter
	}
	if errorGroupFilter != nil {
		return result, errorGroupFilter
	}
	query, errorOrder := gql.Order(sorts, query, columns)
	if errorOrder != nil {
		return result, errorOrder
	}
	if limit != nil {
		if *limit >= 1 {
			query = query.Limit(*limit)
		}
	}
	if offset != nil {
		if *offset >= 1 {
			query = query.Offset(*offset)
		}
	}

	err := query.Find(&result).Error

	return result, err
}

func SubOronNutagGamshigErsdelColumns() []string {
	return []string{"id", "o_zovlol_id", "tosov", "ognoo"}
}
