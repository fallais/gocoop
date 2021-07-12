package email

import (
	"github.com/fallais/gocoop/pkg/notifiers"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Type is the type of the notifier.
const Type = "Email"

// Vendor is the vendor of the notifier.
const Vendor = "Email"

type email struct {
	host      string
	port      int
	recipient string
	subject   string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewProvider returns a new notifier for SMTP.
func NewProvider(settings map[string]interface{}) notifiers.Notifier {
	// Initial values
	host := ""
	port := 25
	recipient := ""
	subject := ""

	// Process the values
	for key, value := range settings {
		switch key {
		case "host":
			host = value.(string)
		case "port":
			port = int(value.(float64))
		case "recipient":
			recipient = value.(string)
		case "subject":
			subject = value.(string)
		default:
			logrus.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).Infoln("Wrong setting for SMTP")
		}
	}

	return &email{
		host:      host,
		port:      port,
		recipient: recipient,
		subject:   subject,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify sends a notification.
func (p *email) Notify(msg string) error {
	// Create the message
	m := gomail.NewMessage()
	m.SetHeader("From", p.recipient)
	m.SetHeader("To", p.recipient)
	m.SetHeader("Subject", p.subject)
	m.SetBody("text/html", msg)

	// Dial
	d := gomail.Dialer{Host: p.host, Port: p.port}

	// Send the email
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

// Type returns the type.
func (p *email) Type() string {
	return Type
}

// Vendor returns the vendor.
func (p *email) Vendor() string {
	return Vendor
}
