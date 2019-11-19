package host

import (
	"github.com/davidnobels/loyalti-go-echo/src/api/host/consumer"
)
func StartKafka() {
	consumer.NewMerchantConsumer()
}
