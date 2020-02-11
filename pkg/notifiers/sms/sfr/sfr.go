package sfr

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fallais/gocoop/pkg/notifiers"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Type is the type of the notifier.
const Type = "SMS"

// Vendor is the vendor of the notifier.
const Vendor = "SFR"

type sfr struct {
	client *http.Client
	url    string
	token  string
	to     string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewProvider returns a new notifier for SFR.
func NewProvider(settings map[string]interface{}) notifiers.Notifier {
	// Set HTTP transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Set timeout
	timeout := time.Duration(12 * time.Second)

	// Set HTTP Client
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	// Initial values
	token := ""
	to := ""

	// Process the values
	for key, value := range settings {
		switch key {
		case "token":
			token = value.(string)
		case "to":
			to = value.(string)
		default:
			logrus.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).Infoln("Wrong setting for SFR")
		}
	}

	return &sfr{
		client: client,
		url:    "http://ws.red.sfr.fr",
		token:  token,
		to:     to,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify sends a notification.
func (s *sfr) Notify(msg string) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(s.url)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/red-ws/red-b2c/resources/sms/send"
	parameters := url.Values{}
	parameters.Add("responseType", "json")
	parameters.Add("token", s.token)
	parameters.Add("to", s.to)
	parameters.Add("type", "PhoneNumber")
	parameters.Add("msg", msg)
	reqURL.RawQuery = parameters.Encode()

	// Create the request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Do the request
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error while sending the SMS. Status code is %d", resp.StatusCode)
	}

	return nil
}

// Type returns the type.
func (s *sfr) Type() string {
	return Type
}

// Vendor returns the vendor.
func (s *sfr) Vendor() string {
	return Vendor
}
