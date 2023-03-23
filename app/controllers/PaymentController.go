package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
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
	DB.DB.Save(&updateOrder)

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
	request.Header.Set("Authorization", "Basic QPayPaymentCheck")

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
	var orderNumber = c.Params("invoice_id")

	checkOrder := models.Orders{}
	DB.DB.Where("order_number = ?", orderNumber).First(&checkOrder)

	if QpayCallBackCheck(checkOrder.InvoiceID) > 0 {
		//Odoo hiij bgaa
		return c.Status(http.StatusOK).JSON("SUCCESS")
	}

	return c.Status(http.StatusOK).JSON("SUCCESS")
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

	UpdateStatus(orderStatus.OrderNumber, "mmk", "success")

	return c.Status(http.StatusOK).JSON("success")
}
