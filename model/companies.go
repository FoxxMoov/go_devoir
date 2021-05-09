package model

/*
Company models a company entry.
*/
type Company struct {
	Siret       int     `yaml:"siret" gorm:"primary_key;auto_increment" json:"siret"`
	Date        string  `yaml:"date" json:"date"`
	TotalSalary float64 `yaml:"total_salary" json:"total_salary"`
	Staff       []Staff `gorm:"foreignKey:Siret" yaml:"staff" json:"staff"`
}

// CreateCompany Create company
func (db *DataModel) CreateCompany(c Company) error {
	return db.Db.Debug().Create(c).Error
}

// ModifyCompanySiret Modify the siret of one company
func (db *DataModel) ModifyCompanySiret(c Company, newSiret string) error {
	return db.Db.Debug().Model(&c).Update("siret", newSiret).Error
}

// ListCompanies Retrieve all compagnies
func (db *DataModel) ListCompanies(c Company) error {
	return db.Db.Debug().Find(&c).Error
}

// DeleteCompanyBySiret Delete company by it's siret
func (db *DataModel) DeleteCompanyBySiret(siret int) error {
	c := &Company{Siret: siret}
	return db.Db.Debug().Delete(c).Error
}
