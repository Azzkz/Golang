package handlers

import (
	"net/http"
	"strconv"

	"GO3/models"
	"github.com/gin-gonic/gin"
)

// Глобальная переменная для хранения данных
var items = []models.Item{
	{ID: 1, Title: "Первый элемент"},
}

// Обработчик для получения списка элементов
func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// Обработчик для создания нового элемента
func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// Обработчик для обновления элемента
func UpdateItem(c *gin.Context) {
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = id
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// Обработчик для удаления элемента
func DeleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
