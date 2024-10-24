package controllers

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateContact(c echo.Context) error {
	var contact models.Contact
	if err := c.Bind(&contact); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := config.DB.Create(&contact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, contact)
}

func GetContacts(c echo.Context) error {
	var contacts []models.Contact
	if err := config.DB.Find(&contacts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, contacts)
}

func GetContactByID(c echo.Context) error {
	var contact models.Contact
	id := c.Param("id")
	if err := config.DB.First(&contact, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Contact not found"})
	}
	return c.JSON(http.StatusOK, contact)
}

func UpdateContact(c echo.Context) error {
	var contact models.Contact
	id := c.Param("id")
	if err := config.DB.First(&contact, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Contact not found"})
	}
	if err := c.Bind(&contact); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := config.DB.Save(&contact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, contact)
}

func DeleteContact(c echo.Context) error {
	var contact models.Contact
	id := c.Param("id")
	if err := config.DB.First(&contact, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Contact not found"})
	}
	if err := config.DB.Delete(&contact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Contact deleted"})
}
