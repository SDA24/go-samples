package person_service

import (
	person_service_abstraction "builder/service_abstraction"
	"encoding/json"
)

type person struct {
	streetAddress, postcode, city string
	companyName, position         string
	annualIncome                  int
}

func (p person) to_Person_service() person_service_abstraction.Person {
	return person_service_abstraction.Person{
		StreetAddress: p.streetAddress,
		Postcode:      p.postcode,
		City:          p.city,
		CompanyName:   p.companyName,
		Position:      p.position,
		AnnualIncome:  p.annualIncome,
	}
}

func (p person) Person_info_json() []byte {
	b, err := json.Marshal(p.to_Person_service())
	if err != nil {
		return nil
	}
	return b
}

////////////////////////////////
type PersonBuilder struct {
	person *person // needs to be inited
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&person{}}
}

func (it *PersonBuilder) Build() person {
	return person{
		streetAddress: it.person.streetAddress,
		postcode:      it.person.postcode,
		city:          it.person.city,
		companyName:   it.person.companyName, 
		position:      it.person.position,
		annualIncome:  it.person.annualIncome,
	}
}

func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.companyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.annualIncome = annualIncome
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.streetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.city = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.postcode = postcode
	return it
}
