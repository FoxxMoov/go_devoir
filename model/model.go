// /*
// package model contains the UI-independent and migration-independent types and
// variables. It does not pretend to be a clean model.
//  */
package model

// import (
// 	// À vous...
// 	"io"
// 	"strconv"

// 	"github.com/bojanz/currency"
// )

// /*
// Currency is the single monetary unit used in this project.
// */
// const Currency = `EUR`

// /*
// Precision is the number of digits used by the chosen Currency.
// */
// var Precision = 	// À vous...à partir de Currency.

// var (
// 	// OutStd is the stream to be used for all standard output.
// 	OutStd io.Writer
// 	// OutErr is the stream to be used for all non-default output.
// 	OutErr io.Writer
// 	// OutVerb is the stream to be used for output only emitted with the verbose flag.
// 	OutVerb io.Writer
// )

// /*
// unmarshalSalary allows unmarshalling Salary fields without defining a
// local type wrapping currency.Amount.
// */
// func unmarshalSalary(u func(interface{}) error) (currency.Amount, error) {
// 	var salary struct {
// 		Salary float64
// 	}
// 	err := u(&salary)
// 	if err != nil {
// 		return currency.Amount{}, err
// 	}
// 	return currency.NewAmount(strconv.FormatFloat(salary.Salary, 'f', int(Precision), 64), Currency)
// }

// /*
// unmarshalTotalSalary allows unmarshalling TotalSalary fields without defining a
// local type wrapping currency.Amount.
// */
// func unmarshalTotalSalary(u func(interface{}) error) (currency.Amount, error) {
// 	var salary struct {
// 		TotalSalary float64 `yaml:"total_salary"`
// 	}
// 	err := u(&salary)
// 	if err != nil {
// 		return currency.Amount{}, err
// 	}
// 	return currency.NewAmount(strconv.FormatFloat(salary.TotalSalary, 'f', int(Precision), 64), Currency)
// }
