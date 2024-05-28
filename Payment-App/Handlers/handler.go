package Handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	database "github.com/huybuine/Payment-App/Database"
	"github.com/huybuine/Payment-App/Models"
)

// GetPaymentDetails trả về danh sách tất cả chi tiết thanh toán
func GetPaymentDetails(c *fiber.Ctx) error {
	db := database.DB.Db
	var payments []Models.PaymentDetail
	db.Find(&payments)
	if len(payments) == 0 {
		return c.Status(404).JSON(fiber.Map{"data": []Models.PaymentDetail{}})
	}
	return c.Status(200).JSON(payments)
}

// GetPaymentDetail trả về chi tiết thanh toán theo ID
func GetPaymentDetail(c *fiber.Ctx) error {
	db := database.DB.Db
	var PaymentDetail Models.PaymentDetail
	id, er := strconv.Atoi(c.Params("id"))
	if er != nil {
		return er
	}
	db.Find(&PaymentDetail, "ID = ?", id)

	return c.Status(200).JSON(PaymentDetail)
}

// CreatePaymentDetail tạo một chi tiết thanh toán mới
func CreatePaymentDetail(c *fiber.Ctx) error {
	db := database.DB.Db
	PaymentDetail := new(Models.PaymentDetail)

	err := c.BodyParser(PaymentDetail)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}
	err = db.Create(&PaymentDetail).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}
	var payments []Models.PaymentDetail
	db.Find(&payments)
	return c.Status(201).JSON(payments)
}

// UpdatePaymentDetail cập nhật một chi tiết thanh toán đã có
func UpdatePaymentDetail(c *fiber.Ctx) error {
	type UpdatePaymentDetail struct {
		CardOwnerName  string `json:"card_owner_name"`
		CardNumber     string `json:"cardNumber"`
		ExpirationDate string `json:"expirationDate"`
		SecurityCode   string `json:"securityCode"`
	}
	db := database.DB.Db
	var PaymentDetail Models.PaymentDetail
	// get id params
	id, er := strconv.Atoi(c.Params("id"))
	if er != nil {
		return er
	}
	db.Find(&PaymentDetail, "ID = ?", id)

	var updatePaymentDetailData UpdatePaymentDetail
	err := c.BodyParser(&updatePaymentDetailData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}
	PaymentDetail.CardOwnerName = updatePaymentDetailData.CardOwnerName
	PaymentDetail.CardNumber = updatePaymentDetailData.CardNumber
	PaymentDetail.ExpirationDate = updatePaymentDetailData.ExpirationDate
	PaymentDetail.SecurityCode = updatePaymentDetailData.SecurityCode
	db.Save(&PaymentDetail)
	var payments []Models.PaymentDetail
	db.Find(&payments)
	return c.Status(200).JSON(payments)
}

// DeletePaymentDetail xóa một chi tiết thanh toán theo ID
func DeletePaymentDetail(c *fiber.Ctx) error {
	db := database.DB.Db
	var PaymentDetail Models.PaymentDetail
	id, er := strconv.Atoi(c.Params("id"))
	if er != nil {
		return er
	}
	db.Find(&PaymentDetail, "ID = ?", id)

	err := db.Delete(&PaymentDetail, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil})
	}
	// Query all PaymentDetails again after deletion
	var payments []Models.PaymentDetail
	db.Find(&payments)
	return c.Status(200).JSON(payments)
}
