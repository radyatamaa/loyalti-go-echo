package host

import (
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/consumer"
)
func StartKafka() {
	consumer.NewMerchantConsumer()
}
