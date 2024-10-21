package mail_service_abstraction

type Imail_config interface {
	From(string) Imail_config
	To(string) Imail_config
	Subject(string) Imail_config
	Body(string) Imail_config
}

type Imail interface {
	Send_email(func(i Imail_config))
}

type mail struct {
	imail Imail
}

func New_mail(i Imail) mail {
	return mail{imail: i}
}

func (m mail) Send_email(f func(i Imail_config)) {
	m.imail.Send_email(f)
}
