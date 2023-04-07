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
	"github.com/leekchan/accounting"
	"lambda/app/models"
	"lambda/ebarimt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		"callback_url":          "https://mmk.khankhulgun.mn/api/qpay/callback/" + qpayRequest.SenderInvoiceNo,
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
	DB.DB.Where("order_number = ?", qpayRequest.SenderInvoiceNo).First(&updateOrder)
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
	var orderNumber = c.Params("invoice_id")

	checkOrder := models.Orders{}
	DB.DB.Where("order_number = ?", orderNumber).First(&checkOrder)

	log.Println(checkOrder.InvoiceID)

	if QpayCallBackCheck(checkOrder.InvoiceID) > 0 {

		UpdateStatus(orderNumber, checkOrder.ID, "qpay", "success")

		var orderDetails []models.OrderDetail
		DB.DB.Where("order_id = ?", checkOrder.ID).Find(&orderDetails)

		for _, orderDetail := range orderDetails {
			UpdateBalance(orderDetail.FoodID, orderDetail.KitchenID, orderDetail.Qty)
		}

		var orderDetailSets []models.OrderDetailSet
		DB.DB.Where("order_id = ?", checkOrder.ID).Find(&orderDetailSets)

		for _, orderDetailSet := range orderDetailSets {
			UpdateBalance(orderDetailSet.FoodID, orderDetailSet.KitchenID, orderDetailSet.Quantity)
		}

		return c.Status(http.StatusOK).JSON("SUCCESS")
	}

	return c.Status(http.StatusOK).JSON("FAILED")
}

func LaterPay(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	orderStatus := models.OrdersStatus{}
	orders := models.ViewOrder{}

	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Order("id DESC").Find(&orders)
	DB.DB.Where("id = ?", orders.ID).Find(&orderStatus)

	if orders.ID == 0 {
		return c.Status(http.StatusOK).JSON("Not found active order")
	}

	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&orders)

	orderLaterPay := models.OrderLaterPay{}

	orderLaterPay.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	orderLaterPay.OrderNumber = orders.OrderNumber
	orderLaterPay.OrderID = orders.ID
	orderLaterPay.Qty = orders.OrderQuantity
	orderLaterPay.Price = orders.Price
	orderLaterPay.PaymentStatus = GetStringPointer("success")

	oldOrders := models.Orders{}
	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Find(&oldOrders)

	if oldOrders.ID >= 1 {
		UpdateStatus(orders.OrderNumber, oldOrders.ID, "mmk", "success")
		var orderDetails []models.OrderDetail
		DB.DB.Where("order_id = ? AND user_id = ?", oldOrders.ID, orderUser["id"]).Find(&orderDetails)

		for _, orderDetail := range orderDetails {
			UpdateBalance(orderDetail.FoodID, orderDetail.KitchenID, orderDetail.Qty)
		}

		var orderDetailSets []models.OrderDetailSet
		DB.DB.Where("order_id = ?", oldOrders.ID).Find(&orderDetailSets)

		for _, orderDetailSet := range orderDetailSets {
			UpdateBalance(orderDetailSet.FoodID, orderDetailSet.KitchenID, orderDetailSet.Quantity)
		}

	} else {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Идэвхтэй захиалга олдсонгүй",
		})
	}

	DB.DB.Create(&orderLaterPay)

	// create Bill
	bilInput := posapi.PutInput{}

	amount := bill.FormatNumber(float64(orderLaterPay.Price))
	bilInput.Amount = amount
	bilInput.Vat = bill.GetVat(float64(orderLaterPay.Price), true, false)
	bilInput.BillIDSuffix = bill.GenerateBillIdSuffix()
	bilInput.CityTax = "0.00"
	bilInput.BillType = "1"
	bilInput.CashAmount = "0.00"
	bilInput.NonCashAmount = amount
	bilInput.DistrictCode = "23"
	bilInput.PosNo = "0001"
	bilInput.BranchNo = "001"

	var items []posapi.Stock

	var orderPayments []models.ViewOrderDetail

	for _, orderPayment := range orderPayments {
		var item posapi.Stock

		Code := strconv.Itoa(orderPayment.OrderID)
		Price := bill.FormatNumber(float64(orderPayment.Price))
		item.Code = Code
		item.Name = orderPayment.FoodName
		item.MeasureUnit = "01"
		item.Qty = Number(orderPayment.Qty)
		item.UnitPrice = Price
		item.TotalAmount = Price
		item.CityTax = Price
		item.Vat = Price
		item.BarCode = Code

		items = append(items, item)
	}

	fmt.Println(items)
	//appent

	bilInput.Stocks = items

	ebarimtResponse, ebarimtErr := bill.PutBill(bilInput, ebarimt.PosAPI)

	if ebarimtErr != nil {
		// create error info
		fmt.Println(ebarimtErr.Error())
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":            "success",
		"message":           orderLaterPay.OrderNumber + " дугаартай захиалга амжилттай",
		"billId":            ebarimtResponse.BillID,
		"billType":          ebarimtResponse.BillType,
		"date":              ebarimtResponse.Date,
		"internalCode":      ebarimtResponse.InternalCode,
		"lottery":           ebarimtResponse.Lottery,
		"lotteryWarningMsg": ebarimtResponse.LotteryWarningMsg,
		"macAddress":        ebarimtResponse.MacAddress,
		"qrData":            ebarimtResponse.QRData,
		"registerNo":        ebarimtResponse.RegisterNo,
	})
}

func Number(v interface{}) string {
	ac := accounting.Accounting{Precision: 2}
	return strings.ReplaceAll(ac.FormatMoney(v), ",", "")
}
