package main

import (
	mail_service "builder/service"
	mail_service_abstraction "builder/service_abstraction"
)

func main() {
	adapter := mail_service.New_email_builder()
	serve := mail_service_abstraction.New_mail(adapter.Build())
	serve.Send_email(func(i mail_service_abstraction.Imail_config) {
		i.From("Saeed@gmail.com").To("Delaram@gmail.com").Subject("Test").Body("I love you")
	})
}
