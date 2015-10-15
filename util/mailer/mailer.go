package main

import (
	"flag"
	"fmt"
	"github.com/budiariyanto/email"
	"github.com/budiariyanto/lemon/logger"
	"net/smtp"
	"strings"
)

type Message struct {
	*email.Message
}

func NewMessage(subject string, body string) *Message {
	message := new(Message)
	message.Message = email.NewMessage(subject, body)
	return message
}

func (m *Message) Attach(filenames []string) error {
	for _, filename := range filenames {
		if err := m.Message.Attach(filename); err != nil {
			return err
		}
	}

	return nil
}

func (m *Message) Inline(filenames []string) error {
	for _, filename := range filenames {
		if err := m.Message.Inline(filename); err != nil {
			return err
		}
	}

	return nil
}

func SendEmail(host string, port string, auth smtp.Auth, message *Message) error {
	err := email.Send(fmt.Sprintf("%s:%s", host, port), auth, message.Message)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	log := logger.NewStdLogger()

	subject := flag.String("subject", "No Subject", "Email subject")
	from := flag.String("from", "", "Email sender")
	recipients := flag.String("to", "", "Email recipients. Support multiple recipients, separate recipients with comma")
	messageBody := flag.String("message", "", "Email message")
	attachments := flag.String("attachments", "", "Email attachments. Support multiple attachments, separate filepaths with comma")
	username := flag.String("username", "", "Mail server username")
	password := flag.String("password", "", "Mail server password")
	host := flag.String("host", "smtp.gmail.com", "Mail server host")
	port := flag.String("port", "587", "Mail server port")
	flag.Parse()

	message := NewMessage(*subject, *messageBody)
	message.From = *from
	message.To = strings.Split(*recipients, ",")
	err := message.Attach(strings.Split(*attachments, ","))
	if err != nil {
		log.Fatal(err.Error())
	}

	auth := smtp.PlainAuth("", *username, *password, *host)

	err = SendEmail(*host, *port, auth, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Email sent.")
}
