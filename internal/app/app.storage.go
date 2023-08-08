package app

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID         uint64          `json:"id" gorm:"primaryKey"`
	Redirect   string          `json:"redirect" gorm:"not null"`
	Goly       string          `json:"goly" gorm:"unique;not null"`
	Clicked    uint64          `json:"clicked"`
	Random     bool            `json:"random"`
}

func getAllGolies() ([]Goly, error) {
	var golies []Goly
	
	tx := db.Find(&golies)
	if tx.Error != nil {
		return []Goly{}, tx.Error
	}

	return golies, nil
}

func getGoly(id uint64) (Goly, error) {
	var goly Goly

	tx := db.Where("id = ?", id).First(&goly)
	if tx.Error != nil {
		return Goly{}, tx.Error
	}
	return goly, nil
}

func createGoly(goly Goly) error {
	tx := db.Create(&goly)
	return tx.Error
}

func updateGoly(goly Goly) error {
	tx := db.Save(&goly)
	return tx.Error
}

func deleteGoly(id uint64) error {
	tx := db.Unscoped().Delete(&Goly{}, id) 
	return tx.Error
}

func getGolyByUrl(url string) (Goly, error) {
	var goly Goly 
	tx := db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}