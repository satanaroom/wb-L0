package handler

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	broker "github.com/satanaroom/L0"
	"github.com/sirupsen/logrus"
)

// Метод для отображения HTML-страницы
func (h *Handler) CreateHTML(c *gin.Context) {
	page, err := template.ParseFiles("index.html")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = page.Execute(c.Writer, nil)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

// Метод для получения заказа по его id и отображения в специальном поле информации о нем
func (h *Handler) GetModel(c *gin.Context) {
	searchOrder := c.Request.FormValue("order_uid")
	model, err := h.services.GetModelCache(searchOrder)
	if err != nil {
		logrus.Fatalln(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Создана новая модель в которой отсутсвуют поля, добавленные для работы с БД
	var newModel broker.ReturnedModel
	newModel.OrderUid = model.OrderUid
	newModel.TrackNumber = model.TrackNumber
	newModel.Entry = model.Entry
	newModel.Delivery.Address = model.Delivery.Address
	newModel.Delivery.City = model.Delivery.City
	newModel.Delivery.Email = model.Delivery.Email
	newModel.Delivery.Name = model.Delivery.Name
	newModel.Delivery.Phone = model.Delivery.Phone
	newModel.Delivery.Region = model.Delivery.Region
	newModel.Delivery.Zip = model.Delivery.Zip
	newModel.Payment.Amount = model.Payment.Amount
	newModel.Payment.Bank = model.Payment.Bank
	newModel.Payment.Currency = model.Payment.Currency
	newModel.Payment.CustomFee = model.Payment.CustomFee
	newModel.Payment.DeliveryCost = model.Payment.DeliveryCost
	newModel.Payment.GoodsTotal = model.Payment.GoodsTotal
	newModel.Payment.PaymentDt = model.Payment.PaymentDt
	newModel.Payment.Provider = model.Payment.Provider
	newModel.Payment.RequestId = model.Payment.RequestId
	newModel.Payment.Transaction = model.Payment.Transaction
	newModel.Locale = model.Locale
	newModel.InternalSignature = model.InternalSignature
	newModel.CustomerId = model.CustomerId
	newModel.DeliveryService = model.DeliveryService
	newModel.Shardkey = model.Shardkey
	newModel.SmId = model.SmId
	newModel.DateCreated = model.DateCreated
	newModel.OofShard = model.OofShard
	newModel.Items = []struct {
		OrderUid    string
		ChrtId      int    "json:\"chrt_id\""
		TrackNumber string "json:\"track_number\""
		Price       int    "json:\"price\""
		Rid         string "json:\"rid\""
		Name        string "json:\"name\""
		Sale        int    "json:\"sale\""
		Size        string "json:\"size\""
		TotalPrice  int    "json:\"total_price\""
		NmId        int    "json:\"nm_id\""
		Brand       string "json:\"brand\""
		Status      int    "json:\"status\""
	}(model.Items)
	for i := range model.Items {
		newModel.Items[i].ChrtId = model.Items[i].ChrtId
		newModel.Items[i].TrackNumber = model.Items[i].TrackNumber
		newModel.Items[i].Price = model.Items[i].Price
		newModel.Items[i].Rid = model.Items[i].Rid
		newModel.Items[i].Name = model.Items[i].Name
		newModel.Items[i].Sale = model.Items[i].Sale
		newModel.Items[i].Size = model.Items[i].Size
		newModel.Items[i].TotalPrice = model.Items[i].TotalPrice
		newModel.Items[i].NmId = model.Items[i].NmId
		newModel.Items[i].Brand = model.Items[i].Brand
		newModel.Items[i].Status = model.Items[i].Status
	}
	page, err := template.ParseFiles("index.html")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = page.Execute(c.Writer, newModel)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
