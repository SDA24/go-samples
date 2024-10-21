package mail_service

import (
	mail_service_abstraction "builder/service_abstraction"
	"fmt"
	"strings"
)

type email struct {
	from    string
	to      string
	subject string
	body    string
}

type Email_builder struct {
	email *email
}

func New_email_builder() *Email_builder {
	return &Email_builder{&email{}}
}

func (b *Email_builder) From(from string) mail_service_abstraction.Imail_config {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *Email_builder) To(to string) mail_service_abstraction.Imail_config {
	b.email.to = to
	return b
}

func (b *Email_builder) Subject(subject string) mail_service_abstraction.Imail_config {
	b.email.subject = subject
	return b
}

func (b *Email_builder) Body(body string) mail_service_abstraction.Imail_config {
	b.email.body = body
	return b
}

func (b *Email_builder) Build() email {
	return email{
		from:    b.email.from,
		to:      b.email.to,
		subject: b.email.subject,
		body:    b.email.body,
	}
}

func (e email) send_email_implementaion(email *email) {
	fmt.Println("sending email: ", *email)
}
func (e email) Send_email(action func(i mail_service_abstraction.Imail_config)) {
	builder := New_email_builder()
	action(builder)
	e.send_email_implementaion(builder.email)
}
