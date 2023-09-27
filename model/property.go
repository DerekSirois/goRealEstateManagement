package model

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	OwnerId uint
	Address string
}

func GetAllPropertyByOwner(db *gorm.DB, ownerId uint) (*[]Property, error) {
	var p []Property
	result := db.Where("owner_id = ?", ownerId).Find(&p)
	return &p, result.Error
}

func (p *Property) GetPropertyById(db *gorm.DB, id uint) error {
	result := db.First(p, id)
	return result.Error
}

func (p *Property) CreateProperty(db *gorm.DB) error {
	result := db.Create(p)
	return result.Error
}

func (p *Property) UpdateProperty(db *gorm.DB, id uint) error {
	pDb := &Property{}
	err := pDb.GetPropertyById(db, id)
	if err != nil {
		return err
	}
	result := db.Model(pDb).Updates(p)
	return result.Error
}

func DeleteProperty(db *gorm.DB, id uint) error {
	result := db.Delete(&Property{}, id)
	return result.Error
}
