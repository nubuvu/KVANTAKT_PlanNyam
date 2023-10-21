package handler

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	_ "reflect"
	"strconv"
)

func (h *Handler) parseJsonCheque(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	ids, err := h.services.JsonCheque.ParseJsonCheque(listId, jsonData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id_add_items": ids,
	})
}
