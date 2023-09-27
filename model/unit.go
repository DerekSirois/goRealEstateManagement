package model

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	PropertyId uint
	RenterId   uint
	DoorNumber string
	RentAmount float32
}

func GetAllUnitByProperty(db *gorm.DB, ownerId uint) (*[]Unit, error) {
	var u []Unit
	result := db.Where("property_id = ?", ownerId).Find(&u)
	return &u, result.Error
}

func (u *Unit) GetUnitById(db *gorm.DB, id uint) error {
	result := db.First(u, id)
	return result.Error
}

func (u *Unit) CreateUnit(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}

func (u *Unit) UpdateUnit(db *gorm.DB, id uint) error {
	uDb := &Unit{}
	err := uDb.GetUnitById(db, id)
	if err != nil {
		return err
	}
	result := db.Model(uDb).Updates(u)
	return result.Error
}

func DeleteUnit(db *gorm.DB, id uint) error {
	result := db.Delete(&Unit{}, id)
	return result.Error
}
