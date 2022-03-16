package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

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
func (h *Handler) GetModel(c *gin.Context) {
	searchOrder := c.Request.FormValue("order_uid")
	model, err := h.services.GetModel(searchOrder)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(model)
}
