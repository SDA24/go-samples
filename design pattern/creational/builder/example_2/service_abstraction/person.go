package person_service_abstraction

const (
	Json = iota
)

type Person struct {
	StreetAddress string `json:"street_address,omitempty"`
	Postcode      string `json:"post_code,omitempty"`
	City          string `json:"city,omitempty"`
	CompanyName   string `json:"company_name,omitempty"`
	Position      string `json:"position,omitempty"`
	AnnualIncome  int    `json:"annual_income,omitempty"`
}

type Iperson_format interface {
	Person_info_json() []byte
}

type person_format struct {
	iperson_format Iperson_format
}

func New_person(i Iperson_format) person_format {
	return person_format{iperson_format: i}
}

func (p person_format) Person_info_json() []byte {
	return p.iperson_format.Person_info_json()
}
