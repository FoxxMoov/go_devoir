package model

import "gorm.io/gorm"

/*
Staff models a staff entry.
*/
type Staff struct {
	Insee  int     `yaml:"insee" gorm:"primary_key" json:"insee"`
	Name   string  `yaml:"name" json:"name"`
	Salary float64 `yaml:"salary" json:"salary"`
	Siret  int     `yaml:"-" json:"siret"`
}

// CreateStaff Create staff
func (db *DataModel) CreateStaff(c Staff) error {
	return db.Db.Debug().Create(c).Error
}

// DeleteStaffByInsee Delete staff by insee
func (db *DataModel) DeleteStaffByInsee(insee int) error {
	s := &Staff{Insee: insee}
	return db.Db.Debug().Delete(s).Error
}

// DeleteStaffs Delete staffs
func (db *DataModel) DeleteStaffs() error {
	var s Staff
	return db.Db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(s).Error
}
