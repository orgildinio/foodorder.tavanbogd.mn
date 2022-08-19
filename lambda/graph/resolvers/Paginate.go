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
	if target == "lut_document_sorting" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.LutDocumentSorting{}

		TabeColumns := LutDocumentSortingColumns()
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
		Paginate.LutDocumentSorting = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "view_ez_bodlogiin_bb" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.ViewEzBodlogiinBb{}

		TabeColumns := ViewEzBodlogiinBbColumns()
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
		Paginate.ViewEzBodlogiinBb = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "hamtragch_baiguullaga" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.HamtragchBaiguullaga{}

		TabeColumns := HamtragchBaiguullagaColumns()
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
		Paginate.HamtragchBaiguullaga = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_oron_nutag_gamshig_ersdel" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"o_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubOronNutagGamshigErsdel{}

		TabeColumns := SubOronNutagGamshigErsdelColumns()
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
		Paginate.SubOronNutagGamshigErsdel = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_oron_nutag_togtool" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"o_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubOronNutagTogtool{}

		TabeColumns := SubOronNutagTogtoolColumns()
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
		Paginate.SubOronNutagTogtool = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_oron_nutag_ua_heregjilt" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"o_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubOronNutagUaHeregjilt{}

		TabeColumns := SubOronNutagUaHeregjiltColumns()
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
		Paginate.SubOronNutagUaHeregjilt = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_oron_nutag_ua_tolvolgoo" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"o_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubOronNutagUaTolvolgoo{}

		TabeColumns := SubOronNutagUaTolvolgooColumns()
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
		Paginate.SubOronNutagUaTolvolgoo = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "view_oron_nutag_zovlol" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.ViewOronNutagZovlol{}

		TabeColumns := ViewOronNutagZovlolColumns()
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
		Paginate.ViewOronNutagZovlol = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		if len(subs) >= 1 {
			resultWithSubs, errorsub := SetViewOronNutagZovlolSubs(ctx, Paginate.ViewOronNutagZovlol, subs, subSorts, subFilters)
			Paginate.ViewOronNutagZovlol = resultWithSubs
			return &Paginate, errorsub
		} else {
			return &Paginate, nil
		}
	}
	if target == "sub_undes_zovlol_huraldaan" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"undes_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubUndesZovlolHuraldaan{}

		TabeColumns := SubUndesZovlolHuraldaanColumns()
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
		Paginate.SubUndesZovlolHuraldaan = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_undes_zovlol_gamshig_ersdel_buuruulah" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"undes_zovlol_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubUndesZovlolGamshigErsdelBuuruulah{}

		TabeColumns := SubUndesZovlolGamshigErsdelBuuruulahColumns()
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
		Paginate.SubUndesZovlolGamshigErsdelBuuruulah = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "view_undesnii_zovlol" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.ViewUndesniiZovlol{}

		TabeColumns := ViewUndesniiZovlolColumns()
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
		Paginate.ViewUndesniiZovlol = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		if len(subs) >= 1 {
			resultWithSubs, errorsub := SetViewUndesniiZovlolSubs(ctx, Paginate.ViewUndesniiZovlol, subs, subSorts, subFilters)
			Paginate.ViewUndesniiZovlol = resultWithSubs
			return &Paginate, errorsub
		} else {
			return &Paginate, nil
		}
	}
	if target == "lut_angi" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.LutAngi{}

		TabeColumns := LutAngiColumns()
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
		Paginate.LutAngi = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "sub_news_social_typies" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{"news_id"}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.SubNewsSocialTypies{}

		TabeColumns := SubNewsSocialTypiesColumns()
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
		Paginate.SubNewsSocialTypies = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	if target == "view_news" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.ViewNews{}

		TabeColumns := ViewNewsColumns()
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
		Paginate.ViewNews = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		if len(subs) >= 1 {
			resultWithSubs, errorsub := SetViewNewsSubs(ctx, Paginate.ViewNews, subs, subSorts, subFilters)
			Paginate.ViewNews = resultWithSubs
			return &Paginate, errorsub
		} else {
			return &Paginate, nil
		}
	}
	if target == "lut_news_type" {

		requestColumns = append(requestColumns, "id")
		requestColumns = append(requestColumns, []string{}...)
		requestColumns = utils.RemoveDuplicateStr(requestColumns)
		query = query.Select(requestColumns)
		data := []*models.LutNewsType{}

		TabeColumns := LutNewsTypeColumns()
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
		Paginate.LutNewsType = data
		Paginate.LastPage = pagination.LastPage
		Paginate.Total = int(pagination.Total)
		return &Paginate, nil
	}
	return &Paginate, nil
}
