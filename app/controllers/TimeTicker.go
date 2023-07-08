package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
	"time"
)

func TimeTick() {
	var rules []models.TblOrderRule
	var carts []models.CartMenu
	var cartSets []models.CartSet
	var orders []models.Orders

	//ticker := time.NewTicker(1 * time.Second)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-ticker.C:
	now := time.Now()

	afterTenMinutes := now.Add(15 * time.Minute)
	nowTime := afterTenMinutes.Format("15:04:05")

	minusThirtyMinute := now.Add(-30 * time.Minute)
	formattedTimeStamp := minusThirtyMinute.Format("2006-01-02 15:04:05")

	DB.DB.Order("morning_order_end ASC").Find(&rules)
	DB.DB.Order("created_at DESC").Find(&carts)
	DB.DB.Order("created_at DESC").Find(&cartSets)

	//Delete from cart menu
	for _, cart := range carts {
		formattedCreatedAt := cart.CreatedAt.Format("2006-01-02 15:04:05")

		if formattedTimeStamp == formattedCreatedAt {
			DB.DB.Where("id = ?", cart.ID).Delete(&carts)
		}
	}

	//Delete from cart set
	for _, cartSet := range cartSets {
		formattedCreatedAt := cartSet.CreatedAt.Format("2006-01-02 15:04:05")

		if formattedTimeStamp >= formattedCreatedAt {
			DB.DB.Debug().Where("id = ?", cartSet.ID).Delete(&cartSets)
		}
	}

	// Get the current time in the desired time zone

	//Update orders payment status
	minusFiveMinute := now.Add(-5 * time.Minute)
	formattedMinusTimeStamp := minusFiveMinute.Format("2006-01-02 15:04:05")

	DB.DB.Where("TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') >= ? AND payment_status = 'pending'", formattedMinusTimeStamp).Find(&orders)
	for _, order := range orders {
		formattedCreatedAt := order.CreatedAt.Format("2006-01-02 15:04:05")

		if formattedMinusTimeStamp >= formattedCreatedAt {
			DB.DB.Model(&orders).Where("id = ? AND payment_status = 'pending'", order.ID).Update("payment_status", "canceled")
		}
	}

	//Notify
	for _, rule := range rules {

		currentTime := formatDatetime(time.Now())
		ruleEndTime := getTodayStartDatetime(rule.MorningOrderEnd)
		remainingMinutes := ruleEndTime.Sub(currentTime).Minutes()

		if int(remainingMinutes) == 15 {
			if nowTime == rule.MorningOrderEnd {
				menu := models.TblMenu{}
				DB.DB.Where("order_rule_id = ? AND set_date = ?", rule.ID, now.Format("2006-01-02")).Find(&menu)
				LeftTimeSend(menu.ID)
			}
		}

	}
	//}
	//}
	//}()
}

func formatDatetime(dateTime time.Time) time.Time {
	layout := "2006-01-02 15:04:05"
	formattedTime := dateTime.Format(layout)
	formatedTime, _ := time.Parse(layout, formattedTime)

	return formatedTime
}
func getTodayStartDatetime(startTimeString string) time.Time {
	layout := "2006-01-02"

	formattedTime := time.Now().Format(layout)
	formattedTime = formattedTime + " " + startTimeString
	layout2 := "2006-01-02 15:04:05"
	startTime, _ := time.Parse(layout2, formattedTime)

	return formatDatetime(startTime)
}

func OrderLeftTimeTicker(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	orderNumber := c.Params("order_number")

	var orders models.OrdersTimeDiff
	DB.DB.Debug().Where("payment_status = 'pending' AND order_number = ? AND user_id = ?", orderNumber, orderUser["id"]).Find(&orders)

	if orders.ID > 0 {

		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "success",
			"message": "Found orders",
			"time":    orders.Diff,
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "warning",
		"message": "Not found active order",
	})
}
