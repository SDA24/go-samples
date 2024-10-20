package service_abstraction

type Ielement_service interface {
	String() string
}

type element struct {
	ielement_service Ielement_service
}

func New_element(i Ielement_service) element {
	return element{ielement_service: i}
}

func (e element) String() string {
	return e.ielement_service.String()
}
