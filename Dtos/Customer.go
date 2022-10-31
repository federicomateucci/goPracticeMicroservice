package Dtos

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city"  xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Status      string `json:"status,omitempty"`
}

// interface repository methods implement
type ICustomerRepository interface {
	FindAll() ([]Customer, error)
}
