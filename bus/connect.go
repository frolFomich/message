package bus

import (
	doc "github.com/frolFomich/abstract-document"
	"github.com/frolFomich/message"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
)

const (
	NatsUrlEnvVariableName = "NATS_URL"
)

var (
	natsConn      *nats.Conn
	busMutex      sync.Mutex
	js            nats.JetStreamContext
	subscriptions = map[string][]message.Subscriber{}
	publishers    = map[string]*natsSubscriber{}
)

type natsSubscriber struct {
}

func connection() nats.JetStreamContext {
	busMutex.Lock()
	defer busMutex.Unlock()

	if natsConn != nil && js != nil {
		return js
	}
	natsURL := nats.DefaultURL
	if e := os.Getenv(NatsUrlEnvVariableName); e != "" {
		natsURL = e
	}
	log.Infof("Connecting nats at: [%s]", natsURL)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Errorf("Couldn't connect nats: %v", err)
		panic(err)
	}
	natsConn = nc
	log.Printf("creating jet stream")
	js, err = natsConn.JetStream()
	if err != nil {
		log.Errorf("JetStream creation error: %v", err)
		panic(err)
	}
	log.Infof("Nats JetStream created")
	return js
}

//Close closes NATS connections
func Close() {
	busMutex.Lock()
	defer busMutex.Unlock()

	if js != nil {
		js = nil
	}
	if natsConn != nil {
		natsConn.Close()
		natsConn = nil
	}
}

//AddStream add new stream with dedicated subjects
func AddStream(name string, subjects ...string) error {
	js := connection()

	busMutex.Lock()
	defer busMutex.Unlock()

	log.Printf("Getting stream info for [%s]", name)
	stream, err := js.StreamInfo(name)
	if err != nil {
		log.Printf("Error getting stream info for [%s]: %v", name, err)
		//return err
	}
	if stream == nil {
		log.Printf("Adding streams: [%s]", strings.Join(subjects, ", "))
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: subjects,
		})
		if err != nil {
			log.Printf("Error adding streams: %v", err)
			return err
		}
	}
	return nil
}

//PublishMessage publishes message to subject
// subject - subject where message be published
// msg - Message to publish
func PublishMessage(subject string, msg message.Message) error {
	js := connection()

	busMutex.Lock()
	defer busMutex.Unlock()

	log.Printf("Marshalling message ID:[%s]", msg.Id())
	bytes, err := msg.MarshalJson()
	if err != nil {
		log.Errorf("Error marshalling message ID:[%s]: %v", msg.Id(), err)
		return err
	}
	log.Printf("Publishing message ID:[%s]", msg.Id())
	_, err = js.Publish(subject, bytes)
	if err != nil {
		log.Errorf("Error publishing message ID:[%s]: %v", msg.Id(), err)
		return err
	}
	return nil
}

//StartSubscriptions - actually subscribe on NATS jet stream
func StartSubscriptions(consumer string) {
	js := connection()
	for subject := range subscriptions {
		_, err := js.Subscribe(
			subject,
			provideSubscriptionFunc(subject),
			nats.Durable(consumer),
			nats.ManualAck())
		if err != nil {
			log.Errorf("Error setting up consumers on subject [%s]: %v", subject, err)
		}
	}
}

//Subscribe - provide handler func to subject
// subject - subscribe to
// handler - handler func which would be invoked on incoming message
func Subscribe(name string, handler message.Subscriber) {
	_, found := subscriptions[name]
	if !found {
		subscriptions[name] = []message.Subscriber{handler}
	} else {
		subscriptions[name] = append(subscriptions[name], handler)
	}
}

func provideSubscriptionFunc(subject string) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		subscribers, found := subscriptions[subject]
		if !found {
			log.Errorf("Handler func not found for subject [%s]", subject)
		}
		d, err := doc.UnmarshalJson(msg.Data)
		if err != nil {
			log.Errorf("Error unmarshaling message: %v", err)
		}
		m := message.FromDocument(d)
		log.Printf("Invoking handlers for message ID:[%s]", m.Id())
		success := true
		for _, subscriber := range subscribers {
			success = success && subscriber.OnMessage(subject, m)
			log.Printf("Handler finished with success = %t", success)
			if !success {
				break
			}
		}
		if success {
			err := msg.Ack()
			if err != nil {
				log.Errorf("Error sending ACK: %v", err)
			}
		}
		log.Printf("Message ID:[%s] handled successfully = %t", m.Id(), success)
	}
}

func (n *natsSubscriber) OnMessage(topic string, m message.Message) bool {
	err := PublishMessage(topic, m)
	if err != nil {
		log.Errorf("Error while publish message in [%s]: %s", topic, m)
		return false
	}
	return true
}

//Subscriber - returns subscriber which publishes message in NATS stream
func Subscriber(topic string) message.Subscriber {
	s, found := publishers[topic]
	if found {
		return s
	}
	publishers[topic] = &natsSubscriber{}
	return publishers[topic]
}
