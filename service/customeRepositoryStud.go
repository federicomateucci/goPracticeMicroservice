package service

import "practiceRestServices/Dtos"

type CustomerRepositoryStud struct {
	customers []Dtos.Customer
}

func (s CustomerRepositoryStud) FindAll() ([]Dtos.Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStud {

	customers := []Dtos.Customer{
		{Id: "1001", Name: "Angeles Antonini", City: "La Paternal", Zipcode: "1414", DateOfBirth: "22/05/1989", Status: "Hot"},
		{Id: "1002", Name: "Federico Mateucci", City: "Villa Crespo", Zipcode: "1414", DateOfBirth: "28/01/1988", Status: "Hot"},
	}
	return CustomerRepositoryStud{customers: customers}
}
