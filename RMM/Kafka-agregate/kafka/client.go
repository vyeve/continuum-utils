package kafka

import (
	"log"
	"time"

	"github.com/ContinuumLLC/platform-common-lib/src/messaging"
)

var (
	// DT
	// brokers = []string{
	// 	"172.28.48.249:9092",
	// 	"172.28.48.142:9092",
	// 	"172.28.48.52:9092",
	// }

	// QA
	brokers = []string{
		"172.28.17.30:9092",
		"172.28.16.57:9092",
		"172.28.17.103:9092",
	}
	topics = []string{"managed-endpoint-change"}
)

const groupID = "testGroupID"

// Client represents Kafka connection
var Client *ClientType

// ClientType represents Kafka connection type
type ClientType struct {
	Name            string
	Type            string
	ConnectionURLs  []string
	service         messaging.Service
	topicProcessors map[string]MessageProcessor
}

// Load initializes/executes all components related to processing Kafka messages from managed-endpoint-change topic
func Load() (err error) {
	conf := messaging.Config{
		Address: brokers,
		GroupID: groupID,
		Topics:  topics,
	}

	service := messaging.NewService(conf)

	Client = new(ClientType)

	Client.service = service
	Client.Name = "Kafka session"
	Client.Type = groupID
	Client.ConnectionURLs = brokers

	Client.topicProcessors = make(map[string]MessageProcessor)
	writer := NewWriter()
	RegisterProcessor(topics[0], writer)

	StartProcessors()

	return nil
}

// RegisterProcessor binds a message Processor with a topic.
func RegisterProcessor(topic string, mp MessageProcessor) {
	Client.topicProcessors[topic] = mp
}

// StartProcessors runs messages processing loop.
func StartProcessors() {
	go func(Client *ClientType) {
		var err error
		for {
			log.Println(" Kafka: Start listening")
			err = Client.service.Listen(Client.messageDispatcher)
			log.Printf("Kafka: Failed to start listening (err: %v)\nKafka: Reconnecting...", err)
			time.Sleep(10 * time.Second)
		}
	}(Client)
}

// MessageProcessor describes a behaviour of
type MessageProcessor interface {
	Process(message *messaging.Message)
}

func (c ClientType) messageDispatcher(message *messaging.Message) {
	Processor, ok := c.topicProcessors[message.Topic]
	if !ok {
		log.Fatalf("Kafka: Topic %q is not processed", message.Topic)
	}

	go Processor.Process(message)
}
