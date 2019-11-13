package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
	"os"
	"os/signal"
	"time"
)


func NewMerchantConsumer() {

	brokers := []string{"11.11.5.146:9092"}

	kafkaConfig := getKafkaConfig("", "")

	master, err := sarama.NewConsumer(brokers, kafkaConfig)

	if err != nil {

		panic(err)

	}

	defer func() {

		if err := master.Close(); err != nil {

			panic(err)

		}

	}()

	topic := "new-customer-topic"

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {

		panic(err)

	}

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt)

	doneCh := make(chan struct{})

	go func() {

		for {

			select {

			case err := <-consumer.Errors():

				fmt.Println(err)

			case msg := <-consumer.Messages():

				//*messageCountStart++
				//Deserialize
				merchant := model.Merchant{}
				json.Unmarshal([]byte(msg.Value),&merchant)
				fmt.Println(merchant)
				//presistance.CreateMerchant(merchant)
				fmt.Println("Received messages", string(msg.Topic),string(msg.Key), string(msg.Value))

			case <-signals:

				fmt.Println("Interrupt is detected")

				doneCh <- struct{}{}

			}

		}

	}()

	<-doneCh
	master.Close()
	fmt.Println("Processed",  "messages")

}


func getKafkaConfig(username, password string) *sarama.Config {

	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Producer.Return.Successes = true

	kafkaConfig.Net.WriteTimeout = 5 * time.Second

	kafkaConfig.Producer.Retry.Max = 0



	if username != "" {

		kafkaConfig.Net.SASL.Enable = true

		kafkaConfig.Net.SASL.User = username

		kafkaConfig.Net.SASL.Password = password

	}

	return kafkaConfig

}