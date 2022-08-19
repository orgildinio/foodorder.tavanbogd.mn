package resolvers

import (
	"context"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func ViewOronNutagZovlol(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewOronNutagZovlol, error) {

	result := []*models.ViewOronNutagZovlol{}
	requestColumns, subs := gql.GetColumns(ctx, "view_oron_nutag_zovlol")
	requestColumns = append(requestColumns, "id")
	requestColumns = append(requestColumns, []string{}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := ViewOronNutagZovlolColumns()
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
		resultWithSubs, errorsub := SetViewOronNutagZovlolSubs(ctx, result, subs, subSorts, subFilters)
		return resultWithSubs, errorsub
	} else {
		return result, err
	}
}

func ViewOronNutagZovlolColumns() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "aimag_id", "sum_id", "aimagname", "sumname"}
}

var ViewOronNutagZovlolSubs = map[string]model.SubTable{
	"sub_oron_nutag_gamshig_ersdel": model.SubTable{
		Table:           "sub_oron_nutag_gamshig_ersdel",
		ParentIdentity:  "id",
		ConnectionField: "o_zovlol_id",
	},
	"sub_oron_nutag_togtool": model.SubTable{
		Table:           "sub_oron_nutag_togtool",
		ParentIdentity:  "id",
		ConnectionField: "o_zovlol_id",
	},
	"sub_oron_nutag_ua_heregjilt": model.SubTable{
		Table:           "sub_oron_nutag_ua_heregjilt",
		ParentIdentity:  "id",
		ConnectionField: "o_zovlol_id",
	},
	"sub_oron_nutag_ua_tolvolgoo": model.SubTable{
		Table:           "sub_oron_nutag_ua_tolvolgoo",
		ParentIdentity:  "id",
		ConnectionField: "o_zovlol_id",
	},
}

func ViewOronNutagZovlolSub(table string) model.SubTable {
	return ViewOronNutagZovlolSubs[table]
}

func SetViewOronNutagZovlolSubs(ctx context.Context, parents []*models.ViewOronNutagZovlol, subs []gql.Sub, subSorts []*model.SubSort, subFilters []*model.SubFilter) ([]*models.ViewOronNutagZovlol, error) {
	parentIds := ""
	for _, parent := range parents {
		if parentIds == "" {
			parentIds = fmt.Sprintf("%v", *parent.ID)
		} else {
			parentIds = parentIds + "," + fmt.Sprintf("%v", *parent.ID)
		}
	}
	for _, sub := range subs {

		if sub.Table == "sub_oron_nutag_gamshig_ersdel" {
			subItem := ViewOronNutagZovlolSub("sub_oron_nutag_gamshig_ersdel")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_oron_nutag_gamshig_ersdel" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_oron_nutag_gamshig_ersdel" {
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
			SubItems, err := SubOronNutagGamshigErsdel(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.OZovlolID) {
						parents[i].SubOronNutagGamshigErsdel = append(parents[i].SubOronNutagGamshigErsdel, SubItemData)
					}
				}
			}
		}
		if sub.Table == "sub_oron_nutag_togtool" {
			subItem := ViewOronNutagZovlolSub("sub_oron_nutag_togtool")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_oron_nutag_togtool" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_oron_nutag_togtool" {
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
			SubItems, err := SubOronNutagTogtool(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.OZovlolID) {
						parents[i].SubOronNutagTogtool = append(parents[i].SubOronNutagTogtool, SubItemData)
					}
				}
			}
		}
		if sub.Table == "sub_oron_nutag_ua_heregjilt" {
			subItem := ViewOronNutagZovlolSub("sub_oron_nutag_ua_heregjilt")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_oron_nutag_ua_heregjilt" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_oron_nutag_ua_heregjilt" {
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
			SubItems, err := SubOronNutagUaHeregjilt(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.OZovlolID) {
						parents[i].SubOronNutagUaHeregjilt = append(parents[i].SubOronNutagUaHeregjilt, SubItemData)
					}
				}
			}
		}
		if sub.Table == "sub_oron_nutag_ua_tolvolgoo" {
			subItem := ViewOronNutagZovlolSub("sub_oron_nutag_ua_tolvolgoo")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_oron_nutag_ua_tolvolgoo" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_oron_nutag_ua_tolvolgoo" {
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
			SubItems, err := SubOronNutagUaTolvolgoo(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.OZovlolID) {
						parents[i].SubOronNutagUaTolvolgoo = append(parents[i].SubOronNutagUaTolvolgoo, SubItemData)
					}
				}
			}
		}
	}

	return parents, nil
}
