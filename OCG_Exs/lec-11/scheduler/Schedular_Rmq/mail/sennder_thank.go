package mail

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
)

// Mailer defines function for sending email
type Mailer interface {
	Send(*EmailContent) error
}

// EmailUser defines email address info
type EmailUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// encoding to json

func (eu *EmailUser) String() string {
	b, _ := json.Marshal(eu)
	return string(b)
}

// EmailContent defines email content info
type EmailContent struct {
	ID               int64      `json:"id"`
	Subject          string     `json:"subject"`
	FromUser         *EmailUser `json:"from"`
	ToUser           *EmailUser `json:"to"`
	PlainTextContent string     `json:"plaintext_content"`
	HtmlContent      string     `json:"html_content"`
}

// encoding mail
func (em *EmailContent) String() string {
	b, _ := json.Marshal(em)
	return string(b)
}

// Validate will check whether the email content is valid
func (em *EmailContent) Validate() error {
	if em == nil || em.FromUser == nil || em.ToUser == nil || em.PlainTextContent == "" {
		return errors.New("wrong content")
	}
	return nil
}

//new client NewSendgrid creates new Sendgrid client
func NewSendgrid(apiKey string) *Sendgrid {
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		Client: client,
	}
}

// Sendgrid implements logic to send email to destination(location/ diem den) email address via sendgrid
type Sendgrid struct {
	Client *sendgrid.Client
}

// Send will send email based on email content
func (m *Sendgrid) Send(em *EmailContent) error {
	if err := em.Validate(); err != nil {
		return err
	}

	fmt.Println("Sending email: ", em)
	return nil
}
