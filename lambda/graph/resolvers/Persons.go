package resolvers

import (
	"context"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/graphql/gql"
	"lambda/lambda/graph/model"
	"lambda/lambda/graph/models"
)

func Persons(ctx context.Context, sorts []*model.Sort, groupFilters []*model.GroupFilter, filters []*model.Filter, subSorts []*model.SubSort, subFilters []*model.SubFilter, limit *int, offset *int) ([]*models.Persons, error) {

	result := []*models.Persons{}
	requestColumns, subs := gql.GetColumns(ctx, "PERSONS")
	requestColumns = append(requestColumns, "PERSON_ID")
	requestColumns = append(requestColumns, []string{}...)
	requestColumns = gql.RemoveDuplicate(requestColumns)
	query := DB.DB.Select(requestColumns)
	columns := PersonsColumns()
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
		resultWithSubs, errorsub := SetPersonsSubs(ctx, result, subs, subSorts, subFilters)
		return resultWithSubs, errorsub
	} else {
		return result, err
	}
}

func PersonsColumns() []string {
	return []string{"PERSON_ID", "FIRST_NAME", "LAST_NAME", "CREATED_AT", "UPDATED_AT", "DELETED_AT"}
}

var PersonsSubs = map[string]model.SubTable{
	"PERSON_ITEMS": model.SubTable{
		Table:           "PERSON_ITEMS",
		ParentIdentity:  "PERSON_ID",
		ConnectionField: "PERSON_ID",
	},
}

func PersonsSub(table string) model.SubTable {
	return PersonsSubs[table]
}

func SetPersonsSubs(ctx context.Context, parents []*models.Persons, subs []gql.Sub, subSorts []*model.SubSort, subFilters []*model.SubFilter) ([]*models.Persons, error) {
	parentIds := ""
	for _, parent := range parents {
		if parentIds == "" {
			parentIds = fmt.Sprintf("%v", parent.PersonID)
		} else {
			parentIds = parentIds + "," + fmt.Sprintf("%v", parent.PersonID)
		}
	}
	for _, sub := range subs {

		if sub.Table == "PERSON_ITEMS" {
			subItem := PersonsSub("PERSON_ITEMS")
			sorts := []*model.Sort{}
			filters := []*model.Filter{}
			for _, sort := range subSorts {
				if sort.Table == "PERSON_ITEMS" {
					newSort := model.Sort{
						Column: sort.Column,
						Order:  sort.Order,
					}
					sorts = append(sorts, &newSort)
				}
			}
			for _, filter := range subFilters {
				if filter.Table == "PERSON_ITEMS" {
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
			SubItems, err := PersonItems(ctx, sorts, []*model.GroupFilter{}, filters, nil, nil)
			if err != nil {
				return parents, err
			}
			for _, SubItemData := range SubItems {
				for i, _ := range parents {
					if fmt.Sprintf("%v", parents[i].PersonID) == fmt.Sprintf("%v", SubItemData.PersonID) {
						parents[i].PersonItems = append(parents[i].PersonItems, SubItemData)
					}
				}
			}
		}
	}

	return parents, nil
}
func CreatePersons(ctx context.Context, input model.PersonsInput) (*models.Persons, error) {
	row := models.Persons{}
	row.FirstName = input.FirstName
	row.LastName = input.LastName

	err := DB.DB.Create(&row).Error
	return &row, err
}
func UpdatePersons(ctx context.Context, id string, input model.PersonsInput) (*models.Persons, error) {
	row := models.Persons{}
	DB.DB.Where("PERSON_ID = ?", id).Find(&row)
	row.FirstName = input.FirstName
	row.LastName = input.LastName

	err := DB.DB.Save(&row).Error

	return &row, err
}
func DeletePersons(ctx context.Context, id string) (*models.Persons, error) {
	row := models.Persons{}
	DB.DB.Where("PERSON_ID = ?", id).Find(&row)
	err := DB.DB.Where("PERSON_ID = ?", id).Delete(&models.Persons{}).Error
	return &row, err
}
