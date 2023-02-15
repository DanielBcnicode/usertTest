package broker

import (
	"log"

	"github.com/streadway/amqp"
	"usertest.com/event"
)

type Broker interface {
	PublishDomainEvent(event *event.DomainEvent) error
	Close()
}

type Rabbit struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	domainQueue *amqp.Queue
}

const DomainQueue string = "DomainQueue"

func NewRabbitConnectionForDomain(url string) (*Rabbit, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("amqp connection error: %v", err)
		return &Rabbit{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR opening rabbit channel: %s\n", err)
		conn.Close()
		return &Rabbit{}, err
	}

	q, err := ch.QueueDeclare(
		DomainQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("ERROR opening DomainQueue queue: %s\n", err)
		ch.Close()
		conn.Close()

		return &Rabbit{}, err
	}

	return &Rabbit{
		conn:        conn,
		channel:     ch,
		domainQueue: &q,
	}, nil
}

func (r *Rabbit) Close() {
	r.channel.Close()
	r.conn.Close()
}

func (r *Rabbit) Consumer(queue string) (<- chan amqp.Delivery, error) {
	return r.channel.Consume(
		queue,
		"", 
		true, 
		false, 
		false, 
		false, 
		nil)
}

func (r *Rabbit) PublishDomainEvent(event *event.DomainEvent) error {

	body, err := event.Serialize()
	if err != nil {
		log.Printf("ERROR Serializing event: %s\n", err)
		return err
	}

	err = r.channel.Publish(
		"",
		DomainQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Printf("ERROR publishing event: %s\n", err)
		return err
	}

	log.Printf("Domain Event published: %s\n", event.Type)
	
	return nil
}
