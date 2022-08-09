package models

import "time"

type Owner struct {
	OwnerID          uint      `json:"id" gorm:"primary_key"`
	Name             string    `json:"name"`
	IDCard           string    `json:"idCard"`
	Address          string    `json:"address"`
	NoWa             string    `json:"noWa"`
	Email            string    `json:"email"`
	BusinessName     string    `json:"businessName"`
	TaxObject        string    `json:"taxObject"`
	TaxObjectAddress string    `json:"taxObjectAddress"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
