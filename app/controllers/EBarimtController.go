package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/ebarimt/posapi"
	"io/ioutil"
	"lambda/app/models"
	"lambda/ebarimt"
	"net/http"
)

func EBarimtInfo(c *fiber.Ctx) error {

	info, err := ebarimt.PosAPI.GetInformation()
	if err != nil {
		return err
	}

	return c.JSON(info)
}

func EBarimtCheck(c *fiber.Ctx) error {

	res, err := ebarimt.PosAPI.CheckApi()
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func EBarimtSend(c *fiber.Ctx) error {

	res, err := ebarimt.PosAPI.SendData()
	if err != nil {
		return err
	}

	return c.JSON(res)
}
func GetOrganizationInfo(regID string) (models.OrganizationInfo, error) {
	var response models.OrganizationInfo
	url := "http://info.ebarimt.mn/rest/merchant/info?regno=" + regID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return response, err
	}
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func EBarimtPut(c *fiber.Ctx) error {

	data := `{
  "amount":"1100.00",
  "bankTransactions":[],
  "billIdSuffix":"123459",
  "billType":"1",
  "branchNo":"001",
  "cashAmount":"0.00",
  "cityTax":"0.00",
  "customerNo":"",
  "districtCode":"24",
  "invoiceId":"",
  "nonCashAmount":"1100.00",
  "posNo":"1234",
  "reportMonth":"",
  "returnBillId":"",
  "stocks":[
    {
      "barCode":"1234567890",
      "cityTax":"0.00",
      "code":"1234567",
      "measureUnit":"01",
      "name":"Хоол",
      "qty":"1.00",
      "totalAmount":"1100.00",
      "unitPrice":"1100.00",
      "vat":"100.00"
    }
  ],
  "taxType":"",
  "vat":"100.00"
}`
	input := posapi.PutInput{}
	json.Unmarshal([]byte(data), &input)
	result, err := ebarimt.PosAPI.Put(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to call Put function"})
	}

	return c.JSON(fiber.Map{"result": result})
}
