package controllers

import (
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"time"
)

func TimeTick(orderID int) {
	var rules []models.TblOrderRule
	var carts []models.CartMenu
	var cartSets []models.CartSet
	var orders []models.Orders

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				now := time.Now()

				afterTenMinutes := now.Add(10 * time.Minute)
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

					if formattedTimeStamp == formattedCreatedAt {
						DB.DB.Debug().Where("id = ?", cartSet.ID).Delete(&cartSets)
					}
				}

				// Get the current time in the desired time zone

				//Update orders payment status
				minusFiveMinute := now.Add(-5 * time.Minute)
				formattedMinusTimeStamp := minusFiveMinute.Format("2006-01-02 15:04:05")

				DB.DB.Where("TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') >= ?", formattedMinusTimeStamp).Find(&orders)
				fmt.Println(formattedMinusTimeStamp)

				for _, order := range orders {
					fmt.Println("Something there!!!")
					formattedCreatedAt := order.CreatedAt.Format("2006-01-02 15:04:05")

					if formattedMinusTimeStamp == formattedCreatedAt {
						fmt.Println("Right here!!!")
						DB.DB.Debug().Model(&orders).Where("id = ?", order.ID).Update("payment_status", "canceled")
					}
				}

				//Notify
				for _, rule := range rules {

					if nowTime == rule.MorningOrderEnd {
						menu := models.TblMenu{}
						DB.DB.Where("order_rule_id = ? AND set_date = ?", rule.ID, now.Format("2006-01-02")).Find(&menu)

						LeftTimeSend(menu.ID)
					}
				}
			}
		}
	}()
}
