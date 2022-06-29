package controllers

import (
	"fmt"
	"github.com/DeeGrant/golang-basic-crm/models"
	"github.com/gofiber/fiber"
	"strconv"
)

func GetLeads(c *fiber.Ctx) {
	leads := models.GetLeads()
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	leadId := c.Params("id")
	id, err := strconv.ParseInt(leadId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	lead := models.GetLead(id)
	c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) {
	var lead models.Lead
	err := c.BodyParser(&lead)
	if err != nil {
		panic("Error while parsing request body.")
	}

	newLead := lead.CreateLead()

	c.JSON(newLead)
}

func DeleteLead(c *fiber.Ctx) {
	leadId := c.Params("id")
	id, err := strconv.ParseInt(leadId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	lead := models.DeleteLead(id)
	c.JSON(lead)
}

func UpdateLead(c *fiber.Ctx) {
	leadId := c.Params("id")
	id, err := strconv.ParseInt(leadId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	lead := models.GetLead(id)

	err = c.BodyParser(&lead)
	if err != nil {
		panic("Error while parsing request body.")
	}

	l := lead.UpdateLead()

	c.JSON(l)
}
