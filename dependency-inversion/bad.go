package main

type BadEmailNotification struct {
	emailService *EmailService
}

func (b *BadEmailNotification) Notify(to, msg string) {
	b.emailService.Notify(to, msg)
}

type BadSMSNotification struct {
	smsService *SMSService
}

func (b *BadSMSNotification) Notify(to, msg string) {
	b.smsService.Notify(to, msg)
}

func initBad() {
	emailService := EmailService{}
	badEmailNotification := BadEmailNotification{emailService: &emailService}
	badEmailNotification.Notify("bad@gmail.com", "hello world")

	smsService := SMSService{}
	badSMSNotification := BadSMSNotification{smsService: &smsService}
	badSMSNotification.Notify("bad@gmail.com", "hello world")
}
