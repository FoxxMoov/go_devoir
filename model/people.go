package model

/*
Staff models a staff entry.
*/
type People struct {
	Insee  int     `yaml:"insee" gorm:"primary_key" json:"insee"`
	Given   string  `yaml:"given" json:"given"`
	Last string `yaml:"last" json:"last"`
	Salary float64 `yaml:"salary" json:"salary"`
}