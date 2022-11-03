package resolvers

import (
	"context"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"github.com/lambda-platform/lambda/utils"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func Paginate(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, page int, size int) (*model.Paginate, error) {

	target, _, err := gql.GetPaginationTargetAndColumns(ctx)
	requestColumns, subs := gql.GetColumns(ctx, target)

	Paginate := model.Paginate{
		Page:     0,
		Total:    0,
		LastPage: 0,
	}
	if err != nil {
		return &Paginate, err
	}
	query := DB.DB
	if target == "persons" {

		requestColumns = append(requestColumns, "PERSON_ID")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.Persons{}

		TabeColumns := PersonsColumns()
		query, errorFilter := gql.Filter(filters, query, TabeColumns)
		if errorFilter != nil {
			return &Paginate, errorFilter
		}
		query, errorGroupFilter := gql.GroupFilter(groupFilters, query, TabeColumns)
		if errorGroupFilter != nil {
			return &Paginate, errorGroupFilter
		}
		query, errorOrder := gql.Order(sorts, query, TabeColumns)
		if errorOrder != nil {
			return &Paginate, errorOrder
		}

		pagination := utils.Paging(&utils.Param{
			DB:    query,
			Page:  page,
			Limit: size,
		}, &data)
		Paginate.Persons = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		if len(subs) >= 1 {
			resultWithSubs, errorsub := SetPersonsSubs(ctx, Paginate.Persons, subs, subSorts, subFilters)
			Paginate.Persons = resultWithSubs
			return &Paginate, errorsub
		} else {
			return &Paginate, nil
		}
	}
	if target == "person_items" {

		requestColumns = append(requestColumns, "ID")
		requestColumns = append(requestColumns, []string{"PERSON_ID"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.PersonItems{}

		TabeColumns := PersonItemsColumns()
		query, errorFilter := gql.Filter(filters, query, TabeColumns)
		if errorFilter != nil {
			return &Paginate, errorFilter
		}
		query, errorGroupFilter := gql.GroupFilter(groupFilters, query, TabeColumns)
		if errorGroupFilter != nil {
			return &Paginate, errorGroupFilter
		}
		query, errorOrder := gql.Order(sorts, query, TabeColumns)
		if errorOrder != nil {
			return &Paginate, errorOrder
		}

		pagination := utils.Paging(&utils.Param{
			DB:    query,
			Page:  page,
			Limit: size,
		}, &data)
		Paginate.PersonItems = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	return &Paginate, nil
}
