package models

import (
	"github.com/DeeGrant/golang-basic-crm/database"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func init() {
	database.Connect()
	db = database.GetDb()
	db.AutoMigrate(&Lead{})
}

func (l *Lead) CreateLead() *Lead {
	db.Create(&l)
	return l
}

func GetLead(id int64) *Lead {
	var lead Lead
	db.Find(&lead, id)
	return &lead
}

func GetLeads() []Lead {
	var leads []Lead
	db.Find(&leads)
	return leads
}

func DeleteLead(id int64) Lead {
	var lead Lead
	db.Delete(&lead, id)
	return lead
}

func (l *Lead) UpdateLead() *Lead {
	db.Save(&l)
	return l
}
