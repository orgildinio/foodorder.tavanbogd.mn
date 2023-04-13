package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"github.com/lambda-platform/lambda/config"
	"github.com/lambda-platform/lambda/datagrid"
	"github.com/lambda-platform/lambda/utils"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
)

func HuraanguiTailanAggergation(c *fiber.Ctx, datagrid datagrid.Datagrid) error {

	query := DB.DB.Table("users")

	query = query.Joins("LEFT JOIN orders ON users.id = orders.user_id").Where("users.pos_name IS NOT NULL")

	query, _ = Filter(c, datagrid, query, "user_short_report")

	if len(datagrid.Condition) > 0 {
		query = query.Where(datagrid.Condition)
	}

	query = query.Select(datagrid.Aggergation)

	rows, _ := query.Rows()

	data := []interface{}{}
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	/*end*/

	for rows.Next() {

		/*start */

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		var myMap = make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)

			if ok {

				v, error := strconv.ParseInt(string(b), 10, 64)
				if error != nil {
					stringValue := string(b)

					myMap[col] = stringValue
				} else {
					myMap[col] = v
				}

			} else {
				myMap[col] = val
			}

		}
		/*end*/

		data = append(data, myMap)

	}

	return c.JSON(data)
}

func HuraanguiTailanFetchData(c *fiber.Ctx, datagrid datagrid.Datagrid) error {

	pageLimit := c.Query("paginate")
	page := c.Query("page")

	query := DB.DB.Table("users")

	query = query.Select(`users.id,
    users.company_id,
    users.company_name,
    users.pos_name,
    sum(orders.order_quantity) AS order_quantity,
    sum(orders.price) AS price,
    concat(users.last_name, ' ', users.first_name) AS full_name,
    max(orders.created_at) AS latest_create`)

	query = query.Joins("LEFT JOIN orders ON users.id = orders.user_id").Where("users.pos_name IS NOT NULL")

	query = query.Group("users.id, users.company_id, users.company_name, users.pos_name, CONCAT(users.last_name, ' ', users.first_name)")

	query = query.Order("users.last_name ASC")

	//DB.DB.LogMode(true)
	query, _ = Filter(c, datagrid, query, "user_short_report")

	var Page_ int = 1
	if page != "" {
		Page_, _ = strconv.Atoi(page)
	}
	Limit_, _ := strconv.Atoi(pageLimit)

	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  Page_,
		Limit: Limit_,
	}, datagrid.Data)

	if len(datagrid.Relations) >= 1 {
		data.Data = datagrid.FillVirtualColumns(datagrid.Data)
	}

	return c.JSON(data)

}

func HuraanguiTailanCompanyAggergation(c *fiber.Ctx, datagrid datagrid.Datagrid) error {

	query := DB.DB.Table("orders")

	query = query.Joins("LEFT JOIN users u ON u.id = orders.user_id LEFT JOIN view_order_detail_searcher d ON orders.id = d.order_id").Where("u.company_id IS NOT NULL")

	query, _ = Filter(c, datagrid, query, "user_short_report")

	if len(datagrid.Condition) > 0 {
		query = query.Where(datagrid.Condition)
	}

	query = query.Select(datagrid.Aggergation)

	rows, _ := query.Rows()

	data := []interface{}{}
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	/*end*/

	for rows.Next() {

		/*start */

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		var myMap = make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)

			if ok {

				v, error := strconv.ParseInt(string(b), 10, 64)
				if error != nil {
					stringValue := string(b)

					myMap[col] = stringValue
				} else {
					myMap[col] = v
				}

			} else {
				myMap[col] = val
			}

		}
		/*end*/

		data = append(data, myMap)

	}

	return c.JSON(data)
}

func HuraanguiTailanCompanyFetData(c *fiber.Ctx, datagrid datagrid.Datagrid) error {

	pageLimit := c.Query("paginate")
	page := c.Query("page")

	query := DB.DB.Table("orders")

	query = query.Select(`
    users.company_name,
       users.company_id,
       CASE
           WHEN tbl_order_rule.food_order_time_name IS NULL THEN 'Захиалгат хоол'::text::character varying
           ELSE tbl_order_rule.food_order_time_name
           END                      AS food_order_time_name,
       max(orders.created_at::date) AS created_at,
       sum(orders.order_quantity)   AS order_quantity,
       sum(orders.price)            AS price,
       tbl_menu.order_rule_id
`)

	query = query.Joins(" LEFT JOIN users ON users.id = orders.user_id LEFT JOIN view_order_detail_searcher d ON orders.id = d.order_id LEFT JOIN tbl_menu ON d.menu_id = tbl_menu.id LEFT JOIN tbl_order_rule ON tbl_menu.order_rule_id = tbl_order_rule.id").Where("users.company_id IS NOT NULL")

	query = query.Group("users.company_name, tbl_order_rule.food_order_time_name, users.company_id, tbl_menu.order_rule_id")

	query = query.Order("users.company_name ASC")

	//DB.DB.LogMode(true)
	query, _ = Filter(c, datagrid, query, "user_short_report")

	var Page_ int = 1
	if page != "" {
		Page_, _ = strconv.Atoi(page)
	}
	Limit_, _ := strconv.Atoi(pageLimit)

	data := utils.Paging(&utils.Param{
		DB:    query,
		Page:  Page_,
		Limit: Limit_,
	}, datagrid.Data)

	if len(datagrid.Relations) >= 1 {
		data.Data = datagrid.FillVirtualColumns(datagrid.Data)
	}

	return c.JSON(data)

}

func Filter(c *fiber.Ctx, datagrid datagrid.Datagrid, query *gorm.DB, reportType string) (*gorm.DB, string) {

	customHeader := ""
	bodyBytes := utils.GetBody(c)
	var filterData map[string]interface{}
	json.Unmarshal([]byte(bodyBytes), &filterData)

	if len(filterData) >= 1 {

		if val, ok := filterData["customHeader"]; ok {
			customHeader = val.(string)
		}

		for k, v := range filterData {
			if k == "user_condition" {

				for _, userCondition := range v.([]interface{}) {
					codintion := reflect.ValueOf(userCondition).Interface().(map[string]interface{})
					User := agentUtils.AuthUserObject(c)

					query = query.Where(codintion["grid_field"].(string)+" = ?", User[codintion["user_field"].(string)])
				}

			} else if k == "custom_condition" {

				if reflect.TypeOf(v).String() == "map[string]interface {}" {
					for kc, vc := range v.(map[string]interface{}) {
						query = query.Where(kc+" = ?", fmt.Sprintf("%v", vc))
					}
				} else {
					query = query.Where(fmt.Sprintf("%v", v))
				}

			} else {
				filterType := datagrid.Filters[k]

				if filterType != "" {
					switch filterType {
					case "Select":
						query = query.Where(k+" = ?", v)
					case "Tag":
						query = query.Where(k+" IN (?)", v)
					case "Date":
						if config.Config.Database.Connection == "oracle" {
							query = query.Where(k+" = TO_DATE(?,'YYYY-MM-DD')", fmt.Sprintf("%v", v))
						} else {
							query = query.Where(k+" = ?", fmt.Sprintf("%v", v))
						}

					case "DateRange":

						if reportType == "user_short_report" {
							query = query.Where("orders.created_at BETWEEN ? AND ?", reflect.ValueOf(v).Index(0).Interface().(string), reflect.ValueOf(v).Index(1).Interface().(string))
						} else if reportType == "food_short_report" {
							query = query.Where("orders.created_at BETWEEN ? AND ?", reflect.ValueOf(v).Index(0).Interface().(string), reflect.ValueOf(v).Index(1).Interface().(string))
						} else {
							query = query.Where(k+" BETWEEN ? AND ?", reflect.ValueOf(v).Index(0).Interface().(string), reflect.ValueOf(v).Index(1).Interface().(string))
						}

					case "DateRangeDouble":
						start := reflect.ValueOf(v).Index(0).Interface().(string)
						end := reflect.ValueOf(v).Index(1).Interface().(string)
						if start != "" && end != "" {
							query = query.Where(k+" BETWEEN ? AND ?", start, end)
						} else if start != "" && end == "" {
							query = query.Where(k+" >= ?", start)
						} else if start == "" && end != "" {
							query = query.Where(k+" <= ?", end)
						}

					default:
						switch vtype := v.(type) {
						case map[string]interface{}:
							fmt.Println(vtype)
							vmap := v.(map[string]interface{})
							switch vmap["type"] {
							case "contains":
								query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", vmap["filter"]))+"%")
							case "equals":
								query = query.Where(k+" = ?", fmt.Sprintf("%v", vmap["filter"]))
							case "lessThan":
								query = query.Where(k+" <= ?", fmt.Sprintf("%v", vmap["filter"]))
							case "greaterThan":
								query = query.Where(k+" >= ?", fmt.Sprintf("%v", vmap["filter"]))
							case "notContains":
								query = query.Where(k+" != ?", fmt.Sprintf("%v", vmap["filter"]))
							default:
								query = query.Where(k+" = ?", fmt.Sprintf("%v", vmap["filter"]))
								//query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", v))+"%")
							}
						default:
							if filterType == "Text" {

								query = query.Where("LOWER("+k+") LIKE ?", "%"+strings.ToLower(fmt.Sprintf("%v", v))+"%")

							} else {
								query = query.Where(k+" = ?", fmt.Sprintf("%v", v))

							}

						}

					}
				} else {
					query = query.Where(k+" = ?", fmt.Sprintf("%v", v))
				}
			}

		}

	}

	return query, customHeader
}
