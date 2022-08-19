package resolvers

import (
	"context"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func ViewEzBodlogiinBb(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, limit *int, offset *int) ([]*models.ViewEzBodlogiinBb, error) {

	result := []*models.ViewEzBodlogiinBb{}
	requestColumns, _ := gql.GetColumns(ctx, "view_ez_bodlogiin_bb")
	requestColumns = append(requestColumns, "id")
	requestColumns = append(requestColumns, []string{}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := ViewEzBodlogiinBbColumns()
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

func ViewEzBodlogiinBbColumns() []string {
	return []string{"id", "document_sorting_id", "sub_document_sorting_id", "ez_barimt_bichgiin_ner", "hawsralt", "created_at", "updated_at", "deleted_at", "barimt_ner_en", "file_en", "document_sorting", "document_sorting_en"}
}
