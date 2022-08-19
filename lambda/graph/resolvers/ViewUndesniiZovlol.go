package resolvers

import (
	"context"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func ViewUndesniiZovlol(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.ViewUndesniiZovlol, error) {

	result := []*models.ViewUndesniiZovlol{}
	requestColumns, subs := gql.GetColumns(ctx, "view_undesnii_zovlol")
	requestColumns = append(requestColumns, "id")
	requestColumns = append(requestColumns, []string{}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := ViewUndesniiZovlolColumns()
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
		resultWithSubs, errorsub := SetViewUndesniiZovlolSubs(ctx, result, subs, subSorts, subFilters)
		return resultWithSubs, errorsub
	} else {
		return result, err
	}
}

func ViewUndesniiZovlolColumns() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "huraldaan", "tailbar_mn_bureldehuun"}
}

var ViewUndesniiZovlolSubs = map[string]model.SubTable{
	"sub_undes_zovlol_gamshig_ersdel_buuruulah": model.SubTable{
		Table:           "sub_undes_zovlol_gamshig_ersdel_buuruulah",
		ParentIdentity:  "id",
		ConnectionField: "undes_zovlol_id",
	},
	"sub_undes_zovlol_huraldaan": model.SubTable{
		Table:           "sub_undes_zovlol_huraldaan",
		ParentIdentity:  "id",
		ConnectionField: "undes_zovlol_id",
	},
}

func ViewUndesniiZovlolSub(table string) model.SubTable {
	return ViewUndesniiZovlolSubs[table]
}

func SetViewUndesniiZovlolSubs(ctx context.Context, parents []*models.ViewUndesniiZovlol, subs []gql.Sub, subSorts []*model.SubSort, subFilters []*model.SubFilter) ([]*models.ViewUndesniiZovlol, error) {
	parentIds := ""
	for _, parent := range parents {
		if parentIds == "" {
			parentIds = fmt.Sprintf("%v", *parent.ID)
		} else {
			parentIds = parentIds + "," + fmt.Sprintf("%v", *parent.ID)
		}
	}
	for _, sub := range subs {

		if sub.Table == "sub_undes_zovlol_gamshig_ersdel_buuruulah" {
			subItem := ViewUndesniiZovlolSub("sub_undes_zovlol_gamshig_ersdel_buuruulah")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_undes_zovlol_gamshig_ersdel_buuruulah" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_undes_zovlol_gamshig_ersdel_buuruulah" {
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
			SubItems, err := SubUndesZovlolGamshigErsdelBuuruulah(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.UndesZovlolID) {
						parents[i].SubUndesZovlolGamshigErsdelBuuruulah = append(parents[i].SubUndesZovlolGamshigErsdelBuuruulah, SubItemData)
					}
				}
			}
		}
		if sub.Table == "sub_undes_zovlol_huraldaan" {
			subItem := ViewUndesniiZovlolSub("sub_undes_zovlol_huraldaan")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "sub_undes_zovlol_huraldaan" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "sub_undes_zovlol_huraldaan" {
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
			SubItems, err := SubUndesZovlolHuraldaan(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", *parents[i].ID) == fmt.Sprintf("%v", *SubItemData.UndesZovlolID) {
						parents[i].SubUndesZovlolHuraldaan = append(parents[i].SubUndesZovlolHuraldaan, SubItemData)
					}
				}
			}
		}
	}

	return parents, nil
}
