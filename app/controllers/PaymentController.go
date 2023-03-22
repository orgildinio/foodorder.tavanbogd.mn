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

	//sender_invoice_no
	//updateOrder := models.Orders{}
	//DB.DB.Where("order_number = ?", qpayRequest.SenderInvoiceNo).First(&updateOrder)
	//updateOrder.InvoiceId = fmt.Sprintf("%s", res["invoice_id"])
	//DB.DB.Save(&updateOrder)

	return c.JSON(res)
}

func QPayPaymentCheck(c *fiber.Ctx) error {
	qpayCheck := &models.QpayCheck{}
	err := c.BodyParser(&qpayCheck)
	request, _ := http.NewRequest("POST", "https://merchant.qpay.mn/v2/auth/token", nil)
	request.Header.Set("Authorization", "Basic Q0lSQ0xFX1BBUks6Wjl1NVhwb3A=")

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
		UpdateStatus(qpayCheck.OrderNumber, "qPay", qpayCheck.EbarimtType, &qpayCheck.OrgRegisterNumber)
	}
	return c.JSON(qpayResponse)
}

func QpayCallBackCheck(invoiceId string) int {
	request, _ := http.NewRequest("POST", "https://merchant.qpay.mn/v2/auth/token", nil)
	request.Header.Set("Authorization", "Basic Q0lSQ0xFX1BBUks6Wjl1NVhwb3A=")

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

	if QpayCallBackCheck(string(checkOrder.InvoiceID)) > 0 {
		// ene ued tuhain order payment status = success
		return c.Status(http.StatusOK).JSON("SUCCESS")
	}

	return c.Status(http.StatusOK).JSON("SUCCESS")
}

func LaterPay(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	//orderStatus := models.OrdersStatus{}

	orders := models.ViewOrder{}
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&orders)

	if orders.UserID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Order not found",
		})
	}

	orderLaterPay := models.OrderLaterPay{}

	orderLaterPay.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	orderLaterPay.OrderNumber = orders.OrderNumber
	orderLaterPay.OrderID = orders.ID
	orderLaterPay.Qty = orders.OrderQuantity
	orderLaterPay.Price = orders.Price
	orderLaterPay.PaymentStatus = GetStringPointer("success")

	oldOrders := models.Orders{}

	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Find(&oldOrders)
	fmt.Println(oldOrders.ID)

	if oldOrders.ID >= 1 {

		oldOrders.PaymentStatus = GetStringPointer("success")
		oldOrders.PaymentType = GetStringPointer("mmk")
		DB.DB.Save(&oldOrders)

	} else {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Идэвхтэй захиалга олдсонгүй",
		})
	}

	DB.DB.Create(&orderLaterPay)

	fmt.Println(orderLaterPay.OrderNumber)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": orderLaterPay.OrderNumber + " дугаартай захиалга",
	})

}
