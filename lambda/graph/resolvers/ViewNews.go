package resolvers

import (
	"context"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func ViewNews(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewNews, error) {

	result := []*models.ViewNews{}
	requestColumns, subs := gql.GetColumns(ctx, "view_news")
	requestColumns = append(requestColumns, "id")
	requestColumns = append(requestColumns, []string{}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := ViewNewsColumns()
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

	if len(subs) >= 1 {
		resultWithSubs, errorsub := SetViewNewsSubs(ctx, result, subs, subSorts, subFilters)
		return resultWithSubs, errorsub
	} else {
		return result, err
	}
}

func ViewNewsColumns() []string {
	return []string{"id", "news_type_id", "garchig", "delgerengui", "zurag", "ursah_eseh", "created_at", "updated_at", "deleted_at", "handalt", "logo", "share", "news_type"}
}

var ViewNewsSubs = map[string]model.SubTable{
	"sub_news_social_typies": model.SubTable{
		Table:           "sub_news_social_typies",
		ParentIdentity:  "id",
		ConnectionField: "news_id",
	},
}

func ViewNewsSub(table string) model.SubTable {
	return ViewNewsSubs[table]
}

func SetViewNewsSubs(ctx context.Context, parents []*models.ViewNews, subs []gql.Sub, subSorts []*model.SubSort, subFilters []*model.SubFilter) ([]*models.ViewNews, error) {
	parentIds := ""
	for _, parent := range parents {
		if parentIds == "" {
			parentIds = fmt.Sprintf("%v", *parent.ID)
		} else {
			parentIds = parentIds + "," + fmt.Sprintf("%v", *parent.ID)
		}
	}
	for _, sub := range subs {

		if sub.Table == "sub_news_social_typies" {
			subItem := ViewNewsSub("sub_news_social_typies")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_news_social_typies" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_news_social_typies" {
					newFilter := model.Filter{
						Column:    filter.Column,
						Condition: filter.Condition,
						Value:     filter.Value,
					}
					filters = append(filters, &newFilter)
				}
			}
			parentFilter := model.Filter{}

			parentFilter.Condition = "whereIn"
			parentFilter.Column = subItem.ConnectionField
			parentFilter.Value = parentIds
			filters = append(filters, &parentFilter)

			sub.Columns = append(sub.Columns, subItem.ConnectionField)
			SubItems, err := SubNewsSocialTypies(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.NewsID) {
						parents[i].SubNewsSocialTypies = append(parents[i].SubNewsSocialTypies, SubItemData)
					}
				}
			}
		}
	}

	return parents, nil
}
