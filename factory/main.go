package main

import "fmt"

// interface Envia Notificación
type NotificationFactory interface {
	SendNotification()
	GetSender() Sender
}

func getNotificationFactory(tNot string) (NotificationFactory, error) {
	if tNot == "SMS" {
		return &SMSNotification{}, nil
	}
	if tNot == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("notificacion no soportada")
}
func sendNotification(n NotificationFactory) {
	n.SendNotification()
}
func getMethod(n NotificationFactory) {
	fmt.Println(n.GetSender().GetSenderMethod())
}

//Determina el Remitente
type Sender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//Determinar que se envia SMS
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Enviando SMS")
}

func (SMSNotification) GetSender() Sender {
	return SMSNotificationSender{}
}

//Deteminar remitente SMS
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Claro"
}

//Determina que se envia EMAIl
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando EMAIL")
}
func (EmailNotification) GetSender() Sender {
	return EmailNotificationSender{}
}

//Determinan el remitente EMAIl
type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Hotmail"
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emialFactor, _ := getNotificationFactory("Email")
	sendNotification(smsFactory)
	sendNotification(emialFactor)
	getMethod(smsFactory)
	getMethod(emialFactor)
}
