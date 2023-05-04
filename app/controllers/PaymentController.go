package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/ebarimt/bill"
	"github.com/lambda-platform/ebarimt/posapi"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"lambda/ebarimt"
	"log"
	"net/http"
	"strconv"
)

func QPayInvoice(c *fiber.Ctx) error {
	qpayRequest := &models.QpayRequest{}
	err := c.BodyParser(&qpayRequest)
	request, _ := http.NewRequest("POST", "https://merchant.qpay.mn/v2/auth/token", nil)
	request.Header.Set("Authorization", "Basic Rk9PRE9SREVSOnBCeDVYdVZU")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP Token Request Failed with Error %s\n", err)
	}

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)

	accessToken := fmt.Sprint(res["access_token"])
	bearerToken := "Bearer " + accessToken

	jsonData := map[string]string{
		"invoice_code":          "FOODORDER_INVOICE",
		"sender_invoice_no":     qpayRequest.SenderInvoiceNo,
		"invoice_receiver_code": qpayRequest.InvoiceReceiverCode,
		"invoice_description":   qpayRequest.InvoiceDescription,
		"amount":                qpayRequest.Amount,
		//"callback_url":          "https://mmk.khankhulgun.mn/api/qpay/callback/" + qpayRequest.SenderInvoiceNo,
		"callback_url": "https://foodorder.tavanbogd.mn/api/qpay/callback/" + qpayRequest.SenderInvoiceNo,
	}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ = http.NewRequest("POST", "https://merchant.qpay.mn/v2/invoice", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", bearerToken)
	response, err = client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP Invoice Request Failed with Error %s\n", err)
	}

	json.NewDecoder(response.Body).Decode(&res)

	updateOrder := models.Orders{}
	DB.DB.Debug().Where("order_number = ?", qpayRequest.SenderInvoiceNo).First(&updateOrder)
	updateOrder.InvoiceID = fmt.Sprintf("%s", res["invoice_id"])
	DB.DB.Omit("is_delivery").Save(&updateOrder)

	return c.JSON(res)
}

func QPayPaymentCheck(c *fiber.Ctx) error {
	qpayCheck := &models.QpayCheck{}
	err := c.BodyParser(&qpayCheck)
	request, _ := http.NewRequest("POST", "https://merchant.qpay.mn/v2/auth/token", nil)
	request.Header.Set("Authorization", "Basic Rk9PRE9SREVSOnBCeDVYdVZU")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP Token Request Failed with Error %s\n", err)
	}

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)

	accessToken := fmt.Sprint(res["access_token"])
	bearerToken := "Bearer " + accessToken

	jsonData := map[string]string{
		"object_type": qpayCheck.ObjectType,
		"object_id":   qpayCheck.ObjectId,
	}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ = http.NewRequest("POST", "https://merchant.qpay.mn/v2/payment/check", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", bearerToken)
	response, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP Invoice Request Failed with Error %s\n", err)
	}

	qpayResponse := models.QpayResonpse{}
	json.NewDecoder(response.Body).Decode(&qpayResponse)

	if qpayResponse.Count >= 1 {
		//UpdateStatus(qpayCheck.OrderNumber, "qPay", qpayCheck.EbarimtType, &qpayCheck.OrgRegisterNumber)
	}
	return c.JSON(qpayResponse)
}

func QpayCallBackCheck(invoiceId string) int {
	request, _ := http.NewRequest("POST", "https://merchant.qpay.mn/v2/auth/token", nil)
	request.Header.Set("Authorization", "Basic Rk9PRE9SREVSOnBCeDVYdVZU")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP Token Request Failed with Error %s\n", err)
	}

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)

	accessToken := fmt.Sprint(res["access_token"])
	bearerToken := "Bearer " + accessToken

	jsonData := map[string]string{
		"object_type": "INVOICE",
		"object_id":   invoiceId,
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, _ = http.NewRequest("POST", "https://merchant.qpay.mn/v2/payment/check", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", bearerToken)
	response, err = client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP Invoice Request Failed with Error %s\n", err)
	}

	qpayResponse := models.QpayResonpse{}
	json.NewDecoder(response.Body).Decode(&qpayResponse)

	return qpayResponse.Count
}

func QPayCallBack(c *fiber.Ctx) error {
	//orderUser := agentUtils.AuthUserObject(c)
	//order := models.ViewOrder{}
	var orderNumber = c.Params("invoice_id")

	checkOrder := models.Orders{}
	DB.DB.Where("order_number = ?", orderNumber).First(&checkOrder)

	if QpayCallBackCheck(checkOrder.InvoiceID) > 0 {

		var orderDetails []models.OrderDetail
		DB.DB.Where("order_id = ?", checkOrder.ID).Find(&orderDetails)

		var orderDetailSets []models.OrderDetailSet
		DB.DB.Where("order_id = ?", checkOrder.ID).Find(&orderDetailSets)

		UpdateStatus(orderNumber, checkOrder.ID, "qpay", "success")

		for _, orderDetail := range orderDetails {
			UpdateBalance(orderDetail.FoodID, orderDetail.KitchenID, orderDetail.Qty)
		}

		for _, orderDetailSet := range orderDetailSets {
			UpdateBalance(orderDetailSet.FoodID, orderDetailSet.KitchenID, orderDetailSet.Quantity)
		}

		//go CreateEbarimt(order)

		return c.Status(http.StatusOK).JSON("SUCCESS")

	}

	return c.Status(http.StatusOK).JSON("FAILED")
}

func LaterPay(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	orderStatus := models.OrdersStatus{}
	order := models.ViewOrder{}

	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Order("id DESC").Find(&order)
	DB.DB.Where("id = ?", order.ID).Find(&orderStatus)

	if order.ID == 0 {
		return c.Status(http.StatusOK).JSON("Not found active order")
	}

	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&order)

	orderLaterPay := models.OrderLaterPay{}

	orderLaterPay.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	orderLaterPay.OrderNumber = order.OrderNumber
	orderLaterPay.OrderID = order.ID
	orderLaterPay.Qty = order.OrderQuantity
	orderLaterPay.Price = order.Price
	orderLaterPay.PaymentStatus = GetStringPointer("success")

	oldOrders := models.Orders{}
	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Find(&oldOrders)

	if oldOrders.ID >= 1 {

		//ChekingBalance(c, oldOrders, order)

		var orderDetails []models.OrderDetail
		DB.DB.Where("order_id = ? AND user_id = ?", oldOrders.ID, orderUser["id"]).Find(&orderDetails)

		var orderDetailSets []models.OrderDetailSet
		DB.DB.Where("order_id = ?", oldOrders.ID).Find(&orderDetailSets)

		for _, orderDetailSet := range orderDetailSets {

			viewBalance := models.ViewFoodBalance{}
			DB.DB.Where("food_id = ? AND kitchen_id = ?", orderDetailSet.FoodID, orderDetailSet.KitchenID).Find(&viewBalance)

			if viewBalance.Quantity < orderDetailSet.Quantity {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"status":  "warning",
					"message": "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
				})
			}

			UpdateStatus(order.OrderNumber, oldOrders.ID, "mmk", "success")
			UpdateBalance(orderDetailSet.FoodID, orderDetailSet.KitchenID, orderDetailSet.Quantity)
		}

		for _, orderDetail := range orderDetails {
			UpdateBalance(orderDetail.FoodID, orderDetail.KitchenID, orderDetail.Qty)
		}

	} else {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Идэвхтэй захиалга олдсонгүй",
		})
	}

	DB.DB.Create(&orderLaterPay)

	go CreateEbarimt(order)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": orderLaterPay.OrderNumber + " дугаартай захиалга амжилттай",
	})
}

func CreateEbarimt(order models.ViewOrder) {
	// create Bill
	bilInput := posapi.PutInput{}

	amount := bill.FormatNumber(float64(order.Price))
	bilInput.Amount = amount
	bilInput.Vat = bill.GetVat(float64(order.Price), true, false)
	bilInput.BillIDSuffix = bill.GenerateBillIdSuffix()
	bilInput.CityTax = "0.00"
	bilInput.BillType = "1"
	bilInput.CashAmount = "0.00"
	bilInput.NonCashAmount = amount
	bilInput.DistrictCode = "23"
	bilInput.PosNo = "0001"
	bilInput.BranchNo = "001"

	oEbarimt := models.OrderEbarimt{}
	DB.DB.Where("order_id = ?", order.ID).Find(&oEbarimt)

	//if *oEbarimt.EbarimtType == 1 {
	//	bilInput.BillType = "1"
	//} else {
	//	bilInput.BillType = "3"
	//	CustomerNo := strconv.Itoa(*oEbarimt.OrgRegisterNumber)
	//	bilInput.CustomerNo = CustomerNo
	//}

	var items []posapi.Stock

	var orderItems []models.ViewOrderDetail

	DB.DB.Where("order_id = ?", order.ID).Find(&orderItems)

	for _, orderPayment := range orderItems {
		var item posapi.Stock
		Code := strconv.Itoa(orderPayment.OrderID)
		Price := bill.FormatNumber(float64(orderPayment.Price))
		TotalAmount := bill.FormatNumber(float64(orderPayment.Price * orderPayment.Qty))
		item.Code = Code
		item.Name = orderPayment.FoodName
		item.MeasureUnit = "01"
		item.Qty = bill.FormatNumber(float64(orderPayment.Qty))
		item.UnitPrice = Price
		item.TotalAmount = TotalAmount
		item.CityTax = Price
		item.Vat = Price
		item.BarCode = Code
		items = append(items, item)

	}

	bilInput.Stocks = items

	log.Println("========== hey2")

	ebarimtResponse, ebarimtErr := bill.PutBill(bilInput, ebarimt.PosAPI)

	log.Println("========== hey3", ebarimtErr)
	log.Println("========== hey4", ebarimtResponse.Success)

	if ebarimtErr != nil {
		// create error info
		fmt.Println(ebarimtErr.Error())
	}

	if ebarimtResponse.Success {
		jsonStrin, _ := json.Marshal(ebarimtResponse)
		var orderEbarimt models.OrderEbarimt
		orderEbarimt.Ebarimt = string(jsonStrin)
		orderEbarimt.OrderID = order.ID

		DB.DB.Create(&orderEbarimt)
	}
}
