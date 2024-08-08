package main

import "fmt"

type Notifer interface {
	Notify(to, msg string)
}

type EmailService struct{}

func (e *EmailService) Notify(to, msg string) {
	fmt.Println("email notify")
}

type SMSService struct{}

func (s *SMSService) Notify(to, msg string) {
	fmt.Println("sms notify")
}

type Notification struct {
	notifer Notifer
}

func (n *Notification) Notify(to, msg string) {
	n.notifer.Notify(to, msg)
}

func goodInit() {
	emailService := EmailService{}
	notification := Notification{notifer: &emailService}
	notification.notifer.Notify("test@gmail.com", "Hello world")

	smsService := SMSService{}
	notification = Notification{notifer: &smsService}
	notification.notifer.Notify("test@gmail.com", "Hello world")
}
