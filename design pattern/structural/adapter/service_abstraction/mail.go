package srv_abstraction

type Iemail_service interface {
	SendRegistrationEmail(userID uint) error
}

type email_service struct {
	send_registration_email Iemail_service
}

func New_email_service(ie Iemail_service) email_service {
	return email_service{
		send_registration_email: ie,
	}
}

func (e email_service) SendRegistrationEmail(userID int) error {
	return e.send_registration_email.SendRegistrationEmail(uint(userID))
}
