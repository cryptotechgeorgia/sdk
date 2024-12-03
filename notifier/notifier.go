package notifier

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

var (
	DefaultNotifier Notifier
)

type Config struct {
	Destination     string
	DestinationType DestinationType
	Template        string
	Language        string
	Subject         string
	Source          string
}

// publisher payload
type request struct {
	Template        string `json:"template"`
	Language        string `json:"language"`
	Payload         string `json:"payload"`
	Subject         string `json:"subject"`
	Source          string `json:"source"`
	DestinationType string `json:"destinationType"`
	Destination     string `json:"destination"`
	Expire          struct {
		Key string  `json:"key"`
		Ttl float64 `json:"ttl"`
	} `json:"expire"`
}

type Publisher interface {
	Publish(ctx context.Context, message interface{}) error
}

type Notifier struct {
	publisher Publisher
	opts      Config
}

func NewNotifier(client Publisher, opts Config) *Notifier {
	return &Notifier{
		publisher: client,
		opts:      opts,
	}
}

func (a *Notifier) WithPublisher(client Publisher) *Notifier {
	a.publisher = client
	return a
}

func (a *Notifier) GetPublisher() Publisher {
	return a.publisher
}

func (a *Notifier) WithDestinationType(dstType DestinationType) *Notifier {
	a.opts.DestinationType = dstType
	return a
}

func (a *Notifier) WithDestination(dst string) *Notifier {
	a.opts.Destination = dst
	return a
}

func (a *Notifier) WithLanguage(language string) *Notifier {
	a.opts.Language = language
	return a
}

func (a *Notifier) WithSubject(subject string) *Notifier {
	a.opts.Subject = subject
	return a
}

func (n *Notifier) Notify(ctx context.Context, message []byte, ttl float64) {
	req := request{
		Template:        n.opts.Template,
		Language:        n.opts.Language,
		Payload:         string(message),
		Subject:         n.opts.Subject,
		Source:          n.opts.Source,
		DestinationType: n.opts.DestinationType.String(),
		Destination:     n.opts.Destination,
		Expire: struct {
			Key string  "json:\"key\""
			Ttl float64 "json:\"ttl\""
		}{Key: uuid.NewString(), Ttl: ttl},
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Printf("notifier: Notify marshall request %s\n", err.Error())
		return
	}

	if err := n.publisher.Publish(ctx, string(b)); err != nil {
		log.Printf("notifier:  Notify  publish: %s\n", err.Error())
	}

}
